package models

import (
	"context"
	"errors"
	"reflect"

	"github.com/cucumber/godog/formatters"
)

var typeOfBytes = reflect.TypeOf([]byte(nil))

// matchable errors
var (
	ErrUnmatchedStepArgumentNumber = errors.New("func expected more arguments than given")
	ErrCannotConvert               = errors.New("cannot convert argument")
	ErrUnsupportedParameterType    = errors.New("func has unsupported parameter type")
)

// StepDefinition ...
type StepDefinition struct {
	formatters.StepDefinition

	Args         []interface{}
	HandlerValue reflect.Value
	File         string
	Line         int

	// multistep related
	Nested    bool
	Undefined []string
}

var typeOfContext = reflect.TypeOf((*context.Context)(nil)).Elem()

// Run a step with the matched arguments using reflect
// Returns one of ...
// (context, error)
// (context, godog.Steps)
func (sd *StepDefinition) Run(ctx context.Context) (context.Context, interface{}) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// the error here is that the declared function has an unsupported param type - really this ought to be trapped at registration ti,e

// the problem is the function decl is not using a support slice type as the param

// Note that the step fn return types were validated at Initialise in test_context.go stepWithKeyword()

// single return value may be one of ...
// error
// context.Context
// godog.Steps

// if the single return value is a context then just return it

// return type is presumably one of nil, "error" or "Steps" so place it into second return position

// multi-value value return must be
//  (context, error) and the context value must not be nil

func (sd *StepDefinition) shouldBeString(idx int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetInternalStepDefinition ...
func (sd *StepDefinition) GetInternalStepDefinition() *formatters.StepDefinition {
	_ = "STUB: not implemented"
	return nil
}
