package repositoryerrors

type ServiceUnavailable struct {
	*repositoryError
}

func NewServiceUnavailableError(message string, err error) *ServiceUnavailable {
	return &ServiceUnavailable{newRepositoryError(message, err)}
}
