package godog

import (
	"context"
	"testing"

	"github.com/cucumber/godog/internal/models"
	"github.com/cucumber/godog/internal/storage"
)

const (
	exitSuccess int = iota
	exitFailure
	exitOptionError
)

type (
	testSuiteInitializer func(*TestSuiteContext)
	scenarioInitializer  func(*ScenarioContext)
)

type runner struct {
	randomSeed            int64
	stopOnFailure, strict bool

	defaultContext context.Context
	testingT       *testing.T

	features []*models.Feature

	testSuiteInitializer testSuiteInitializer
	scenarioInitializer  scenarioInitializer

	storage *storage.Storage
	fmt     Formatter
}

func (r *runner) concurrent(rate int) (failed bool) { _ = "STUB: not implemented"; return false }

// run before suite handlers

// reserve space in queue

// free a space in queue

// Copy base suite.

// if running concurrently, only print at end of scenario to keep
// scenario logs segregated

// Running within the same goroutine for concurrency 1
// to preserve original stacks and simplify debugging.

// wait until last are processed

// run after suite handlers

// print summary

func runWithOptions(suiteName string, runner runner, opt Options) int {
	_ = "STUB: not implemented"
	return 0
}

// user may have specified -1 option to create random seed

// store chosen seed in environment, so it could be seen in formatter summary report

// determine tested package

// @TODO: should prevent from having these

func runsFromPackage(fp string) string { _ = "STUB: not implemented"; return "" }

// TestSuite allows for configuration
// of the Test Suite Execution
type TestSuite struct {
	Name                 string
	TestSuiteInitializer func(*TestSuiteContext)
	ScenarioInitializer  func(*ScenarioContext)
	Options              *Options
}

// Run will execute the test suite.
//
// If options are not set, it will reads
// all configuration options from flags.
//
// The exit codes may vary from:
//
//	0 - success
//	1 - failed
//	2 - command line usage error
//	128 - or higher, os signal related error exit codes
//
// If there are flag related errors they will be directed to os.Stderr
func (ts TestSuite) Run() int { _ = "STUB: not implemented"; return 0 }

// RetrieveFeatures will parse and return the features based on test suite option
// Any modification on the parsed features will not have any impact on the next Run of the Test Suite
func (ts TestSuite) RetrieveFeatures() ([]*models.Feature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDefaultOptions() (*Options, error) { _ = "STUB: not implemented"; return nil, nil }
