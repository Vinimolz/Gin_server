package routes

import (
	"github.com/Vinimolz/Gin_server/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetAllStudents)
	r.GET("/:name", controllers.HelloThere)
	r.Run()
}