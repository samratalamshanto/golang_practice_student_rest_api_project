package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/controller"
)

func CourseRotes(r *gin.Engine) {
	courseRoutes := r.Group("/course")
	{
		courseRoutes.GET("/", controller.GetAllCourses)
		courseRoutes.GET("/:id", controller.GetCourse)
		courseRoutes.POST("/add", controller.AddCourse)
		courseRoutes.PUT("/:id", controller.UpdateCourse)
		courseRoutes.DELETE("/:id", controller.DeleteCourse)
		courseRoutes.DELETE("/all", controller.DeleteAllCourses)
	}
}
