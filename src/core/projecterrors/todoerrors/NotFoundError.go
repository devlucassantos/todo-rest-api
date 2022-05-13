package todoerrors

import "todo/src/core/projecterrors/todoerrors/msgs"

type NotFound struct {
	message string
}

func NewNotFoundError() *NotFound {
	return &NotFound{msgs.NotFoundError}
}

func (err NotFound) Error() string {
	return err.message
}
