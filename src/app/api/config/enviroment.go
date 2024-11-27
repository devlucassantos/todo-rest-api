package config

import (
	"fmt"
	"os"
	"todo/src/app/api/endpoints/routes"
)

func NewServer() {
	loadEnvFile()

	app := routes.LoadRoutes()

	address := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	app.Logger.Fatal(app.Start(address))
}
