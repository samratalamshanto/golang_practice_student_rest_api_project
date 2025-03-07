package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/controller"
)

func TeacherRoutes(r *gin.Engine) {

	teacherRoutes := r.Group("/teacher")
	{
		teacherRoutes.GET("/", controller.GetAllTeachers)
		teacherRoutes.GET("/:id", controller.GetTeacher)
		teacherRoutes.POST("/add", controller.AddTeacher)
		teacherRoutes.PUT("/:id", controller.UpdateTeacher)
		teacherRoutes.DELETE("/:id", controller.DeleteTeacher)
		teacherRoutes.DELETE("/all", controller.DeleteAllTeachers)
	}
}
