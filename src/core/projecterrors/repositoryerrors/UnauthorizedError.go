package repositoryerrors

type Unauthorized struct {
	*repositoryError
}

func NewUnauthorizedError(message string, err error) *Unauthorized {
	return &Unauthorized{newRepositoryError(message, err)}
}
