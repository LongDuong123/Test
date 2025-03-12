package main

import (
	"log"

	"example.com/api"
	"example.com/mysql"
	"example.com/repository"
	"example.com/service"
)

func main() {
	database, err := mysql.NewMySQL() // Connect to the database
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
		return
	}

	repo := repository.NewUserRepository(database) // Create a new repository
	serv := service.NewUserService(repo)           // Create a new service
	ctr := api.NewUserController(serv)             // Create a new controller
	router := api.Routes(ctr)                      // Create a new router
	router.Run(":8080")
}
