package formatters

import (
	"io"
	"regexp"

	messages "github.com/cucumber/messages/go/v21"
)

type registeredFormatter struct {
	name        string
	description string
	fmt         FormatterFunc
}

var registeredFormatters []*registeredFormatter

// FindFmt searches available formatters registered
// and returns FormaterFunc matched by given
// format name or nil otherwise
func FindFmt(name string) FormatterFunc { _ = "STUB: not implemented"; return *new(FormatterFunc) }

// Format registers a feature suite output
// formatter by given name, description and
// FormatterFunc constructor function, to initialize
// formatter with the output recorder.
func Format(name, description string, f FormatterFunc) { _ = "STUB: not implemented"; return }

// AvailableFormatters gives a map of all
// formatters registered with their name as key
// and description as value
func AvailableFormatters() map[string]string { _ = "STUB: not implemented"; return nil }

// Formatter is an interface for feature runner
// output summary presentation.
//
// New formatters may be created to represent
// suite results in different ways. These new
// formatters needs to be registered with a
// godog.Format function call
type Formatter interface {
	TestRunStarted()
	Feature(*messages.GherkinDocument, string, []byte)
	Pickle(*messages.Pickle)
	Defined(*messages.Pickle, *messages.PickleStep, *StepDefinition)
	Failed(*messages.Pickle, *messages.PickleStep, *StepDefinition, error)
	Passed(*messages.Pickle, *messages.PickleStep, *StepDefinition)
	Skipped(*messages.Pickle, *messages.PickleStep, *StepDefinition)
	Undefined(*messages.Pickle, *messages.PickleStep, *StepDefinition)
	Pending(*messages.Pickle, *messages.PickleStep, *StepDefinition)
	Ambiguous(*messages.Pickle, *messages.PickleStep, *StepDefinition, error)
	Summary()
}

// FlushFormatter is a `Formatter` but can be flushed.
type FlushFormatter interface {
	Formatter
	Flush()
}

// FormatterFunc builds a formatter with given
// suite name and io.Writer to record output
type FormatterFunc func(string, io.Writer) Formatter

// StepDefinition is a registered step definition
// contains a StepHandler and regexp which
// is used to match a step. Args which
// were matched by last executed step
//
// This structure is passed to the formatter
// when step is matched and is either failed
// or successful
type StepDefinition struct {
	Expr    *regexp.Regexp
	Handler interface{}
	Keyword Keyword
}

type Keyword int64

const (
	Given Keyword = iota
	When
	Then
	None
)
