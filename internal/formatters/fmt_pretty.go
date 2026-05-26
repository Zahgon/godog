package formatters

import (
	"io"
	"regexp"

	messages "github.com/cucumber/messages/go/v21"

	"github.com/cucumber/godog/colors"
	"github.com/cucumber/godog/formatters"
	"github.com/cucumber/godog/internal/models"
)

func init() {
	formatters.Format("pretty", "Prints every feature with runtime statuses.", PrettyFormatterFunc)
}

// PrettyFormatterFunc implements the FormatterFunc for the pretty formatter
func PrettyFormatterFunc(suite string, out io.Writer) formatters.Formatter {
	_ = "STUB: not implemented"
	return *new(formatters.Formatter)
}

var outlinePlaceholderRegexp = regexp.MustCompile("<[^>]+>")

// Pretty is a formatter for readable output.
type Pretty struct {
	*Base
	firstFeature *bool
}

// TestRunStarted is triggered on test start.
func (f *Pretty) TestRunStarted() { _ = "STUB: not implemented"; return }

// Feature receives gherkin document.
func (f *Pretty) Feature(gd *messages.GherkinDocument, p string, c []byte) {
	_ = "STUB: not implemented"
	return
}

// Pickle takes a gherkin node for formatting.
func (f *Pretty) Pickle(pickle *messages.Pickle) { _ = "STUB: not implemented"; return }

// Passed captures passed step.
func (f *Pretty) Passed(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Skipped captures skipped step.
func (f *Pretty) Skipped(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Undefined captures undefined step.
func (f *Pretty) Undefined(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// Failed captures failed step.
func (f *Pretty) Failed(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition, err error) {
	_ = "STUB: not implemented"
	return
}

// Failed captures failed step.
func (f *Pretty) Ambiguous(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition, err error) {
	_ = "STUB: not implemented"
	return
}

// Pending captures pending step.
func (f *Pretty) Pending(pickle *messages.Pickle, step *messages.PickleStep, match *formatters.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

func (f *Pretty) printFeature(feature *messages.Feature) { _ = "STUB: not implemented"; return }

func keywordAndName(keyword, name string) string { _ = "STUB: not implemented"; return "" }

func (f *Pretty) scenarioLengths(pickle *messages.Pickle) (scenarioHeaderLength int, maxLength int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (f *Pretty) printScenarioHeader(pickle *messages.Pickle, astScenario *messages.Scenario, spaceFilling int) {
	_ = "STUB: not implemented"
	return
}

func (f *Pretty) printUndefinedPickle(pickle *messages.Pickle) { _ = "STUB: not implemented"; return }

//  do not print scenario headers and examples multiple times

// Summary renders summary information.
func (f *Pretty) Summary() { _ = "STUB: not implemented"; return }

func (f *Pretty) printOutlineExample(pickle *messages.Pickle, step *messages.PickleStep, backgroundSteps int) {
	_ = "STUB: not implemented"
	return
}

// do not print empty examples

// do not print examples unless all steps has finished

// determine example row status

// in first example, we need to print steps

// print the step outline

// an example table header

func (f *Pretty) printTableRow(row *messages.TableRow, max []int, clr colors.ColorFunc) {
	_ = "STUB: not implemented"
	return
}

func (f *Pretty) printTableHeader(row *messages.TableRow, max []int) {
	_ = "STUB: not implemented"
	return
}

func isFirstScenarioInRule(rule *messages.Rule, scenario *messages.Scenario) bool {
	_ = "STUB: not implemented"
	return false
}

func isFirstPickleAndNoRule(feature *models.Feature, pickle *messages.Pickle, rule *messages.Rule) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *Pretty) printStep(pickle *messages.Pickle, pickleStep *messages.PickleStep) {
	_ = "STUB: not implemented"
	return
}

func (f *Pretty) printDocString(docString *messages.DocString) { _ = "STUB: not implemented"; return }

// print table with aligned table cells
// @TODO: need to make example header cells bold
func (f *Pretty) printTable(t *messages.PickleTable, c colors.ColorFunc) {
	_ = "STUB: not implemented"
	return
}

// longest gives a list of longest columns of all rows in Table
func maxColLengths(t *messages.PickleTable, clrs ...colors.ColorFunc) []int {
	_ = "STUB: not implemented"
	return nil
}

func longestExampleRow(t *messages.Examples, clrs ...colors.ColorFunc) []int {
	_ = "STUB: not implemented"
	return nil
}

func (f *Pretty) longestStep(steps []*messages.Step, pickleLength int) int {
	_ = "STUB: not implemented"
	return 0
}

// a line number representation in feature file
func line(path string, loc *messages.Location) string {
	_ = "STUB: not implemented"
	// Path can contain a line number already.
	// This line number has to be trimmed to avoid duplication.
	return ""
}

func (f *Pretty) lengthPickleStep(keyword, text string) int { _ = "STUB: not implemented"; return 0 }

func (f *Pretty) lengthPickle(keyword, name string) int { _ = "STUB: not implemented"; return 0 }
