package models

import (
	messages "github.com/cucumber/messages/go/v21"
)

// Feature is an internal object to group together
// the parsed gherkin document, the pickles and the
// raw content.
type Feature struct {
	*messages.GherkinDocument
	Pickles []*messages.Pickle
	Content []byte
}

// FindRule returns the rule to which the given scenario belongs
func (f Feature) FindRule(astScenarioID string) *messages.Rule {
	_ = "STUB: not implemented"
	return nil
}

// FindScenario returns the scenario in the feature or in a rule in the feature
func (f Feature) FindScenario(astScenarioID string) *messages.Scenario {
	_ = "STUB: not implemented"
	return nil
}

// FindBackground ...
func (f Feature) FindBackground(astScenarioID string) *messages.Background {
	_ = "STUB: not implemented"
	return nil
}

// FindExample ...
func (f Feature) FindExample(exampleAstID string) (*messages.Examples, *messages.TableRow) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindStep ...
func (f Feature) FindStep(astStepID string) *messages.Step { _ = "STUB: not implemented"; return nil }
