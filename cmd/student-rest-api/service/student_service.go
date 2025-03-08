package service

import (
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/database"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/models"
)

func GetStudents() (*[]models.Student, error) {
	var students []models.Student
	res := database.DB.Find(&students)
	return &students, res.Error
}

func GetStudent(id int) (*models.Student, error) {
	var student models.Student
	res := database.DB.First(&student, id)
	return &student, res.Error
}

func CreateOrUpdateStudent(student models.Student) (*models.Student, error) {
	res := database.DB.Save(&student)
	return &student, res.Error
}

func DeleteAllStudents() error {
	res := database.DB.Where("1=1").Delete(&models.Student{})
	return res.Error
}

func DeleteStudent(id int) error {
	res := database.DB.Delete(&models.Student{}, id)
	return res.Error
}
