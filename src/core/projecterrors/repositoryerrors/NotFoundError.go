package repositoryerrors

type NotFound struct {
	*repositoryError
}

func NewNotFoundError(message string, err error) *NotFound {
	return &NotFound{newRepositoryError(message, err)}
}
