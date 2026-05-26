package formatters

import (
	"fmt"
	"io"

	"github.com/cucumber/godog/formatters"
	messages "github.com/cucumber/messages/go/v21"
)

const nanoSec = 1000000
const spec = "0.1.0"

func init() {
	formatters.Format("events", fmt.Sprintf("Produces JSON event stream, based on spec: %s.", spec), EventsFormatterFunc)
}

// EventsFormatterFunc implements the FormatterFunc for the events formatter
func EventsFormatterFunc(suite string, out io.Writer) formatters.Formatter {
	_ = "STUB: not implemented"
	return *new(formatters.Formatter)
}

// Events - Events formatter
type Events struct {
	*Base
}

func (f *Events) event(ev interface{}) { _ = "STUB: not implemented"; return }

// Pickle receives scenario.
func (f *Events) Pickle(pickle *messages.Pickle) { _ = "STUB: not implemented"; return }

// @TODO: is status undefined or passed? when there are no steps
// for this scenario

// TestRunStarted is triggered on test start.
func (f *Events) TestRunStarted() { _ = "STUB: not implemented"; return }

// Feature receives gherkin document.
func (f *Events) Feature(ft *messages.GherkinDocument, p string, c []byte) {
	_ = "STUB: not implemented"
	return
}

// Summary pushes summary information to JSON stream.
func (f *Events) Summary() {
	_ = "STUB: not implemented"
	// @TODO: determine status
	return
}

// @TODO not sure that could be correctly implemented

func (f *Events) step(pickle *messages.Pickle, pickleStep *messages.PickleStep) {
	_ = "STUB: not implemented"
	return
}

// Defined receives step definition.
func (f *Events) Defined(pickle *messages.Pickle, pickleStep *messages.PickleStep, def *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Passed captures passed step.
func (f *Events) Passed(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Skipped captures skipped step.
func (f *Events) Skipped(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Undefined captures undefined step.
func (f *Events) Undefined(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Failed captures failed step.
func (f *Events) Failed(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition, err error) {
	_ = "STUB: not implemented"
	return
}

// Pending captures pending step.
func (f *Events) Pending(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Ambiguous captures ambiguous step.
func (f *Events) Ambiguous(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition, err error) {
	_ = "STUB: not implemented"
	return
}

func (f *Events) scenarioLocation(pickle *messages.Pickle) string {
	_ = "STUB: not implemented"
	return ""
}

func isLastStep(pickle *messages.Pickle, step *messages.PickleStep) bool {
	_ = "STUB: not implemented"
	return false
}
