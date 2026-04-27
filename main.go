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
	results := radar.ProbeAll(known.Endpoints)
	if err := radar.WriteState(statePath, results); err != nil {
		return err
	}
	if err := radar.WriteReport(reportPath, results); err != nil {
		return err
	}
	if alerts := radar.CriticalAlerts(results); len(alerts) > 0 {
		if err := radar.SendAlerts(alerts); err != nil {
			fmt.Fprintln(os.Stderr, "alert:", err)
		}
	}
	return nil
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
