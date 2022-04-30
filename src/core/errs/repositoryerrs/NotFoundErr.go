package repositoryerrs

type NotFound struct {
	*repositoryErr
}

func NewNotFoundErr(message string, err error) *NotFound {
	return &NotFound{newRepositoryErr(message, err)}
}
