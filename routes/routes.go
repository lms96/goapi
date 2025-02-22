package routes

import (
	"goapi/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/students", controllers.GetAllStudents)
	r.POST("/student", controllers.CreateStudent)
	r.GET("/student/:id", controllers.GetStudent)
	r.PATCH("/student/:id", controllers.EditStudent)
	r.DELETE("/student/:id", controllers.DeleteStudent)
	r.GET("/student/cpf/:cpf", controllers.GetStudentByCPF)
	r.Run()
}
