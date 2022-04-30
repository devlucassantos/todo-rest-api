package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"todo/src/app/api/endpoints/dto/response"
	"todo/src/core/errs/handlererrs/msgs"
	"todo/src/core/errs/serviceerrs"
)

func handleServiceErrors(c echo.Context, err error) error {
	fmt.Println(reflect.TypeOf(err).String())
	switch castedErr := err.(type) {
	case *serviceerrs.Conflict:
		return writeConflictErr(c, *castedErr)
	case *serviceerrs.UnexpectedInternal:
		return writeUnexpectedInternalErr(c, *castedErr)
	case *serviceerrs.MissingInfo:
		return writeMissingInfoErr(c, *castedErr)
	default:
		return writeUnexpectedErr(c, msgs.UnexpectedInternalErr)
	}
}

func writeUnexpectedErr(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusInternalServerError, message)
}

func writeUnexpectedInternalErr(c echo.Context, err serviceerrs.UnexpectedInternal) error {
	unexpectedInternalErr := &response.GenericErrResponse{Message: err.Error()}
	return c.JSON(http.StatusInternalServerError, *unexpectedInternalErr)
}

func writeValidationErr(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusBadRequest, err)
}

func writeConflictErr(c echo.Context, err serviceerrs.Conflict) error {
	conflictFields := err.Fields()
	conflictErr := &response.GenericErrResponse{Message: err.Error(), Conflicts: conflictFields}
	return c.JSON(http.StatusConflict, *conflictErr)
}

func writeMissingInfoErr(c echo.Context, err serviceerrs.MissingInfo) error {
	missingInfoErr := &response.GenericErrResponse{Message: err.Error()}
	return c.JSON(http.StatusBadRequest, *missingInfoErr)
}

func writeCreatedResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusCreated, data)
}
