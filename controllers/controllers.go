package controllers

import (
	"log"
	"net/http"

	db "github.com/Vinimolz/Gin_server/database"
	student "github.com/Vinimolz/Gin_server/models"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	var studentSlice []student.Student
	db.Db.Find(&studentSlice)
	c.JSON(200, studentSlice)
}

func GetStudentById(c *gin.Context) {
	studentId := c.Params.ByName("student_id")
	var desiredStudent student.Student
	db.Db.First(&desiredStudent, studentId)

	if desiredStudent.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"404:": "Student was not found"})
		return
	}

	c.JSON(200, desiredStudent)
}

func DeleteStudentById(c *gin.Context) {
	studentId := c.Params.ByName("student_id")
	var deleteThisStudent student.Student
	db.Db.Where("ID = ?", studentId).Delete(&deleteThisStudent)
	c.JSON(http.StatusOK, gin.H{"Success:" : "Student was deleted from db"})
}

func EditStudent(c *gin.Context) {
	var editStudent student.Student
	studentId := c.Params.ByName("student_id")

	db.Db.First(&editStudent, studentId)

	if err := c.ShouldBindJSON(&editStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error: " : "Could not edit the student"})
		return
	}

	db.Db.Model(&editStudent).UpdateColumns(editStudent)

	c.JSON(http.StatusOK, editStudent)
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
