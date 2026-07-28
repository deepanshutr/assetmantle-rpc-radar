// Command radar tracks the health and existence of AssetMantle public endpoints.
//
//	radar health    -- probe all entries in known.yaml, write data/state.json + data/REPORT.md
//	radar discover  -- scrape discovery sources, append new entries to known.yaml
//	radar all       -- discover then health
package main

import (
	"fmt"
	"os"

	"github.com/deepanshutr/assetmantle-rpc-radar/radar"
)

const (
	knownPath  = "known.yaml"
	statePath  = "data/state.json"
	reportPath = "data/REPORT.md"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	switch os.Args[1] {
	case "health":
		if err := runHealth(); err != nil {
			fmt.Fprintln(os.Stderr, "health:", err)
			os.Exit(1)
		}
	case "discover":
		if err := runDiscover(); err != nil {
			fmt.Fprintln(os.Stderr, "discover:", err)
			os.Exit(1)
		}
	case "all":
		if err := runDiscover(); err != nil {
			fmt.Fprintln(os.Stderr, "discover:", err)
			os.Exit(1)
		}
		if err := runHealth(); err != nil {
			fmt.Fprintln(os.Stderr, "health:", err)
			os.Exit(1)
		}
	default:
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: radar {health|discover|all}")
}

func runHealth() error {
	known, err := radar.LoadKnown(knownPath)
	if err != nil {
		return err
	}
	// Read the previous run's state before it is overwritten. The health
	// workflow commits data/state.json back to the repo, so the checkout is
	// the baseline: alerts fire on transitions, never on standing state.
	prev, err := radar.ReadState(statePath)
	if err != nil {
		return err
	}
	results := radar.ProbeAll(known.Endpoints)
	if err := radar.WriteReport(reportPath, results); err != nil {
		return err
	}

	// Deliver BEFORE advancing the baseline.
	//
	// state.json records what the operator has already been TOLD, not merely
	// what was last observed. Writing it first and then swallowing a send
	// failure would consume the transition: the next run sees was == now and
	// stays silent, so one dropped Telegram call silences a real outage
	// permanently. That is strictly worse than the daily repetition this
	// change removes — repetition is annoying, a lost page is the failure
	// monitoring exists to prevent.
	//
	// On failure the baseline stays put, the run goes red, and the workflow's
	// commit step is skipped, so the next run re-derives and re-fires.
	alerts := radar.CriticalAlerts(prev, results)
	if len(alerts) > 0 {
		if err := radar.SendAlerts(alerts, runURL()); err != nil {
			return fmt.Errorf("alert delivery failed, baseline deliberately not advanced: %w", err)
		}
	}
	if err := radar.WriteState(statePath, results); err != nil {
		return err
	}
	return nil
}

// runURL links back to the Actions run; empty outside CI so local runs stay bare.
func runURL() string {
	server := os.Getenv("GITHUB_SERVER_URL")
	repo := os.Getenv("GITHUB_REPOSITORY")
	id := os.Getenv("GITHUB_RUN_ID")
	if server == "" || repo == "" || id == "" {
		return ""
	}
	return server + "/" + repo + "/actions/runs/" + id
}

func runDiscover() error {
	known, err := radar.LoadKnown(knownPath)
	if err != nil {
		return err
	}
	added, err := radar.Discover(&known)
	if err != nil {
		return err
	}
	if added == 0 {
		fmt.Println("discover: no new endpoints")
		return nil
	}
	if err := radar.SaveKnown(knownPath, known); err != nil {
		return err
	}
	fmt.Printf("discover: added %d new endpoints\n", added)
	return nil
}
