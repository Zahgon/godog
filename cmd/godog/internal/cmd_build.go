package internal

import (
	"github.com/spf13/cobra"
)

var buildOutput string
var buildOutputDefault = "godog.test"

// CreateBuildCmd creates the build subcommand.
func CreateBuildCmd() cobra.Command { _ = "STUB: not implemented"; return *new(cobra.Command) }

func buildCmdRunFunc(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
