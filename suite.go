package godog

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	messages "github.com/cucumber/messages/go/v21"

	"github.com/cucumber/godog/formatters"
	"github.com/cucumber/godog/internal/models"
	"github.com/cucumber/godog/internal/storage"
)

var (
	errorInterface   = reflect.TypeOf((*error)(nil)).Elem()
	contextInterface = reflect.TypeOf((*context.Context)(nil)).Elem()
)

// more than one regex matched the step text
var ErrAmbiguous = fmt.Errorf("ambiguous step definition")

// ErrUndefined is returned in case if step definition was not found
var ErrUndefined = fmt.Errorf("step is undefined")

// ErrPending should be returned by step definition if
// step implementation is pending
var ErrPending = fmt.Errorf("step implementation is pending")

// ErrSkip should be returned by step definition or a hook if scenario and further steps are to be skipped.
var ErrSkip = fmt.Errorf("skipped")

// StepResultStatus describes step result.
type StepResultStatus = models.StepResultStatus

const (
	// StepPassed indicates step that passed.
	StepPassed StepResultStatus = models.Passed
	// StepFailed indicates step that failed.
	StepFailed = models.Failed
	// StepSkipped indicates step that was skipped.
	StepSkipped = models.Skipped
	// StepUndefined indicates undefined step.
	StepUndefined = models.Undefined
	// StepPending indicates step with pending implementation.
	StepPending = models.Pending
	// StepAmbiguous indicates step text matches more than one step def
	StepAmbiguous = models.Ambiguous
)

type suite struct {
	steps []*models.StepDefinition

	fmt     Formatter
	storage *storage.Storage

	failed        bool
	randomSeed    int64
	stopOnFailure bool
	strict        bool

	defaultContext context.Context
	testingT       *testing.T

	// suite event handlers
	beforeScenarioHandlers []BeforeScenarioHook
	beforeStepHandlers     []BeforeStepHook
	afterStepHandlers      []AfterStepHook
	afterScenarioHandlers  []AfterScenarioHook
}

type Attachment struct {
	Body      []byte
	FileName  string
	MediaType string
}

type attachmentKey struct{}

func Attach(ctx context.Context, attachments ...Attachment) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func Attachments(ctx context.Context) []Attachment { _ = "STUB: not implemented"; return nil }

func clearAttach(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func pickleAttachments(ctx context.Context) []models.PickleAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (s *suite) matchStep(step *messages.PickleStep) (*models.StepDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *suite) runStep(ctx context.Context, pickle *Scenario, step *Step, scenarioErr error, isFirst, isLast bool) (rctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// user multistep definitions may panic

// FailNow or SkipNow called on dogTestingT, so clear the error to let the normal
// below getTestingT(ctx).isFailed() call handle the reasons.

// Check for any calls to Fail on dogT

// Run after step handlers.

// Trigger after scenario on failing or last step to attach possible hook error to step.

// extract any accumulated attachments and clear them

// run before scenario handlers

// run before step handlers

func (s *suite) runBeforeStepHooks(ctx context.Context, step *Step, err error) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (s *suite) runAfterStepHooks(ctx context.Context, step *Step, status StepResultStatus, err error) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// Adding hook error to resulting error without breaking hooks loop.

func (s *suite) runBeforeScenarioHooks(ctx context.Context, pickle *messages.Pickle) (context.Context, error) {
	_ = "STUB: not implemented"

	// run before scenario handlers
	return *new(context.Context), nil
}

func (s *suite) runAfterScenarioHooks(ctx context.Context, pickle *messages.Pickle, lastStepErr error) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// run after scenario handlers

// Adding hook error to resulting error without breaking hooks loop.

func (s *suite) maybeUndefined(ctx context.Context, text string, arg interface{}, stepType messages.PickleStepType) (context.Context, []string, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

// @TODO: we cannot currently parse table or content body from nested steps

func (s *suite) maybeSubSteps(ctx context.Context, result interface{}) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (s *suite) runSubStep(ctx context.Context, text string, def *models.StepDefinition) (_ context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (s *suite) matchStepTextAndType(text string, stepType messages.PickleStepType) (*models.StepDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// since we need to assign arguments
// better to copy the step definition

func keywordMatches(k formatters.Keyword, stepType messages.PickleStepType) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *suite) runSteps(ctx context.Context, pickle *Scenario, steps []*Step) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (s *suite) shouldFail(err error) bool { _ = "STUB: not implemented"; return false }

func (s *suite) runPickle(pickle *messages.Pickle) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Before scenario hooks are called in context of first evaluated step
// so that error from handler can be added to step.

// scenario

// Running scenario as a subtest.

// After scenario handlers are called in context of last evaluated step
// so that error from handler can be added to step.
