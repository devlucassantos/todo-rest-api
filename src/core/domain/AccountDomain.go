package domain

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
	"os"
	"regexp"
	"strings"
	"time"
	"todo/src/core/domain/msgs"
	"todo/src/core/projecterrors/todoerrors"
)

type Account struct {
	id       int
	name     string
	email    string
	password string
	hash     string
	token    string
}

func NewValidatedAccount(id int, name, email, password, hash, token string) (*Account, *todoerrors.Validation) {
	formattedEmail := strings.ToLower(strings.TrimSpace(email))
	matched, err := regexp.MatchString("^[a-zA-Z\\d.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z\\d](?:[a-zA-Z\\d-]{0,61}[a-zA-Z\\d])?(?:\\.[a-zA-Z\\d](?:[a-zA-Z\\d-]{0,61}[a-zA-Z\\d])?)*$", formattedEmail)
	if !matched || err != nil {
		log.Error(err)
		invalidFields := todoerrors.InvalidFields{}
		invalidFields.AppendField(msgs.AccountEmail, msgs.InvalidAccountEmail)
		return nil, todoerrors.NewValidationError(msgs.InvalidAccountDetails, invalidFields)
	}

	return &Account{
		id:       id,
		name:     strings.TrimSpace(name),
		email:    formattedEmail,
		password: password,
		hash:     hash,
		token:    token,
	}, nil
}

func NewAccount(id int, name, email, password, hash, token string) *Account {
	return &Account{
		id:       id,
		name:     strings.TrimSpace(name),
		email:    strings.ToLower(strings.TrimSpace(email)),
		password: password,
		hash:     hash,
		token:    token,
	}
}

func (d Account) Id() int {
	return d.id
}

func (d Account) Name() string {
	return d.name
}

func (d Account) Email() string {
	return d.email
}

func (d Account) Password() string {
	return d.password
}

func (d *Account) SetPassword(password string) {
	d.password = password
}

func (d Account) Hash() string {
	return d.hash
}

func (d *Account) SetHash(hash string) {
	d.hash = hash
}

func (d Account) Token() string {
	return d.token
}

func (d *Account) GenerateToken() (string, error) {
	claims := d.buildClaims()
	secretKey := os.Getenv("SERVER_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Error(err)
		return "", err
	}

	d.token = signedToken

	return signedToken, nil
}

func (d Account) buildClaims() *jwt.MapClaims {
	now := time.Now().Unix()
	exp := time.Now().Add(time.Minute * 60).Unix()

	return &jwt.MapClaims{
		"exp":   exp,
		"iat":   now,
		"typ":   "Bearer",
		"iss":   fmt.Sprintf("https://%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT")),
		"email": d.email,
	}
}
