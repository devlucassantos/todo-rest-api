package response

import "todo/src/core/domain"

type Auth struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

func NewAuth(account domain.Account) *Auth {
	return &Auth{
		Id:          account.Id(),
		Name:        account.Name(),
		Email:       account.Email(),
		AccessToken: account.Token(),
	}
}
