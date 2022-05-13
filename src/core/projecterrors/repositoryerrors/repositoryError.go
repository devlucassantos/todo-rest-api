package repositoryerrors

type repositoryError struct {
	friendlyMessage string
	internalError   error
}

func newRepositoryError(friendlyMessage string, internalError error) *repositoryError {
	return &repositoryError{
		friendlyMessage: friendlyMessage,
		internalError:   internalError,
	}
}

func (err repositoryError) GetFriendlyMessage() string {
	return err.friendlyMessage
}

func (err repositoryError) PredecessorError() string {
	return err.internalError.Error()
}

func (err repositoryError) Error() string {
	if err.internalError == nil {
		return err.friendlyMessage
	}
	return err.internalError.Error()
}
