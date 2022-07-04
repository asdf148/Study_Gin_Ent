package myutils

import "fmt"

// TODO:
// method의 generic 공부해서 만들기

type ErrorHandler interface {
	// ErrorHandling[E error](err error, error_message string)
}

type errorHandler struct {}

func (handler *errorHandler) ErrorHandling(err error, error_message string) {
	if err != nil {
		fmt.Println("failed finding user at service: %w", err)
		// return "", err
	}
}