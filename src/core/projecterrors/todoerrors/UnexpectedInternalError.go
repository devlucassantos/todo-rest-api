package todoerrors

type UnexpectedInternal struct {
	message string
}

func NewUnexpectedInternalError(message string) *UnexpectedInternal {
	return &UnexpectedInternal{message}
}

func (err UnexpectedInternal) Error() string {
	return err.message
}
