package repositoryerrs

type Duplicated struct {
	*repositoryErr
	identifiers []string
}

func NewDuplicatedErr(message string, err error, identifiers ...string) *Duplicated {
	return &Duplicated{newRepositoryErr(message, err), identifiers}
}

func (err Duplicated) Identifiers() []string {
	return err.identifiers
}
