package formatters

import (
	"io"

	"github.com/cucumber/godog/formatters"
	"github.com/cucumber/godog/internal/storage"
	messages "github.com/cucumber/messages/go/v21"
)

// MultiFormatter passes test progress to multiple formatters.
type MultiFormatter struct {
	formatters []formatter
	repeater   repeater
}

type formatter struct {
	fmt formatters.FormatterFunc
	out io.Writer
}

type repeater []formatters.Formatter

type storageFormatter interface {
	SetStorage(s *storage.Storage)
}

// SetStorage passes storage to all added formatters.
func (r repeater) SetStorage(s *storage.Storage) { _ = "STUB: not implemented"; return }

// TestRunStarted triggers TestRunStarted for all added formatters.
func (r repeater) TestRunStarted() { _ = "STUB: not implemented"; return }

// Feature triggers Feature for all added formatters.
func (r repeater) Feature(document *messages.GherkinDocument, s string, bytes []byte) {
	_ = "STUB: not implemented"
	return
}

// Pickle triggers Pickle for all added formatters.
func (r repeater) Pickle(pickle *messages.Pickle) { _ = "STUB: not implemented"; return }

// Defined triggers Defined for all added formatters.
func (r repeater) Defined(pickle *messages.Pickle, step *messages.PickleStep, definition *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Failed triggers Failed for all added formatters.
func (r repeater) Failed(pickle *messages.Pickle, step *messages.PickleStep, definition *formatters.StepDefinition, err error) {
	_ = "STUB: not implemented"
	return
}

// Passed triggers Passed for all added formatters.
func (r repeater) Passed(pickle *messages.Pickle, step *messages.PickleStep, definition *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Skipped triggers Skipped for all added formatters.
func (r repeater) Skipped(pickle *messages.Pickle, step *messages.PickleStep, definition *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Undefined triggers Undefined for all added formatters.
func (r repeater) Undefined(pickle *messages.Pickle, step *messages.PickleStep, definition *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Pending triggers Pending for all added formatters.
func (r repeater) Pending(pickle *messages.Pickle, step *messages.PickleStep, definition *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Ambiguous triggers Ambiguous for all added formatters.
func (r repeater) Ambiguous(pickle *messages.Pickle, step *messages.PickleStep, definition *formatters.StepDefinition, err error) {
	_ = "STUB: not implemented"
	return
}

// Summary triggers Summary for all added formatters.
func (r repeater) Summary() { _ = "STUB: not implemented"; return }

// Add adds formatter with output writer.
func (m *MultiFormatter) Add(name string, out io.Writer) { _ = "STUB: not implemented"; return }

// FormatterFunc implements the FormatterFunc for the multi formatter.
func (m *MultiFormatter) FormatterFunc(suite string, out io.Writer) formatters.Formatter {
	_ = "STUB: not implemented"
	return *new(formatters.Formatter)
}
