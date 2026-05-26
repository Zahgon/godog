package godog

import (
	"flag"
	"io"

	"github.com/cucumber/godog/colors"
	"github.com/cucumber/godog/internal/utils"
)

// repeats a space n times
var s = utils.S

var descFeaturesArgument = "Optional feature(s) to run. Can be:\n" +
	s(4) + "- dir " + colors.Yellow("(features/)") + "\n" +
	s(4) + "- feature " + colors.Yellow("(*.feature)") + "\n" +
	s(4) + "- scenario at specific line " + colors.Yellow("(*.feature:10)") + "\n" +
	"If no feature paths are listed, suite tries " + colors.Yellow("features") + " path by default.\n" +
	"Multiple comma-separated values can be provided.\n"

var descConcurrencyOption = "Run the test suite with concurrency level:\n" +
	s(4) + "- " + colors.Yellow(`= 1`) + ": supports all types of formats.\n" +
	s(4) + "- " + colors.Yellow(`>= 2`) + ": only supports " + colors.Yellow("progress") + ". Note, that\n" +
	s(4) + "your context needs to support parallel execution."

var descTagsOption = "Filter scenarios by tags. Expression can be:\n" +
	s(4) + "- " + colors.Yellow(`"@wip"`) + ": run all scenarios with wip tag\n" +
	s(4) + "- " + colors.Yellow(`"~@wip"`) + ": exclude all scenarios with wip tag\n" +
	s(4) + "- " + colors.Yellow(`"@wip && ~@new"`) + ": run wip scenarios, but exclude new\n" +
	s(4) + "- " + colors.Yellow(`"@wip,@undone"`) + ": run wip or undone scenarios"

var descRandomOption = "Randomly shuffle the scenario execution order.\n" +
	"Specify SEED to reproduce the shuffling from a previous run.\n" +
	s(4) + `e.g. ` + colors.Yellow(`--random`) + " or " + colors.Yellow(`--random=5738`)

// FlagSet allows to manage flags by external suite runner
// builds flag.FlagSet with godog flags binded
//
// Deprecated:
func FlagSet(opt *Options) *flag.FlagSet { _ = "STUB: not implemented"; return nil }

// BindFlags binds godog flags to given flag set prefixed
// by given prefix, without overriding usage
func BindFlags(prefix string, set *flag.FlagSet, opt *Options) { _ = "STUB: not implemented"; return }

// override flag defaults if any corresponding properties were supplied on the incoming `opt`

type flagged struct {
	short, long, descr, dflt string
}

func (f *flagged) name() string { _ = "STUB: not implemented"; return "" }

// `random` is special in that we will later assign it randomly
// if the user specifies `--random` without specifying one,
// so mask the "default" value here to avoid UI confusion about
// what the value will end up being.

func usage(set *flag.FlagSet, w io.Writer) func() { _ = "STUB: not implemented"; return nil }

// prints an option or argument with a description, or only description

// --- GENERAL ---

// --- OPTIONS ---

// randomSeed implements `flag.Value`, see https://golang.org/pkg/flag/#Value
type randomSeed struct {
	ref *int64
}

func (rs *randomSeed) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (rs *randomSeed) String() string { _ = "STUB: not implemented"; return "" }

// If a Value has an IsBoolFlag() bool method returning true, the command-line
// parser makes -name equivalent to -name=true rather than using the next
// command-line argument.
func (rs *randomSeed) IsBoolFlag() bool { _ = "STUB: not implemented"; return false }
