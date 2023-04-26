package main

import (
	"./ "
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	host := "database-1.cpgqobnybsuv.us-east-1.rds.amazonaws.com"
	port := 5432 // Replace with your RDS port if different
	user := "postgres"
	password := "mypassword1"
	dbname := "database-1"

	userRepo, err := user.NewPostgresUserRepository(host, port, user, password, dbname)
	if err != nil {
		log.Fatalf("Failed to create user repository: %v", err)
	}

	userHandler := user.NewUserHandler(userRepo)
	userHandler.RegisterRoutes(router)

	router.Run(":8000")
}

