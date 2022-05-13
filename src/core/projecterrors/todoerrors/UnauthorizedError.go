package todoerrors

import (
	"todo/src/app/api/endpoints/handlers/msgs"
)

type Unauthorized struct {
	message string
}

func NewUnauthorizedError() *Unauthorized {
	return &Unauthorized{msgs.UnauthorizedError}
}

func (err Unauthorized) Error() string {
	return err.message
}
