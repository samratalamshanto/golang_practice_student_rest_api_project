package routers

import "github.com/gin-gonic/gin"

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	CourseRotes(r)
	StudentRoutes(r)
	TeacherRoutes(r)

	return r
}
