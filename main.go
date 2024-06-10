package main

import (
	db "github.com/Vinimolz/Gin_server/database"
	student "github.com/Vinimolz/Gin_server/models"
	"github.com/Vinimolz/Gin_server/routes"
)

func main() {
	db.DatabaseConnection()
	student.Students = []student.Student{
		{Name: "Vinicius", Student_id: "123445"},
		{Name: "Anna", Student_id: "987654"},
	}
	routes.HandleRequests()
}
