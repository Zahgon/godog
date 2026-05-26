package formatters

import (
	"regexp"
	"text/template"

	messages "github.com/cucumber/messages/go/v21"
)

// some snippet formatting regexps
var snippetExprCleanup = regexp.MustCompile(`([\/\[\]\(\)\\^\$\.\|\?\*\+\'])`)
var snippetExprQuoted = regexp.MustCompile(`(\W|^)"(?:[^"]*)"(\W|$)`)
var snippetMethodName = regexp.MustCompile(`[^a-zA-Z\_\ ]`)
var snippetNumbers = regexp.MustCompile(`(\d+)`)

var snippetHelperFuncs = template.FuncMap{
	"backticked": func(s string) string {
		return "`" + s + "`"
	},
}

var undefinedSnippetsTpl = template.Must(template.New("snippets").Funcs(snippetHelperFuncs).Parse(`
{{ range . }}func {{ .Method }}({{ .Args }}) error {
	return godog.ErrPending
}

{{end}}func InitializeScenario(ctx *godog.ScenarioContext) { {{ range . }}
	ctx.Step({{ backticked .Expr }}, {{ .Method }}){{end}}
}
`))

type undefinedSnippet struct {
	Method   string
	Expr     string
	argument *messages.PickleStepArgument
}

func (s undefinedSnippet) Args() (ret string) { _ = "STUB: not implemented"; return "" }

type snippetSortByMethod []undefinedSnippet

func (s snippetSortByMethod) Len() int { _ = "STUB: not implemented"; return 0 }

func (s snippetSortByMethod) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s snippetSortByMethod) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
