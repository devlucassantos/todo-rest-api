package todoerrors

import "todo/src/core/projecterrors/todoerrors/msgs"

type Conflict struct {
	conflictingFields []string
}

func NewConflictError(conflictingFields ...string) *Conflict {
	return &Conflict{conflictingFields}
}

func (err Conflict) Error() string {
	return msgs.ConflictError
}

func (err Conflict) Fields() []string {
	return err.conflictingFields
}
