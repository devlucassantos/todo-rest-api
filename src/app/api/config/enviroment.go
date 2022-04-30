package config

import (
	"fmt"
	"os"
	"todo/src/app/api/endpoints/routes"
)

func NewServer() {
	loadEnvFile()

	app := routes.LoadRoutes()

	address := fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT"))
	app.Logger.Fatal(app.Start(address))
}
