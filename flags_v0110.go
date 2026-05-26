package godog

import (
	"github.com/spf13/pflag"
)

// Choose randomly assigns a convenient pseudo-random seed value.
// The resulting seed will be between `1-99999` for later ease of specification.
func makeRandomSeed() int64 { _ = "STUB: not implemented"; return 0 }

func flagSet(opt *Options) *pflag.FlagSet { _ = "STUB: not implemented"; return nil }

// BindCommandLineFlags binds godog flags to given flag set prefixed
// by given prefix, without overriding usage
func BindCommandLineFlags(prefix string, opts *Options) { _ = "STUB: not implemented"; return }
