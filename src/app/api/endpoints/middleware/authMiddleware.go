package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"os"
	"strings"
	"todo/src/app/api/endpoints/handlers"
	"todo/src/app/api/endpoints/handlers/msgs"
	"todo/src/core/projecterrors/todoerrors"
)

type authMiddleware struct{}

func NewAuthMiddleware() *authMiddleware {
	return &authMiddleware{}
}

func (m authMiddleware) Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		userId := ctx.Param("userId")
		token, err := m.getToken(authHeader)
		if err != nil {
			log.Error(err)
			return handlers.WriteUnauthorizedError(ctx, err.Error())
		}

		secretKey := os.Getenv("SERVER_SECRET")
		newToken, err := jwt.Parse(
			token,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			},
		)
		if err != nil {
			log.Error(err)
			return handlers.WriteUnauthorizedError(ctx, msgs.UnauthorizedError)
		}

		if !newToken.Valid {
			return handlers.WriteUnauthorizedError(ctx, msgs.UnauthorizedError)
		}

		claims := newToken.Claims.(jwt.MapClaims)
		if fmt.Sprint(claims["id"]) != userId {
			return handlers.WriteForbiddenError(ctx, msgs.ForbiddenError)
		}

		return next(ctx)
	}
}

func (m authMiddleware) getToken(authHeader string) (string, error) {
	splitAuthHeader := strings.Split(authHeader, "Bearer")
	if len(splitAuthHeader) < 2 {
		return "", todoerrors.NewUnauthorizedError()
	}

	token := strings.TrimSpace(splitAuthHeader[1])
	if token == "" {
		return "", todoerrors.NewUnauthorizedError()
	}

	return token, nil
}
