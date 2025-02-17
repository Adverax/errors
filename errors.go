package errors

import (
	"errors"
	"strings"
)

// Errors is aggregator and represents compound error and contains slice of real errors.
type Errors struct {
	errors []error
}

func NewErrors(es ...error) *Errors {
	res := &Errors{
		errors: make([]error, 0),
	}
	for _, e := range es {
		res.Add(e)
	}
	return res
}

func (that *Errors) Check(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			that.Add(err)
		}
	}
	return that.ResError()
}

func (that *Errors) ResError() error {
	if that.IsEmpty() {
		return nil
	}

	return that
}

func (that *Errors) Error() string {
	if that == nil {
		return ""
	}

	if len(that.errors) == 0 {
		return ""
	}

	errStrings := make([]string, 0)
	for _, e := range that.errors {
		errStrings = append(errStrings, e.Error())
	}
	return strings.Join(errStrings, "\n")
}

func (that *Errors) Add(err error) {
	that.errors = append(that.errors, err)
}

func (that *Errors) AddErrors(errs *Errors) {
	that.errors = append(that.errors, errs.errors...)
}

func (that *Errors) IsEmpty() bool {
	return len(that.errors) == 0
}

func (that *Errors) Size() int {
	return len(that.errors)
}

func (that *Errors) Contains(err error) bool {
	for _, e := range that.errors {
		if errors.Is(e, err) {
			return true
		}
	}

	return false
}
