package serviceerrs

import "todo/src/core/errs/serviceerrs/msgs"

type Conflict struct {
	conflictingFields []string
}

func NewConflictErr(conflictingFields ...string) *Conflict {
	return &Conflict{conflictingFields}
}

func (err Conflict) Error() string {
	return msgs.ConflictErrMsg
}

func (err Conflict) Fields() []string {
	return err.conflictingFields
}
