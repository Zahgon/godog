package godog

import (
	"context"
	"fmt"
	"testing"
)

// T returns a TestingT compatible interface from the current test context. It will return nil if
// called outside the context of a test. This can be used with (for example) testify's assert and
// require packages.
func T(ctx context.Context) TestingT {
	_ = "STUB: not implemented"
	return *

	// TestingT is a subset of the public methods implemented by go's testing.T. It allows assertion
	// libraries to be used with godog, provided they depend only on this subset of methods.
	new(TestingT)
}

type TestingT interface {
	// Name returns the name of the current pickle under test
	Name() string
	// Log will log to the current testing.T log if set, otherwise it will log to stdout
	Log(args ...interface{})
	// Logf will log a formatted string to the current testing.T log if set, otherwise it will log
	// to stdout
	Logf(format string, args ...interface{})
	// Error fails the current test and logs the provided arguments. Equivalent to calling Log then
	// Fail.
	Error(args ...interface{})
	// Errorf fails the current test and logs the formatted message. Equivalent to calling Logf then
	// Fail.
	Errorf(format string, args ...interface{})
	// Fail marks the current test as failed, but does not halt execution of the step.
	Fail()
	// FailNow marks the current test as failed and halts execution of the step.
	FailNow()
	// Fatal logs the provided arguments, marks the test as failed and halts execution of the step.
	Fatal(args ...interface{})
	// Fatal logs the formatted message, marks the test as failed and halts execution of the step.
	Fatalf(format string, args ...interface{})
	// Skip logs the provided arguments and marks the test as skipped but does not halt execution
	// of the step.
	Skip(args ...interface{})
	// Skipf logs the formatted message and marks the test as skipped but does not halt execution
	// of the step.
	Skipf(format string, args ...interface{})
	// SkipNow marks the current test as skipped and halts execution of the step.
	SkipNow()
	// Skipped returns true if the test has been marked as skipped.
	Skipped() bool
}

// Logf will log test output. If called in the context of a test and testing.T has been registered,
// this will log using the step's testing.T, else it will simply log to stdout.
func Logf(ctx context.Context, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Log will log test output. If called in the context of a test and testing.T has been registered,
// this will log using the step's testing.T, else it will simply log to stdout.
func Log(ctx context.Context, args ...interface{}) { _ = "STUB: not implemented"; return }

// LoggedMessages returns an array of any logged messages that have been recorded during the test
// through calls to godog.Log / godog.Logf or via operations against godog.T(ctx)
func LoggedMessages(ctx context.Context) []string { _ = "STUB: not implemented"; return nil }

// errStopNow should be returned inside a panic within the test to immediately halt execution of that
// test
var errStopNow = fmt.Errorf("FailNow or SkipNow called")

type testingT struct {
	name         string
	t            *testing.T
	failed       bool
	skipped      bool
	failMessages []string
	logMessages  []string
}

// check interface against our testingT and the upstream testing.B/F/T:
var (
	_ TestingT = &testingT{}
	_ TestingT = (*testing.T)(nil)
)

func (dt *testingT) Name() string { _ = "STUB: not implemented"; return "" }

func (dt *testingT) Log(args ...interface{}) { _ = "STUB: not implemented"; return }

func (dt *testingT) Logf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (dt *testingT) Error(args ...interface{}) { _ = "STUB: not implemented"; return }

func (dt *testingT) Errorf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (dt *testingT) Fail() { _ = "STUB: not implemented"; return }

func (dt *testingT) FailNow() { _ = "STUB: not implemented"; return }

func (dt *testingT) Fatal(args ...interface{}) { _ = "STUB: not implemented"; return }

func (dt *testingT) Fatalf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (dt *testingT) Skip(args ...interface{}) { _ = "STUB: not implemented"; return }

func (dt *testingT) Skipf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (dt *testingT) SkipNow() { _ = "STUB: not implemented"; return }

func (dt *testingT) Skipped() bool {
	_ = "STUB: not implemented"

	// isFailed will return an error representing the calls to Fail made during this test
	return false
}

func (dt *testingT) isFailed() error { _ = "STUB: not implemented"; return nil }

type testingTCtxVal struct{}

func setContextTestingT(ctx context.Context, dt *testingT) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getTestingT(ctx context.Context) *testingT { _ = "STUB: not implemented"; return nil }
