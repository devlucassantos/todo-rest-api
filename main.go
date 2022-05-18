package main

import (
	"todo/src/app/api/config"
)

// @Title           			To Do List API
// @Version        				1.0.0
// @Description    				This repository is an example of an API made in Go with the To Do List theme.
// @Contact.name   				Lucas Santos
// @Contact.url    				https://github.com/devlucassantos
// @License.name				MIT
// @Host 						localhost:8000
// @BasePath 					/api
// @securitySchemes 			http
// @securityDefinitions.apikey 	bearerAuth
// @type 						http
// @scheme 						bearer
// @in 							header
// @name 						Authorization
// @bearerFormat 				JWT
func main() {
	config.NewServer()
}
