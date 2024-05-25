package main

import (
	"go-ddd-practice/pkg/controller"
	"go-ddd-practice/pkg/infrastracture/db"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	db.DbInit()
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	routes := controller.InitRouter()
	routes.Run(os.Getenv("SERVER_PORT"))
}
