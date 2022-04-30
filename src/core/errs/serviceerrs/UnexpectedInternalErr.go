package serviceerrs

type UnexpectedInternal struct {
	message string
}

func NewUnexpectedInternalErr(message string) *UnexpectedInternal {
	return &UnexpectedInternal{message}
}

func (err UnexpectedInternal) Error() string {
	return err.message
}
