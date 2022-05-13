package todoerrors

type MissingInfo struct {
	message string
}

func NewMissingInfoError(message string) *MissingInfo {
	return &MissingInfo{message}
}

func (err MissingInfo) Error() string {
	return err.message
}
