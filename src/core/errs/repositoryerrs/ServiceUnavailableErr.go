package repositoryerrs

type ServiceUnavailable struct {
	*repositoryErr
}

func NewServiceUnavailableErr(message string, err error) *ServiceUnavailable {
	return &ServiceUnavailable{newRepositoryErr(message, err)}
}
