package formatters

import (
	"sync"

	"github.com/cucumber/godog/formatters"
	messages "github.com/cucumber/messages/go/v21"
)

// WrapOnFlush wrap a `formatters.Formatter` in a `formatters.FlushFormatter`, which only
// executes when `Flush` is called
func WrapOnFlush(fmt formatters.Formatter) formatters.FlushFormatter {
	_ = "STUB: not implemented"
	return *new(formatters.FlushFormatter)
}

type onFlushFormatter struct {
	fmt formatters.Formatter
	fns []func()
	mu  *sync.Mutex
}

func (o *onFlushFormatter) Pickle(pickle *messages.Pickle) { _ = "STUB: not implemented"; return }

func (o *onFlushFormatter) Passed(pickle *messages.Pickle, step *messages.PickleStep, definition *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Ambiguous implements formatters.Formatter.
func (o *onFlushFormatter) Ambiguous(pickle *messages.Pickle, step *messages.PickleStep, definition *formatters.StepDefinition, err error) {
	_ = "STUB: not implemented"
	return
}

// Defined implements formatters.Formatter.
func (o *onFlushFormatter) Defined(pickle *messages.Pickle, step *messages.PickleStep, definition *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Failed implements formatters.Formatter.
func (o *onFlushFormatter) Failed(pickle *messages.Pickle, step *messages.PickleStep, definition *formatters.StepDefinition, err error) {
	_ = "STUB: not implemented"
	return
}

// Feature implements formatters.Formatter.
func (o *onFlushFormatter) Feature(pickle *messages.GherkinDocument, p string, c []byte) {
	_ = "STUB: not implemented"
	return
}

// Pending implements formatters.Formatter.
func (o *onFlushFormatter) Pending(pickle *messages.Pickle, step *messages.PickleStep, definition *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Skipped implements formatters.Formatter.
func (o *onFlushFormatter) Skipped(pickle *messages.Pickle, step *messages.PickleStep, definition *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Summary implements formatters.Formatter.
func (o *onFlushFormatter) Summary() { _ = "STUB: not implemented"; return }

// TestRunStarted implements formatters.Formatter.
func (o *onFlushFormatter) TestRunStarted() { _ = "STUB: not implemented"; return }

// Undefined implements formatters.Formatter.
func (o *onFlushFormatter) Undefined(pickle *messages.Pickle, step *messages.PickleStep, definition *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Flush the logs.
func (o *onFlushFormatter) Flush() { _ = "STUB: not implemented"; return }
