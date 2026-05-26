package formatters

import (
	"encoding/xml"
	"io"
	"time"

	"github.com/cucumber/godog/formatters"
	"github.com/cucumber/godog/internal/models"
)

func init() {
	formatters.Format("junit", "Prints junit compatible xml to stdout", JUnitFormatterFunc)
}

// JUnitFormatterFunc implements the FormatterFunc for the junit formatter
func JUnitFormatterFunc(suite string, out io.Writer) formatters.Formatter {
	_ = "STUB: not implemented"
	return *new(formatters.Formatter)
}

// JUnit renders test results in JUnit format.
type JUnit struct {
	*Base
}

// Summary renders summary information.
func (f *JUnit) Summary() { _ = "STUB: not implemented"; return }

func junitTimeDuration(from, to time.Time) string { _ = "STUB: not implemented"; return "" }

// getPickleResult deals with the fact that if there's no result due to 'StopOnFirstFailure' being
// set, MustGetPickleResult panics.
func (f *JUnit) getPickleResult(pickleID string) (res *models.PickleResult) {
	_ = "STUB: not implemented"
	return nil
}

func (f *JUnit) getPickleStepResult(stepID string) (res *models.PickleStepResult) {
	_ = "STUB: not implemented"
	return nil
}

func (f *JUnit) getPickleStepResultsByPickleID(pickleID string) (res []models.PickleStepResult) {
	_ = "STUB: not implemented"
	return nil
}

func (f *JUnit) buildJUNITPackageSuite() JunitPackageSuite {
	_ = "STUB: not implemented"
	return *new(JunitPackageSuite)
}

type junitFailure struct {
	Message string `xml:"message,attr"`
	Type    string `xml:"type,attr,omitempty"`
}

type junitError struct {
	XMLName xml.Name `xml:"error,omitempty"`
	Message string   `xml:"message,attr"`
	Type    string   `xml:"type,attr"`
}

type junitTestCase struct {
	XMLName xml.Name      `xml:"testcase"`
	Name    string        `xml:"name,attr"`
	Status  string        `xml:"status,attr"`
	Time    string        `xml:"time,attr"`
	Failure *junitFailure `xml:"failure,omitempty"`
	Error   []*junitError
}

type junitTestSuite struct {
	XMLName   xml.Name `xml:"testsuite"`
	Name      string   `xml:"name,attr"`
	Tests     int      `xml:"tests,attr"`
	Skipped   int      `xml:"skipped,attr"`
	Failures  int      `xml:"failures,attr"`
	Errors    int      `xml:"errors,attr"`
	Time      string   `xml:"time,attr"`
	TestCases []*junitTestCase
}

// JunitPackageSuite ...
type JunitPackageSuite struct {
	XMLName    xml.Name `xml:"testsuites"`
	Name       string   `xml:"name,attr"`
	Tests      int      `xml:"tests,attr"`
	Skipped    int      `xml:"skipped,attr"`
	Failures   int      `xml:"failures,attr"`
	Errors     int      `xml:"errors,attr"`
	Time       string   `xml:"time,attr"`
	TestSuites []*junitTestSuite
}
