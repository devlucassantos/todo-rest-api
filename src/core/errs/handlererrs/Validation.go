package handlererrs

import (
	"github.com/labstack/gommon/log"
	"strings"
	"todo/src/core/errs/handlererrs/msgs"
)

type validation struct {
	message       string
	invalidFields *InvalidFields
}

func NewValidationErr(message string, fields InvalidFields) *validation {
	return &validation{
		message:       message,
		invalidFields: &fields,
	}
}

func (err validation) Error() string {
	return err.message
}

func (err validation) InvalidFields() *InvalidFields {
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
		log.Info("Name is empty! No field was added.")
		return
	}

	if strings.Trim(description, " ") == "" {
		description = msgs.DefaultInvalidFieldMsg
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
