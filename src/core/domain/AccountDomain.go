package domain

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
	"os"
	"time"
)

type Account struct {
	id       int
	name     string
	email    string
	password string
	hash     string
	token    string
}

func NewAccount(id int, name, email, password, hash, token string) *Account {
	return &Account{
		id:       id,
		name:     name,
		email:    email,
		password: password,
		hash:     hash,
		token:    token,
	}
}

func (a Account) Id() int {
	return a.id
}

func (a Account) Name() string {
	return a.name
}

func (a Account) Email() string {
	return a.email
}

func (a Account) Password() string {
	return a.password
}

func (a *Account) SetPassword(password string) {
	a.password = password
}

func (a Account) Hash() string {
	return a.hash
}

func (a *Account) SetHash(hash string) {
	a.hash = hash
}

func (a Account) Token() string {
	return a.token
}

func (a *Account) GenerateToken() (string, error) {
	claims := a.buildClaims()
	secretKey := os.Getenv("SERVER_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Error(err)
		return "", err
	}

	a.token = signedToken

	return signedToken, nil
}

func (a Account) buildClaims() *jwt.MapClaims {
	now := time.Now().Unix()
	exp := time.Now().Add(time.Minute * 60).Unix()

	return &jwt.MapClaims{
		"exp":   exp,
		"iat":   now,
		"typ":   "Bearer",
		"iss":   fmt.Sprintf("https://%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT")),
		"email": a.email,
	}
}
