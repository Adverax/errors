package errors

import "context"

type ErrorHandler interface {
	HandleError(ctx context.Context, err error)
}

type DefaultErrorHandler struct{}

func (that *DefaultErrorHandler) HandleError(context.Context, error) {}

var defaultErrorHandler = &DefaultErrorHandler{}

func NewDefaultErrorHandler() ErrorHandler {
	return defaultErrorHandler
}
