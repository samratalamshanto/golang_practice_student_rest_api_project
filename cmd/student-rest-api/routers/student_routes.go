package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/controller"
)

func StudentRoutes(r *gin.Engine) {
	studentRoutes := r.Group("/student")
	{
		studentRoutes.GET("/", controller.GetAllStudents)
		studentRoutes.GET("/:id", controller.GetStudent)
		studentRoutes.POST("/add", controller.AddStudent)
		studentRoutes.PUT("/:id", controller.UpdateStudent)
		studentRoutes.DELETE("/:id", controller.DeleteStudent)
		studentRoutes.DELETE("/all", controller.DeleteAllStudents)
	}

}
