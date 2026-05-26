package models

import (
	"time"

	"github.com/cucumber/godog/colors"
)

// TestRunStarted ...
type TestRunStarted struct {
	StartedAt time.Time
}

// PickleResult ...
type PickleResult struct {
	PickleID  string
	StartedAt time.Time
}

// PickleAttachment ...
type PickleAttachment struct {
	Name     string
	MimeType string
	Data     []byte
}

// PickleStepResult ...
type PickleStepResult struct {
	Status     StepResultStatus
	FinishedAt time.Time
	Err        error

	PickleID     string
	PickleStepID string

	Def *StepDefinition

	Attachments []PickleAttachment
}

// NewStepResult ...
func NewStepResult(
	status StepResultStatus,
	pickleID, pickleStepID string,
	match *StepDefinition,
	attachments []PickleAttachment,
	err error,
) PickleStepResult {
	_ = "STUB: not implemented"
	return *new(PickleStepResult)
}

// StepResultStatus ...
type StepResultStatus int

const (
	// Passed ...
	Passed StepResultStatus = iota
	// Failed ...
	Failed
	// Skipped ...
	Skipped
	// Undefined ...
	Undefined
	// Pending ...
	Pending
	// Ambiguous ...
	Ambiguous
)

// Color ...
func (st StepResultStatus) Color() colors.ColorFunc {
	_ = "STUB: not implemented"
	return *new(colors.ColorFunc)
}

// String ...
func (st StepResultStatus) String() string { _ = "STUB: not implemented"; return "" }
