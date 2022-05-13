package todoerrors

import (
	"github.com/labstack/gommon/log"
	"reflect"
	"todo/src/core/projecterrors/repositoryerrors"
	"todo/src/core/projecterrors/todoerrors/msgs"
)

func ConvertRepositoryErrorToServiceError(err error, reporter interface{}) error {
	reporterType := reflect.TypeOf(reporter).String()
	switch castedErr := err.(type) {
	case *repositoryerrors.Duplicated:
		return handleDuplicatedError(castedErr, reporterType)
	case *repositoryerrors.Dependency:
		return handleDependencyError(castedErr, reporterType)
	case *repositoryerrors.ServiceUnavailable:
		return handleServiceUnavailableError(castedErr, reporterType)
	case *repositoryerrors.NotFound:
		return handleNotFoundError()
	case *repositoryerrors.Unauthorized:
		return handleUnauthorizedError()
	default:
		return handleUnknownError(err, reporterType)
	}
}

func handleDuplicatedError(err *repositoryerrors.Duplicated, reporter string) error {
	log.Info(reporter, err.Error(), " - ", err.PredecessorError())
	return NewConflictError(err.Identifiers()...)
}

func handleDependencyError(err *repositoryerrors.Dependency, reporter string) error {
	log.Info(reporter, " - ", err.Error(), " - ", err.PredecessorError())

	invalidFields := InvalidFields{}
	affectedList := err.GetAffectedDependencies()
	for _, affectedField := range affectedList {
		invalidFields.AppendField(affectedField, msgs.FieldNotFound)
	}
	return NewValidationError(err.GetFriendlyMessage(), invalidFields)
}

func handleNotFoundError() error {
	return NewNotFoundError()
}

func handleUnauthorizedError() error {
	return NewUnauthorizedError()
}

func handleServiceUnavailableError(err *repositoryerrors.ServiceUnavailable, reporter string) error {
	log.Error(reporter, " - ", err.PredecessorError())
	return NewUnexpectedInternalError(msgs.UnexpectedInternalError)
}

func handleUnknownError(err error, reporter string) error {
	log.Warn(reporter, " - ", err.Error())
	return NewUnexpectedInternalError(msgs.UnexpectedInternalError)
}
