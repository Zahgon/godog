package main

import (
	"io"

	"github.com/cucumber/godog"
)

const (
	passedEmoji    = "✅"
	skippedEmoji   = "➖"
	failedEmoji    = "❌"
	undefinedEmoji = "❓"
	pendingEmoji   = "🚧"
)

func init() {
	godog.Format("emoji", "Progress formatter with emojis", emojiFormatterFunc)
}

func emojiFormatterFunc(suite string, out io.Writer) godog.Formatter {
	_ = "STUB: not implemented"
	return *new(godog.Formatter)
}

func newEmojiFmt(suite string, out io.Writer) *emojiFmt { _ = "STUB: not implemented"; return nil }

type emojiFmt struct {
	*godog.ProgressFmt

	out io.Writer
}

func (f *emojiFmt) TestRunStarted() { _ = "STUB: not implemented"; return }

func (f *emojiFmt) Passed(scenario *godog.Scenario, step *godog.Step, match *godog.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

func (f *emojiFmt) Skipped(scenario *godog.Scenario, step *godog.Step, match *godog.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

func (f *emojiFmt) Undefined(scenario *godog.Scenario, step *godog.Step, match *godog.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

func (f *emojiFmt) Failed(scenario *godog.Scenario, step *godog.Step, match *godog.StepDefinition, err error) {
	_ = "STUB: not implemented"
	return
}

func (f *emojiFmt) Pending(scenario *godog.Scenario, step *godog.Step, match *godog.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

func (f *emojiFmt) Summary() { _ = "STUB: not implemented"; return }

func (f *emojiFmt) printSummaryLegend() { _ = "STUB: not implemented"; return }

func (f *emojiFmt) step(pickleStepID string) { _ = "STUB: not implemented"; return }
