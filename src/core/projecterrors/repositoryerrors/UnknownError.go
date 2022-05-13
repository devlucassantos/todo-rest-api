package repositoryerrors

import "todo/src/core/projecterrors/repositoryerrors/msgs"

type Unknown struct {
	*repositoryError
}

func NewUnknownError(err error) *Unknown {
	return &Unknown{newRepositoryError(msgs.UnexpectedInternalError, err)}
}
