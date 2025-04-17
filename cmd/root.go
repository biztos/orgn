// Package cmd holds the Cobra-based command definitions.
//
// It is designed to be easily customizable.
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/biztos/orgn/runner"
)

var Version = "0.1.0"
var Name = "orgn"
var Title = "The Origin Server"
var Description = `More to come!

Eventually!
`

var Config = &runner.Config{}

var Exit = os.Exit // for testability
var Stdout io.Writer = os.Stdout
var Stderr io.Writer = os.Stderr

var configFile string

// RootCmd is the Cobra Root and may be changed to suit after initialization.
var RootCmd = &cobra.Command{
	Use:     Name,        // Only the first word of "Use" applies to root.
	Version: Version,     // Nb: Version is *just* the number e.g. "1.2.3".
	Long:    Description, // Short is ignored in root.
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		// Load any config file we may have.
		if configFile != "" {
			if err := Config.LoadFile(configFile); err != nil {
				return fmt.Errorf("error loading config: %w", err)
			}
		}

		return nil
	},
}

func Execute() {
	RootCmd.SetOut(Stdout)
	RootCmd.SetErr(Stderr)
	err := RootCmd.Execute()
	if err != nil {
		Exit(1) // error message is printed above already.
	}
}

// BailErr bails out with the given error code and error message to Stderr.
func BailErr(code int, err error) {
	fmt.Fprintln(Stderr, err.Error())
	Exit(code)
}

// UpdateInfo updates the help and usage info from the package variables.
func UpdateInfo() {
	RootCmd.Use = Name
	RootCmd.Version = Version
	RootCmd.Long = fmt.Sprintf("%s (%s)\n\n%s", Title, Name, Description)

}

func init() {

	UpdateInfo()

	// Config files:
	RootCmd.PersistentFlags().StringVar(&configFile, "config", "",
		"Config file in TOML format (flags take precedence).")

	// Output control:
	RootCmd.PersistentFlags().BoolVarP(&Config.Debug, "debug", "d", false,
		"Log at DEBUG level (default is INFO).")

}
