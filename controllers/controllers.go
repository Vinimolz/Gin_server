package controllers

import (
	"log"
	"net/http"

	db "github.com/Vinimolz/Gin_server/database"
	student "github.com/Vinimolz/Gin_server/models"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	c.JSON(200, student.Students)
}

func HelloThere(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{"Hello, there!": "Your name is: " + name})
}

func CreateNewStudent(c *gin.Context) {
	var newStudent student.Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check for nil pointer before using db.Db
	if db.Db == nil {
		log.Printf("Nil database pointer")
		return
	}

	// Execute database operation
	if err := db.Db.Create(&newStudent).Error; err != nil {
		log.Fatalf(err.Error())
		return
	}

	c.JSON(http.StatusOK, newStudent)
}

func CreateNewStudentAgain(c *gin.Context) {
	var student student.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error:": err.Error()})
		return
	}

	db.Db.Create(&student)
	c.JSON(http.StatusOK, student)
}
