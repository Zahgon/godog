package formatters

import (
	"regexp"

	messages "github.com/cucumber/messages/go/v21"

	"github.com/cucumber/godog/colors"
	"github.com/cucumber/godog/internal/models"
	"github.com/cucumber/godog/internal/utils"
)

var (
	red    = colors.Red
	redb   = colors.Bold(colors.Red)
	green  = colors.Green
	blackb = colors.Bold(colors.Black)
	yellow = colors.Yellow
	cyan   = colors.Cyan
	cyanb  = colors.Bold(colors.Cyan)
	whiteb = colors.Bold(colors.White)
)

// repeats a space n times
var s = utils.S

var (
	passed    = models.Passed
	failed    = models.Failed
	skipped   = models.Skipped
	undefined = models.Undefined
	pending   = models.Pending
	ambiguous = models.Ambiguous
)

type sortFeaturesByName []*models.Feature

func (s sortFeaturesByName) Len() int           { _ = "STUB: not implemented"; return 0 }
func (s sortFeaturesByName) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (s sortFeaturesByName) Swap(i, j int)      { _ = "STUB: not implemented"; return }

type sortPicklesByID []*messages.Pickle

func (s sortPicklesByID) Len() int           { _ = "STUB: not implemented"; return 0 }
func (s sortPicklesByID) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (s sortPicklesByID) Swap(i, j int) { _ = "STUB: not implemented"; return }

type sortPickleStepResultsByPickleStepID []models.PickleStepResult

func (s sortPickleStepResultsByPickleStepID) Len() int { _ = "STUB: not implemented"; return 0 }
func (s sortPickleStepResultsByPickleStepID) Less(i, j int) bool {
	_ = "STUB: not implemented"
	return false
}

func (s sortPickleStepResultsByPickleStepID) Swap(i, j int) { _ = "STUB: not implemented"; return }

func mustConvertStringToInt(s string) int { _ = "STUB: not implemented"; return 0 }

// DefinitionID ...
func DefinitionID(sd *models.StepDefinition) string { _ = "STUB: not implemented"; return "" }

// case when suite is a structure with methods

// case when steps are just plain funcs

var matchFuncDefRef = regexp.MustCompile(`\(([^\)]+)\)`)
