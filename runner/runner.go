// Package runner runs the orgn commands.
package runner

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// Config defines the configuration.
type Config struct {
	Debug bool // Log at Debug level (default is Info).
}

// LoadFile loads the config file at path and uses its nonzero values to
// override zero values in c (thus the existing configs take precedence).
func (c *Config) LoadFile(path string) error {

	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	nc := &Config{}
	if err := toml.Unmarshal(b, nc); err != nil {
		return fmt.Errorf("error unmarshaling TOML: %w", err)
	}

	// Local nonzero values win.
	if !c.Debug {
		c.Debug = nc.Debug
	}
	// ...and so on as we add configs.

	return nil
}
