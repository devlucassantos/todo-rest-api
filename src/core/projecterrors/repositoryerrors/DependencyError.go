package repositoryerrors

type Dependency struct {
	*repositoryError
	affectedDependencies []string
}

func NewDependencyError(message string, err error, affectedDependencies ...string) *Dependency {
	return &Dependency{
		repositoryError:      newRepositoryError(message, err),
		affectedDependencies: affectedDependencies,
	}
}

func (e Dependency) GetAffectedDependencies() []string {
	return e.affectedDependencies
}
