package services

import "todo/src/core/domain"

type IAuth interface {
	SignUp(domain.Account) (*int, *string, error)
	SignIn(account domain.Account) (*string, error)
}
