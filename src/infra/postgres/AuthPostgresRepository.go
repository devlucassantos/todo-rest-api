package postgres

import (
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"todo/src/core/domain"
	"todo/src/core/projecterrors/repositoryerrors"
	"todo/src/infra/postgres/dto"
	"todo/src/infra/postgres/msgs"
	"todo/src/infra/postgres/query"

	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Auth struct {
	iConnectionManager
}

func NewAuthPostgresRepository(connectionManager iConnectionManager) *Auth {
	return &Auth{connectionManager}
}

func (r Auth) SignUp(account domain.Account) (*domain.Account, error) {
	connection, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return nil, repositoryerrors.NewServiceUnavailableError(msgs.ConnectionError, err)
	}
	defer r.closeConnection(connection)

	var id int
	err = connection.QueryRow(query.Auth().Insert(), dto.Account().Insert(account)...).Scan(&id)
	if err != nil {
		log.Error(err)
		return nil, r.handlePostgresError(err)
	}

	account = *domain.NewAccount(
		id,
		account.Name(),
		account.Email(),
		"",
		"",
	)

	return &account, nil
}

func (r Auth) SignIn(account domain.Account) (*domain.Account, error) {
	conn, err := r.getConnection()
	if err != nil {
		log.Error(err)
		return nil, repositoryerrors.NewServiceUnavailableError(msgs.ConnectionError, err)
	}
	defer r.closeConnection(conn)

	accountDto := dto.Account().Select().ByEmail()
	err = conn.Get(&accountDto, query.Auth().Select().ByEmail(), account.Email())
	if err != nil {
		return nil, repositoryerrors.NewUnauthorizedError(msgs.InvalidAccountCredentials, err)
	}

	hashedPassword, err := hex.DecodeString(accountDto.Password)
	if err != nil {
		log.Error(err)
		return nil, repositoryerrors.NewUnknownError(err)
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(account.Password()))
	if err != nil {
		log.Error(msgs.InvalidAccountCredentials)
		return nil, repositoryerrors.NewUnauthorizedError(msgs.InvalidAccountCredentials, err)
	}

	account = *domain.NewAccount(
		accountDto.Id,
		accountDto.Name,
		accountDto.Email,
		"",
		"",
	)

	return &account, nil
}

func (r Auth) handlePostgresError(err error) error {
	errMessage := err.Error()

	if strings.Contains(errMessage, "unique") {
		return repositoryerrors.NewDuplicatedError(msgs.DuplicatedCredentials, err, msgs.Email)
	} else if strings.Contains(errMessage, "sql: no rows in result set") {
		return repositoryerrors.NewUnauthorizedError(msgs.InvalidAccountCredentials, err)
	}

	return repositoryerrors.NewUnknownError(err)
}
