package postgres

import (
	"fmt"
	"os"
)

func getPostgresConnectionURI() (string, error) {
	user := os.Getenv("POSTGRES_SERVICE_USER")
	password := os.Getenv("POSTGRES_SERVICE_USER_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")

	connectionUri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)
	return connectionUri, nil
}
