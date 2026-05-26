package internal

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var version bool
var output string

// CreateRootCmd creates the root command.
func CreateRootCmd() cobra.Command { _ = "STUB: not implemented"; return *new(cobra.Command) }

// Deprecated: Use godog build, godog run or godog version.
// This is to support the legacy direct usage of the root command.

func runRootCmd(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func bindRootCmdFlags(flagSet *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// Since using the root command directly is deprecated.
// All flags will be hidden
