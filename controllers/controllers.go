package controllers

import (
	"github.com/Vinimolz/Gin_server/models"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	c.JSON(200, student.Students)
}

func HelloThere(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{"Hello, there!" : "Your name is: " + name})
}