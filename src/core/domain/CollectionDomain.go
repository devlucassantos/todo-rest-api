package domain

import (
	"github.com/labstack/gommon/log"
	"strings"
	"todo/src/core/domain/msgs"
	"todo/src/core/projecterrors/todoerrors"
)

type Collection struct {
	id   int
	name string
}

func NewValidatedCollection(id int, name string) (*Collection, *todoerrors.Validation) {
	formattedName := strings.TrimSpace(name)
	if formattedName == "" {
		log.Error(msgs.CollectionName)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.CollectionName, msgs.InvalidCollectionName)
		return nil, todoerrors.NewValidationError(msgs.InvalidCollectionDetails, invalidFields)
	}

	return &Collection{
		id:   id,
		name: formattedName,
	}, nil
}

func NewCollection(id int, name string) *Collection {
	return &Collection{
		id:   id,
		name: strings.TrimSpace(name),
	}
}

func (d Collection) Id() int {
	return d.id
}

func (d Collection) Name() string {
	return d.name
}
