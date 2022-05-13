package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
	"todo/src/app/api/endpoints/dto/response"
	"todo/src/app/api/endpoints/handlers/msgs"
	"todo/src/core/projecterrors/todoerrors"
)

func handleServiceErrors(ctx echo.Context, err error) error {
	switch castedErr := err.(type) {
	case *todoerrors.Conflict:
		return writeConflictError(ctx, *castedErr)
	case *todoerrors.Unauthorized:
		return WriteUnauthorizedError(ctx, err.Error())
	case *todoerrors.MissingInfo:
		return writeMissingInfoError(ctx, *castedErr)
	case *todoerrors.Validation:
		return writeValidationError(ctx, *castedErr)
	case *todoerrors.NotFound:
		return writeNotFoundError(ctx, err.Error())
	case *todoerrors.UnexpectedInternal:
		return writeUnexpectedInternalError(ctx, *castedErr)
	default:
		return writeUnexpectedError(ctx, msgs.UnexpectedInternalError)
	}
}

func writeConflictError(ctx echo.Context, err todoerrors.Conflict) error {
	conflictFields := err.Fields()
	conflictErr := &response.GenericErrorResponse{Message: err.Error(), Conflicts: conflictFields}
	return ctx.JSON(http.StatusConflict, *conflictErr)
}

func WriteUnauthorizedError(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusUnauthorized, response.GenericErrorResponse{Message: message})
}

func writeMissingInfoError(ctx echo.Context, err todoerrors.MissingInfo) error {
	missingInfoErr := &response.GenericErrorResponse{Message: err.Error()}
	return ctx.JSON(http.StatusBadRequest, *missingInfoErr)
}

func writeValidationError(ctx echo.Context, err todoerrors.Validation) error {
	var invalidFields todoerrors.InvalidFields

	for _, field := range err.InvalidFields().Fields() {
		invalidFields.AppendField(field.Name(), field.Description())
	}

	var invalidFieldsResponse []response.InvalidFieldsResponse

	for _, field := range invalidFields.Fields() {
		invalidField := &response.InvalidFieldsResponse{
			Name:        field.Name(),
			Description: field.Description(),
		}
		invalidFieldsResponse = append(invalidFieldsResponse, *invalidField)
	}

	validationErr := response.GenericErrorResponse{
		Message:       err.Error(),
		InvalidFields: invalidFieldsResponse,
	}

	return ctx.JSON(http.StatusBadRequest, validationErr)
}

func writeNotFoundError(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusNotFound, response.GenericErrorResponse{Message: message})
}

func writeUnexpectedError(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusInternalServerError, response.GenericErrorResponse{Message: message})
}

func writeUnexpectedInternalError(ctx echo.Context, err todoerrors.UnexpectedInternal) error {
	unexpectedInternalErr := &response.GenericErrorResponse{Message: err.Error()}
	return ctx.JSON(http.StatusInternalServerError, *unexpectedInternalErr)
}

func writeAcceptResponse(ctx echo.Context, data interface{}) error {
	if fmt.Sprint(data) == "[]" {
		return ctx.String(http.StatusOK, fmt.Sprint(data))
	}

	return ctx.JSON(http.StatusOK, data)
}

func writeCreatedResponse(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusCreated, data)
}

func writeNoContentResponse(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNoContent)
}

func convertToInt(value string, paramName string) (int, error) {
	invalidFields := todoerrors.InvalidFields{}
	if value == "" {
		errorMessage := fmt.Sprintf("Parameter not reported: %s", paramName)
		invalidFields.AppendField(paramName, errorMessage)
		return -1, todoerrors.NewValidationError(errorMessage, invalidFields)
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Error(err.Error())
		errorMessage := fmt.Sprintf("Invalid parameter: %s", paramName)
		invalidFields.AppendField(paramName, errorMessage)
		return -1, todoerrors.NewValidationError(paramName, invalidFields)
	}

	return intValue, nil
}
