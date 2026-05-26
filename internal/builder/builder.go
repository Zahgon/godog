package builder

import (
	"go/build"
	"path/filepath"
	"text/template"
)

var (
	tooldir         = findToolDir()
	compiler        = filepath.Join(tooldir, "compile")
	linker          = filepath.Join(tooldir, "link")
	gopaths         = filepath.SplitList(build.Default.GOPATH)
	godogImportPath = "github.com/cucumber/godog"

	// godep
	runnerTemplate = template.Must(template.New("testmain").Parse(`package main

import (
	"github.com/cucumber/godog"
	{{if or .TestSuiteContexts .ScenarioContexts}}_test "{{.ImportPath}}"{{end}}
	{{if or .XTestSuiteContexts .XScenarioContexts}}_xtest "{{.ImportPath}}_test"{{end}}
	{{if or .XTestSuiteContexts .XScenarioContexts}}"testing/internal/testdeps"{{end}}
	"os"
)

{{if or .XTestSuiteContexts .XScenarioContexts}}
func init() {
	testdeps.ImportPath = "{{.ImportPath}}"
}
{{end}}

func main() {
	status := godog.TestSuite{
		Name: "{{ .Name }}",
		TestSuiteInitializer: func (ctx *godog.TestSuiteContext) {
			os.Setenv("GODOG_TESTED_PACKAGE", "{{.ImportPath}}")
			{{range .TestSuiteContexts}}
			_test.{{ . }}(ctx)
			{{end}}
			{{range .XTestSuiteContexts}}
			_xtest.{{ . }}(ctx)
			{{end}}
		},
		ScenarioInitializer: func (ctx *godog.ScenarioContext) {
			{{range .ScenarioContexts}}
			_test.{{ . }}(ctx)
			{{end}}
			{{range .XScenarioContexts}}
			_xtest.{{ . }}(ctx)
			{{end}}
		},
	}.Run()

	os.Exit(status)
}`))

	// temp file for import
	tempFileTemplate = template.Must(template.New("temp").Parse(`package {{.Name}}

import "github.com/cucumber/godog"

var _ = godog.Version
`))
)

// Build creates a test package like go test command at given target path.
// If there are no go files in tested directory, then
// it simply builds a godog executable to scan features.
//
// If there are go test files, it first builds a test
// package with standard go test command.
//
// Finally it generates godog suite executable which
// registers exported godog contexts from the test files
// of tested package.
//
// Returns the path to generated executable
func Build(bin string) error { _ = "STUB: not implemented"; return nil }

// we allow package to be nil, if godog is run only when
// there is a feature file in empty directory

// may need to produce temp file for godog dependency

// @TODO: in case of modules we cannot build it our selves, we need to have this hacky option

// build and compile the tested package.
// generated test executable will be removed
// since we do not need it for godog suite.
// we also print back the temp WORK directory
// go has built. We will reuse it for our suite workdir.

// extract go-build temporary directory as our workdir

// it may have some compilation warnings, in the output, but these are not
// considered to be errors, since command exit status is 0

// may not locate it in output

// check whether workdir exists

// replace _testmain.go file with our own

// godog package may be vendored and may need importmap

// compile godog testmain package archive
// we do not depend on CGO so a lot of checks are not necessary

// link test suite executable

// filterImportCfg strips unsupported lines from imports configuration.
func filterImportCfg(path string) error { _ = "STUB: not implemented"; return nil }

func maybeVendoredGodog() *build.Package { _ = "STUB: not implemented"; return nil }

func normaliseLocalImportPath(dir string) string { _ = "STUB: not implemented"; return "" }

func importPackage(dir string) *build.Package { _ = "STUB: not implemented"; return nil }

// normalize import path for local import packages
// taken from go source code
// see: https://github.com/golang/go/blob/go1.7rc5/src/cmd/go/pkg.go#L279

// from go src
func makeImportValid(r rune) rune {
	_ = "STUB: not implemented"
	// Should match Go spec, compilers, and ../../go/parser/parser.go:/isValidImport.
	return 0
}

// build temporary file content if godog
// package is not present in currently tested package
func buildTempFile(pkg *build.Package) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// maybe we are testing the godog package on it's own

// buildTestMain if given package is valid
// it scans test files for contexts
// and produces a testmain source code.
func buildTestMain(pkg *build.Package) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// parseImport parses the import path to deal with go module.
func parseImport(rawPath, rootPath string) string {
	_ = "STUB: not implemented"
	// with go > 1.11 and go module enabled out of the GOPATH,
	// the import path begins with an underscore and the GOPATH is unknown on build.
	return ""
}

// go < 1.11 or it's a module inside the GOPATH

// for module support, query the module import path

// Unable to read stdout

// Does not using modules

// Unexpected result

// Concatenates the module path with the current sub-folders if needed

type contexts struct {
	deprecatedFeatureCtxs []string
	testSuiteCtxs         []string
	scenarioCtxs          []string
}

func (ctxs contexts) validate() error { _ = "STUB: not implemented"; return nil }

// processPackageTestFiles runs through ast of each test
// file pack and looks for godog suite contexts to register
// on run
func processPackageTestFiles(packs ...[]string) (ctxs contexts, _ error) {
	_ = "STUB: not implemented"
	return *new(contexts), nil
}

func findToolDir() string { _ = "STUB: not implemented"; return "" }
