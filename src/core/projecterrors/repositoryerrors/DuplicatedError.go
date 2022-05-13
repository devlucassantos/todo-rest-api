package repositoryerrors

type Duplicated struct {
	*repositoryError
	identifiers []string
}

func NewDuplicatedError(message string, err error, identifiers ...string) *Duplicated {
	return &Duplicated{newRepositoryError(message, err), identifiers}
}

func (err Duplicated) Identifiers() []string {
	return err.identifiers
}
