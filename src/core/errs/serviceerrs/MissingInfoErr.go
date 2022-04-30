package serviceerrs

type MissingInfo struct {
	message string
}

func NewMissingInfoErr(message string) *MissingInfo {
	return &MissingInfo{message}
}

func (err MissingInfo) Error() string {
	return err.message
}
