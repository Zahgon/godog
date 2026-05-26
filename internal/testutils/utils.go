package testutils

import (
	"testing"

	"github.com/cucumber/godog/internal/models"
)

// BuildTestFeature creates a feature for testing purpose.
//
// The created feature includes:
//   - a background
//   - one normal scenario with three steps
//   - one outline scenario with one example and three steps
func BuildTestFeature(t *testing.T) models.Feature {
	_ = "STUB: not implemented"
	return *new(models.Feature)
}

const featureContent = `Feature: eat godogs
In order to be happy
As a hungry gopher
I need to be able to eat godogs

Background:
  Given there are <begin> godogs

Scenario: Eat 5 out of 12
  When I eat 5
  Then there should be 7 remaining

Scenario Outline: Eat <dec> out of <beginning>
  When I eat <dec>
  Then there should be <remain> remaining

  Examples:
	| begin | dec | remain |
	| 12    | 5   | 7      |`

// BuildTestFeature creates a feature with rules for testing purpose.
//
// The created feature includes:
//   - a background
//   - one normal scenario with three steps
//   - one outline scenario with one example and three steps
func BuildTestFeatureWithRules(t *testing.T) models.Feature {
	_ = "STUB: not implemented"
	return *new(models.Feature)
}

const featureWithRuleContent = `Feature: eat godogs
In order to be happy
As a hungry gopher
I need to be able to eat godogs

Rule: eating godogs

Background:
  Given there are <begin> godogs

Scenario: Eat 5 out of 12
  When I eat 5
  Then there should be 7 remaining

Scenario Outline: Eat <dec> out of <beginning>
  When I eat <dec>
  Then there should be <remain> remaining

  Examples:
	| begin | dec | remain |
	| 12    | 5   | 7      |`
