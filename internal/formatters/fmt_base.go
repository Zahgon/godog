package formatters

import (
	"io"
	"sync"

	messages "github.com/cucumber/messages/go/v21"

	"github.com/cucumber/godog/formatters"
	"github.com/cucumber/godog/internal/storage"
)

// BaseFormatterFunc implements the FormatterFunc for the base formatter.
func BaseFormatterFunc(suite string, out io.Writer) formatters.Formatter {
	_ = "STUB: not implemented"
	return *new(formatters.Formatter)
}

// NewBase creates a new base formatter.
func NewBase(suite string, out io.Writer) *Base { _ = "STUB: not implemented"; return nil }

// Base is a base formatter.
type Base struct {
	suiteName string
	out       io.Writer
	indent    int

	Storage *storage.Storage
	Lock    *sync.Mutex
}

// SetStorage assigns gherkin data storage.
func (f *Base) SetStorage(st *storage.Storage) { _ = "STUB: not implemented"; return }

// TestRunStarted is triggered on test start.
func (f *Base) TestRunStarted() {
	_ = "STUB: not implemented"

	// Feature receives gherkin document.
	return
}

func (f *Base) Feature(*messages.GherkinDocument, string, []byte) {
	_ = "STUB: not implemented"

	// Pickle receives scenario.
	return
}

func (f *Base) Pickle(*messages.Pickle) {
	_ = "STUB: not implemented"

	// Defined receives step definition.
	return
}

func (f *Base) Defined(*messages.Pickle, *messages.PickleStep, *formatters.StepDefinition) {
	_ = "STUB: not implemented"

	// Passed captures passed step.
	return
}

func (f *Base) Passed(*messages.Pickle, *messages.PickleStep, *formatters.StepDefinition) {
	_ = "STUB: not implemented"

	// Skipped captures skipped step.
	return
}

func (f *Base) Skipped(*messages.Pickle, *messages.PickleStep, *formatters.StepDefinition) {
	_ = "STUB: not implemented"

	// Undefined captures undefined step.
	return
}

func (f *Base) Undefined(*messages.Pickle, *messages.PickleStep, *formatters.StepDefinition) {
	_ = "STUB: not implemented"

	// Failed captures failed step.
	return
}

func (f *Base) Failed(*messages.Pickle, *messages.PickleStep, *formatters.StepDefinition, error) {
	_ = "STUB: not implemented"

	// Pending captures pending step.
	return
}

func (f *Base) Pending(*messages.Pickle, *messages.PickleStep, *formatters.StepDefinition) {
	_ = "STUB: not implemented"

	// Ambiguous captures ambiguous step.
	return
}

func (f *Base) Ambiguous(*messages.Pickle, *messages.PickleStep, *formatters.StepDefinition, error) {
	_ = "STUB: not implemented"

	// Summary renders summary information.
	return
}

func (f *Base) Summary() { _ = "STUB: not implemented"; return }

// there may be some scenarios without steps

// go 1.5 and 1.6 prints 0 instead of 0s, if duration is zero.

// prints used randomization seed

func asciiTitle(s string) string { _ = "STUB: not implemented"; return "" }

// Snippets returns code suggestions for undefined steps.
func (f *Base) Snippets() string { _ = "STUB: not implemented"; return "" }

// build snippets

// there may be trailing spaces
