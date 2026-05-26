package godog

import (
	"io"

	"github.com/cucumber/godog/formatters"
	internal_fmt "github.com/cucumber/godog/internal/formatters"
	"github.com/cucumber/godog/internal/models"
	"github.com/cucumber/godog/internal/storage"
)

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
type Formatter = formatters.Formatter

type storageFormatter interface {
	SetStorage(*storage.Storage)
}

// FormatterFunc builds a formatter with given
// suite name and io.Writer to record output
type FormatterFunc = formatters.FormatterFunc

func printStepDefinitions(steps []*models.StepDefinition, w io.Writer) {
	_ = "STUB: not implemented"
	return
}

// NewBaseFmt creates a new base formatter.
func NewBaseFmt(suite string, out io.Writer) *BaseFmt { _ = "STUB: not implemented"; return nil }

// NewProgressFmt creates a new progress formatter.
func NewProgressFmt(suite string, out io.Writer) *ProgressFmt {
	_ = "STUB: not implemented"
	return nil
}

// NewPrettyFmt creates a new pretty formatter.
func NewPrettyFmt(suite string, out io.Writer) *PrettyFmt { _ = "STUB: not implemented"; return nil }

// NewEventsFmt creates a new event streaming formatter.
func NewEventsFmt(suite string, out io.Writer) *EventsFmt { _ = "STUB: not implemented"; return nil }

// NewCukeFmt creates a new Cucumber JSON formatter.
func NewCukeFmt(suite string, out io.Writer) *CukeFmt { _ = "STUB: not implemented"; return nil }

// NewJUnitFmt creates a new JUnit formatter.
func NewJUnitFmt(suite string, out io.Writer) *JUnitFmt { _ = "STUB: not implemented"; return nil }

// BaseFmt exports Base formatter.
type BaseFmt = internal_fmt.Base

// ProgressFmt exports Progress formatter.
type ProgressFmt = internal_fmt.Progress

// PrettyFmt exports Pretty formatter.
type PrettyFmt = internal_fmt.Pretty

// EventsFmt exports Events formatter.
type EventsFmt = internal_fmt.Events

// CukeFmt exports Cucumber JSON formatter.
type CukeFmt = internal_fmt.Cuke

// JUnitFmt exports JUnit formatter.
type JUnitFmt = internal_fmt.JUnit
