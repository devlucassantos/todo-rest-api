package serviceerrs

import (
	"github.com/labstack/gommon/log"
	"reflect"
	"todo/src/core/errs/repositoryerrs"
	"todo/src/core/errs/repositoryerrs/msgs"
)

func ConvertRepositoryErrToServiceErr(err error, reporter interface{}) error {
	reporterType := reflect.TypeOf(reporter).String()

	switch castedErr := err.(type) {
	case *repositoryerrs.Duplicated:
		return handleDuplicatedErr(castedErr, reporterType)
	case *repositoryerrs.ServiceUnavailable:
		return handleServiceUnavailableErr(castedErr, reporterType)
	default:
		return handleUnknownErr(err, reporterType)
	}
}

func handleDuplicatedErr(err *repositoryerrs.Duplicated, reporter string) error {
	log.Info(reporter, err.Error(), " - ", err.PredecessorError())
	return NewConflictErr(err.Identifiers()...)
}

func handleServiceUnavailableErr(err *repositoryerrs.ServiceUnavailable, reporter string) error {
	log.Error(reporter, " - ", err.PredecessorError())
	return NewUnexpectedInternalErr(msgs.UnexpectedInternalErr)
}

func handleUnknownErr(err error, reporter string) error {
	log.Warn(reporter, " - ", err.Error())
	return NewUnexpectedInternalErr(msgs.UnexpectedInternalErr)
}
