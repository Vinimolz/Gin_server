package main

import (
	student "github.com/Vinimolz/Gin_server/models"
	"github.com/Vinimolz/Gin_server/routes"
)

func main() {
	student.Students = []student.Student{
		{Name: "Vinicius", Id: "123445"},
		{Name: "Anna", Id: "987654"},
	}
	routes.HandleRequests()
}
