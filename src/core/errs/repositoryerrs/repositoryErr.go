package repositoryerrs

type repositoryErr struct {
	friendlyMessage string
	internalError   error
}

func newRepositoryErr(friendlyMessage string, internalError error) *repositoryErr {
	return &repositoryErr{
		friendlyMessage: friendlyMessage,
		internalError:   internalError,
	}
}

func (err repositoryErr) GetFriendlyMessage() string {
	return err.friendlyMessage
}

func (err repositoryErr) PredecessorError() string {
	return err.internalError.Error()
}

func (err repositoryErr) Error() string {
	if err.internalError == nil {
		return err.friendlyMessage
	}
	return err.internalError.Error()
}
