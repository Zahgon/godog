package formatters

import (
	"io"

	"github.com/cucumber/godog/formatters"
	messages "github.com/cucumber/messages/go/v21"
)

func init() {
	formatters.Format("progress", "Prints a character per step.", ProgressFormatterFunc)
}

// ProgressFormatterFunc implements the FormatterFunc for the progress formatter.
func ProgressFormatterFunc(suite string, out io.Writer) formatters.Formatter {
	_ = "STUB: not implemented"
	return *new(formatters.Formatter)
}

// NewProgress creates a new progress formatter.
func NewProgress(suite string, out io.Writer) *Progress { _ = "STUB: not implemented"; return nil }

// Progress is a minimalistic formatter.
type Progress struct {
	*Base
	StepsPerRow int
	Steps       *int
}

// Summary renders summary information.
func (f *Progress) Summary() { _ = "STUB: not implemented"; return }

func (f *Progress) step(pickleStepID string) { _ = "STUB: not implemented"; return }

// Passed captures passed step.
func (f *Progress) Passed(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Skipped captures skipped step.
func (f *Progress) Skipped(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Undefined captures undefined step.
func (f *Progress) Undefined(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Failed captures failed step.
func (f *Progress) Failed(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition, err error) {
	_ = "STUB: not implemented"
	return
}

// Ambiguous steps.
func (f *Progress) Ambiguous(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition, err error) {
	_ = "STUB: not implemented"
	return
}

// Pending captures pending step.
func (f *Progress) Pending(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}
