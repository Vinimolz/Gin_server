package main

import (
	db "github.com/Vinimolz/Gin_server/database"
	"github.com/Vinimolz/Gin_server/routes"
)

func main() {
	db.DatabaseConnection()
	routes.HandleRequests()
}
