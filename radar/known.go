package radar

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Endpoint represents one tracked endpoint.
type Endpoint struct {
	Kind   string `yaml:"kind"          json:"kind"`   // rpc | rest | grpc
	URL    string `yaml:"url"           json:"url"`
	Source string `yaml:"source,omitempty" json:"source,omitempty"`
	Notes  string `yaml:"notes,omitempty"  json:"notes,omitempty"`
}

// Known is the on-disk shape of known.yaml.
type Known struct {
	Endpoints []Endpoint `yaml:"endpoints"`
}

// LoadKnown reads known.yaml from path.
func LoadKnown(path string) (Known, error) {
	var k Known
	b, err := os.ReadFile(path)
	if err != nil {
		return k, err
	}
	if err := yaml.Unmarshal(b, &k); err != nil {
		return k, err
	}
	return k, nil
}

// SaveKnown writes known.yaml to path with stable formatting.
func SaveKnown(path string, k Known) error {
	b, err := yaml.Marshal(k)
	if err != nil {
		return err
	}
	header := []byte("# Tracked AssetMantle public endpoints.\n# Keep entries grouped by kind; the discover step appends new ones in-place.\n")
	return os.WriteFile(path, append(header, b...), 0o644)
}
