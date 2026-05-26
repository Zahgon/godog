package storage

import (
	"sync"

	messages "github.com/cucumber/messages/go/v21"
	"github.com/hashicorp/go-memdb"

	"github.com/cucumber/godog/internal/models"
)

const (
	writeMode bool = true
	readMode  bool = false

	tableFeature         string = "feature"
	tableFeatureIndexURI string = "id"

	tablePickle         string = "pickle"
	tablePickleIndexID  string = "id"
	tablePickleIndexURI string = "uri"

	tablePickleStep        string = "pickle_step"
	tablePickleStepIndexID string = "id"

	tablePickleResult              string = "pickle_result"
	tablePickleResultIndexPickleID string = "id"

	tablePickleStepResult                  string = "pickle_step_result"
	tablePickleStepResultIndexPickleStepID string = "id"
	tablePickleStepResultIndexPickleID     string = "pickle_id"
	tablePickleStepResultIndexStatus       string = "status"

	tableStepDefintionMatch            string = "step_defintion_match"
	tableStepDefintionMatchIndexStepID string = "id"
)

// Storage is a thread safe in-mem storage
type Storage struct {
	db *memdb.MemDB

	testRunStarted     models.TestRunStarted
	testRunStartedLock *sync.Mutex
}

// NewStorage will create an in-mem storage that
// is used across concurrent runners and formatters
func NewStorage() *Storage { _ = "STUB: not implemented"; return nil }

// MustInsertPickle will insert a pickle and it's steps,
// will panic on error.
func (s *Storage) MustInsertPickle(p *messages.Pickle) { _ = "STUB: not implemented"; return }

// MustGetPickle will retrieve a pickle by id and panic on error.
func (s *Storage) MustGetPickle(id string) *messages.Pickle { _ = "STUB: not implemented"; return nil }

// MustGetPickles will retrieve pickles by URI and panic on error.
func (s *Storage) MustGetPickles(uri string) (ps []*messages.Pickle) {
	_ = "STUB: not implemented"
	return nil
}

// MustGetPickleStep will retrieve a pickle step and panic on error.
func (s *Storage) MustGetPickleStep(id string) *messages.PickleStep {
	_ = "STUB: not implemented"
	return nil
}

// MustInsertTestRunStarted will set the test run started event and panic on error.
func (s *Storage) MustInsertTestRunStarted(trs models.TestRunStarted) {
	_ = "STUB: not implemented"
	return
}

// MustGetTestRunStarted will retrieve the test run started event and panic on error.
func (s *Storage) MustGetTestRunStarted() models.TestRunStarted {
	_ = "STUB: not implemented"
	return *new(models.TestRunStarted)
}

// MustInsertPickleResult will instert a pickle result and panic on error.
func (s *Storage) MustInsertPickleResult(pr models.PickleResult) { _ = "STUB: not implemented"; return }

// MustInsertPickleStepResult will insert a pickle step result and panic on error.
func (s *Storage) MustInsertPickleStepResult(psr models.PickleStepResult) {
	_ = "STUB: not implemented"
	return
}

// MustGetPickleResult will retrieve a pickle result by id and panic on error.
func (s *Storage) MustGetPickleResult(id string) models.PickleResult {
	_ = "STUB: not implemented"
	return *new(models.PickleResult)
}

// MustGetPickleResults will retrieve all pickle results and panic on error.
func (s *Storage) MustGetPickleResults() (prs []models.PickleResult) {
	_ = "STUB: not implemented"
	return nil
}

// MustGetPickleStepResult will retrieve a pickle strep result by id and panic on error.
func (s *Storage) MustGetPickleStepResult(id string) models.PickleStepResult {
	_ = "STUB: not implemented"
	return *new(models.PickleStepResult)
}

// MustGetPickleStepResultsByPickleID will retrieve pickle step results by pickle id and panic on error.
func (s *Storage) MustGetPickleStepResultsByPickleID(pickleID string) (psrs []models.PickleStepResult) {
	_ = "STUB: not implemented"
	return nil
}

// MustGetPickleStepResultsByPickleIDUntilStep will retrieve pickle step results by pickle id
// from 0..stepID for that pickle.
func (s *Storage) MustGetPickleStepResultsByPickleIDUntilStep(pickleID string, untilStepID string) (psrs []models.PickleStepResult) {
	_ = "STUB: not implemented"
	return nil
}

// MustGetPickleStepResultsByStatus will retrieve pickle strep results by status and panic on error.
func (s *Storage) MustGetPickleStepResultsByStatus(status models.StepResultStatus) (psrs []models.PickleStepResult) {
	_ = "STUB: not implemented"
	return nil
}

// MustInsertFeature will insert a feature and panic on error.
func (s *Storage) MustInsertFeature(f *models.Feature) { _ = "STUB: not implemented"; return }

// MustGetFeature will retrieve a feature by URI and panic on error.
func (s *Storage) MustGetFeature(uri string) *models.Feature { _ = "STUB: not implemented"; return nil }

// MustGetFeatures will retrieve all features by and panic on error.
func (s *Storage) MustGetFeatures() (fs []*models.Feature) { _ = "STUB: not implemented"; return nil }

type stepDefinitionMatch struct {
	StepID         string
	StepDefinition *models.StepDefinition
}

// MustInsertStepDefintionMatch will insert the matched StepDefintion for the step ID and panic on error.
func (s *Storage) MustInsertStepDefintionMatch(stepID string, match *models.StepDefinition) {
	_ = "STUB: not implemented"
	return
}

// MustGetStepDefintionMatch will retrieve the matched StepDefintion for the step ID and panic on error.
func (s *Storage) MustGetStepDefintionMatch(stepID string) *models.StepDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *Storage) mustInsert(table string, obj interface{}) { _ = "STUB: not implemented"; return }

func (s *Storage) mustFirst(table, index string, args ...interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (s *Storage) mustGet(table, index string, args ...interface{}) memdb.ResultIterator {
	_ = "STUB: not implemented"
	return *new(memdb.ResultIterator)
}
