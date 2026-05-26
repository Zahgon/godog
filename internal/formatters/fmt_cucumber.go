package formatters

/*
   The specification for the formatting originated from https://www.relishapp.com/cucumber/cucumber/docs/formatters/json-output-formatter.
   I found that documentation was misleading or out dated.  To validate formatting I create a ruby cucumber test harness and ran the
   same feature files through godog and the ruby cucumber.

   The docstrings in the cucumber.feature represent the cucumber output for those same feature definitions.

   I did note that comments in ruby could be at just about any level in particular Feature, Scenario and Step.  In godog I
   could only find comments under the Feature data structure.
*/

import (
	"io"

	"github.com/cucumber/godog/formatters"
	"github.com/cucumber/godog/internal/models"
	messages "github.com/cucumber/messages/go/v21"
)

func init() {
	formatters.Format("cucumber", "Produces cucumber JSON format output.", CucumberFormatterFunc)
}

// CucumberFormatterFunc implements the FormatterFunc for the cucumber formatter
func CucumberFormatterFunc(suite string, out io.Writer) formatters.Formatter {
	_ = "STUB: not implemented"
	return *new(formatters.Formatter)
}

// Cuke ...
type Cuke struct {
	*Base
}

// Summary renders test result as Cucumber JSON.
func (f *Cuke) Summary() { _ = "STUB: not implemented"; return }

func (f *Cuke) buildCukeFeatures(features []*models.Feature) (res []CukeFeatureJSON) {
	_ = "STUB: not implemented"
	return nil
}

func (f *Cuke) buildCukeElements(pickles []*messages.Pickle) (res []cukeElement) {
	_ = "STUB: not implemented"
	return nil
}

type cukeComment struct {
	Value string `json:"value"`
	Line  int    `json:"line"`
}

type cukeDocstring struct {
	Value       string `json:"value"`
	ContentType string `json:"content_type"`
	Line        int    `json:"line"`
}

type cukeTag struct {
	Name string `json:"name"`
	Line int    `json:"line"`
}

type cukeResult struct {
	Status   string `json:"status"`
	Error    string `json:"error_message,omitempty"`
	Duration *int   `json:"duration,omitempty"`
}

type cukeMatch struct {
	Location string `json:"location"`
}

type cukeEmbedding struct {
	Name     string `json:"name"`
	MimeType string `json:"mime_type"`
	Data     string `json:"data"`
}

type cukeStep struct {
	Keyword    string              `json:"keyword"`
	Name       string              `json:"name"`
	Line       int                 `json:"line"`
	Docstring  *cukeDocstring      `json:"doc_string,omitempty"`
	Match      cukeMatch           `json:"match"`
	Result     cukeResult          `json:"result"`
	DataTable  []*cukeDataTableRow `json:"rows,omitempty"`
	Embeddings []cukeEmbedding     `json:"embeddings,omitempty"`
}

type cukeDataTableRow struct {
	Cells []string `json:"cells"`
}

type cukeElement struct {
	ID          string     `json:"id"`
	Keyword     string     `json:"keyword"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Line        int        `json:"line"`
	Type        string     `json:"type"`
	Tags        []cukeTag  `json:"tags,omitempty"`
	Steps       []cukeStep `json:"steps,omitempty"`
}

// CukeFeatureJSON ...
type CukeFeatureJSON struct {
	URI         string        `json:"uri"`
	ID          string        `json:"id"`
	Keyword     string        `json:"keyword"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Line        int           `json:"line"`
	Comments    []cukeComment `json:"comments,omitempty"`
	Tags        []cukeTag     `json:"tags,omitempty"`
	Elements    []cukeElement `json:"elements,omitempty"`
}

func buildCukeFeature(feat *models.Feature) CukeFeatureJSON {
	_ = "STUB: not implemented"
	return *new(CukeFeatureJSON)
}

func (f *Cuke) buildCukeElement(pickle *messages.Pickle) (cukeElement cukeElement) {
	_ = "STUB: not implemented"
	return *new(cukeElement)
}

func (f *Cuke) buildCukeStep(pickle *messages.Pickle, stepResult models.PickleStepResult) (cukeStep cukeStep) {
	_ = "STUB: not implemented"
	return *new(cukeStep)
}

func makeCukeID(name string) string { _ = "STUB: not implemented"; return "" }
