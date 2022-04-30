package repositoryerrs

import "todo/src/core/errs/repositoryerrs/msgs"

type Unknown struct {
	*repositoryErr
}

func NewUnknownErr(err error) *Unknown {
	return &Unknown{newRepositoryErr(msgs.UnexpectedInternalErr, err)}
}
