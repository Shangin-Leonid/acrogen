package ut /* Unit Tests*/

import (
	"errors"

	"acrogen/utils"
)

// TODO docs
type ErrorsAccumulator struct {
	errs []error
}

func NewErrorsAccumulator() ErrorsAccumulator {
	return ErrorsAccumulator{make([]errors)}
}

// TODO docs
func (ea *ErrorsAccumulator) TakeIntoAccount(err error) {
	if err != nil {
		ea.errs = append(ea.errs, err)
	}
}

// TODO docs
func (ea ErrorsAccumulator) IsNoError() bool {
	return len(ea.errs) == 0
}

// TODO docs
func AssertEq[V any](realValue, expectedValue V, failMes string) error {
	// TODO print values
	return utils.TerOp(realValue == expectedValue, nil, errors.New(failMes))
}
