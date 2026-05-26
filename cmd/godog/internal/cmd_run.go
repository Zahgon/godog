package internal

import (
	"github.com/spf13/cobra"

	"github.com/cucumber/godog/internal/flags"
)

var opts flags.Options

// CreateRunCmd creates the run subcommand.
func CreateRunCmd() cobra.Command { _ = "STUB: not implemented"; return *new(cobra.Command) }

func runCmdRunFunc(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func buildAndRunGodog(args []string) (err error) { _ = "STUB: not implemented"; return nil }

func runGodog(bin string, args []string) (err error) { _ = "STUB: not implemented"; return nil }

// This works on both Unix and Windows. Although package
// syscall is generally platform dependent, WaitStatus is
// defined for both Unix and Windows and in both cases has
// an ExitStatus() method with the same signature.
