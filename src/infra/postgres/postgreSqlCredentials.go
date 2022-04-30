package postgres

import (
	"fmt"
	"os"
)

func getPostgresConnectionURI() (string, error) {
	user := os.Getenv("POSTGRESQL_USER")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	host := os.Getenv("POSTGRESQL_HOST")
	port := os.Getenv("POSTGRESQL_PORT")
	dbName := os.Getenv("POSTGRESQL_NAME")

	connectionUri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)
	return connectionUri, nil
}

func getPostgresTestConnectionURI() (string, error) {
	user := os.Getenv("POSTGRESQL_TlEST_USER")
	password := os.Getenv("POSTGRESQL_TEST_PASSWORD")
	host := os.Getenv("POSTGRESQL_TEST_HOST")
	port := os.Getenv("POSTGRESQL_TEST_PORT")
	dbName := os.Getenv("POSTGRESQL_TEST_NAME")

	connectionUri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)
	return connectionUri, nil
}
