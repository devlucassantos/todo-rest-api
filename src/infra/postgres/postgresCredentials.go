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
	dbName := os.Getenv("POSTGRESQL_DB")

	connectionUri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)
	return connectionUri, nil
}
