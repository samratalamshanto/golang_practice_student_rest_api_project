package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/models"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/service"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/utility"
	"log"
	"strconv"
)

func GetAllStudents(c *gin.Context) {
	res, err := service.GetStudents()
	if err != nil {
		log.Println("Failed to get all students. Error= " + err.Error())
		utility.ErrorResponse(c, "Failed to get all students. Error= "+err.Error(), err)
		return
	}

	utility.SuccessResponse(c, "Successfully get all students", res)
}

func GetStudent(c *gin.Context) {
	studentId := c.Param("id")
	studentIdInt, errConv := strconv.Atoi(studentId)
	if errConv != nil {
		utility.ErrorResponse(c, "studentId must be int", errConv)
		return
	}
	res, err := service.GetStudent(studentIdInt)
	if err != nil {
		utility.ErrorResponse(c, "Failed to get student. Error= "+err.Error(), err)
		return
	}

	utility.SuccessResponse(c, "Successfully get all students", res)
}

func CreateOrUpdateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		utility.ErrorResponse(c, "Bad Request. User type JSON expected.", err)
		return
	}

	res, err := service.CreateOrUpdateStudent(student)
	if err != nil {
		utility.ErrorResponse(c, "Failed to get all students. Error= "+err.Error(), err)
		return
	}

	utility.SuccessResponse(c, "Successfully create student. ", res)
}

func DeleteStudent(c *gin.Context) {
	studentId := c.Param("id")
	studentIdInt, errConv := strconv.Atoi(studentId)
	if errConv != nil {
		utility.ErrorResponse(c, "studentId must be int", errConv)
	}
	err := service.DeleteStudent(studentIdInt)
	if err != nil {
		utility.ErrorResponse(c, "Failed to delete the student. Error= "+err.Error(), err)
	}
	utility.SuccessResponse(c, "Successfully delete student", nil)

}

func DeleteAllStudents(c *gin.Context) {
	err := service.DeleteAllStudents()
	if err != nil {
		utility.ErrorResponse(c, "Failed to delete all students. Error= "+err.Error(), err)
		return
	}
	utility.SuccessResponse(c, "Successfully delete all students", nil)
}
