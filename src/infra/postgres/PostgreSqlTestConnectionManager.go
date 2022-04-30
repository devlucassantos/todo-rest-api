package postgres

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type testConnectionManager struct{}

func NewPostgresTestConnectionManager() *testConnectionManager {
	return &testConnectionManager{}
}

func (c testConnectionManager) getConnection() (*sqlx.DB, error) {
	uri, err := getPostgresTestConnectionURI()
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

func (c testConnectionManager) closeConnection(connection *sqlx.DB) {
	err := connection.Close()
	if err != nil {
		log.Print("Error while closing connection: " + err.Error())
	}
}
