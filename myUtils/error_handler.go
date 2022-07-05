package myutils

import "fmt"

type ErrorHandler interface {
	ErrorHandling(err error, error_message string)
}

type errorHandler struct {}

func NewErrorHandler() ErrorHandler {
	return &errorHandler{}
}

func (handler *errorHandler) ErrorHandling(err error, error_message string) {
	if err != nil {
		fmt.Println(error_message + ":", err)
		panic(err)
	}
}