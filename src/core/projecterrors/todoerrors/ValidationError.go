package todoerrors

import (
	"github.com/labstack/gommon/log"
	"strings"
	"todo/src/core/projecterrors/todoerrors/msgs"
)

type Validation struct {
	message       string
	invalidFields *InvalidFields
}

func NewValidationError(message string, fields InvalidFields) *Validation {
	return &Validation{
		message:       message,
		invalidFields: &fields,
	}
}

func (err Validation) Error() string {
	return err.message
}

func (err Validation) InvalidFields() *InvalidFields {
	return err.invalidFields
}

type InvalidFields struct {
	fields []invalidField
}

func (f InvalidFields) Fields() []invalidField {
	return f.fields
}

func (f InvalidFields) HasInvalidFields() bool {
	return len(f.fields) > 0
}

func (f *InvalidFields) AppendField(name, description string) {
	if strings.Trim(name, " ") == "" {
		log.Info(msgs.EmptyFieldName)
		return
	}

	if strings.Trim(description, " ") == "" {
		description = msgs.DefaultInvalidField
	}

	f.fields = append(f.fields, invalidField{name, description})
}

type invalidField struct {
	name, description string
}

func (i invalidField) Name() string {
	return i.name
}

func (i invalidField) Description() string {
	return i.description
}
