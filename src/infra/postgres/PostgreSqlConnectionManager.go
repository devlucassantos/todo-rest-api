package postgres

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type iConnectionManager interface {
	getConnection() (*sqlx.DB, error)
	closeConnection(*sqlx.DB)
}

type connectionManager struct{}

func NewPostgresConnectionManager() *connectionManager {
	return &connectionManager{}
}

func (c connectionManager) getConnection() (*sqlx.DB, error) {
	uri, err := getPostgresConnectionURI()
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Open("postgres", uri)
	if err != nil {
		log.Print("Error while accessing database: " + err.Error())
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func (c connectionManager) closeConnection(connection *sqlx.DB) {
	err := connection.Close()
	if err != nil {
		log.Print("Error while closing connection: " + err.Error())
	}
}
