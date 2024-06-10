package routes

import (
	"github.com/Vinimolz/Gin_server/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetAllStudents)
	r.GET("/students/:student_id", controllers.GetStudentById)
	r.DELETE("/students/:student_id", controllers.DeleteStudentById)
	r.PATCH("/students/:student_id", controllers.EditStudent)
	r.GET("/:name", controllers.HelloThere)
	r.POST("/students", controllers.CreateNewStudentAgain)
	r.Run()
}