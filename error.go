package main

import (
	"fmt"
)

type Error struct {
	Status        int    `json:"-"`
	Message       string `json:"message"`
	ErrorDetails  string `json:"error_details"`
	internalError *error
}

func NewError(status int, message string, internalError *error) *Error {
	err := Error{
		Status:        status,
		Message:       message,
		internalError: internalError,
	}

	if err.internalError != nil {
		err.ErrorDetails = (*err.internalError).Error()
	}

	return &err
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d] %s\n%s", e.Status, e.Message, e.ErrorDetails)
}
