package service

import (
	"fmt"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/database"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/models"
	"gorm.io/gorm"
	"log"
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

func CreateOrUpdateStudent(student []models.Student) (*[]models.Student, error) {
	res := database.DB.Save(&student) //batch insert, update for better performance
	return &student, res.Error
}

func DeleteAllStudents() error {
	res := database.DB.Where("1=1").Unscoped().Delete(&models.Student{})
	return res.Error
}

func DeleteStudent(id int) error {
	res := database.DB.Unscoped().Delete(&models.Student{}, id)
	return res.Error
}

// -------------- Raw Queries

func GetAllStudentsBySectionAndClass(classId int, section string) (*[]models.Student, error) {
	var students []models.Student

	res := database.DB.Where("class=? and section=?", classId, section).Find(&students)

	res1 := database.DB.Raw("select * from pp_students where class=? and section=?", classId, section).Scan(&students)
	fmt.Println(res1)

	return &students, res.Error
}

func DeleteAllStudentsBySectionAndClass(classId int, section string) error {
	//for update,insert,delete--> no return result
	res := database.DB.Exec("delete from pp_students where class=? and section=?", classId, section)
	log.Println("Rows affected:", res.RowsAffected) //Print Rows Affected
	return res.Error
}

func TestTransactional() error {
	tx := database.DB.Begin() //begin transaction

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var students []models.Student
	resStd := tx.Where("class=?", 1).Find(&students)
	if resStd.Error != nil {
		tx.Rollback()
		return resStd.Error
	}

	var teacher models.Teacher
	resTeacher := tx.Where("name=?", "Sam").Find(&teacher) // For transactional, should use same tx, not database.DB (create new tx)
	if resTeacher.Error != nil {
		tx.Rollback()
		return resTeacher.Error
	}

	return tx.Commit().Error
}

func TestTransactionalByGorm() error {
	return database.DB.Transaction(func(tx *gorm.DB) error {

		var students []models.Student
		resStd := tx.Where("class=?", 1).Find(&students)
		if resStd.Error != nil {
			return resStd.Error
		}

		var teacher models.Teacher
		resTeacher := tx.Where("name=?", "Sam").Find(&teacher) // For transactional, should use same tx, not database.DB (create new tx)
		if resTeacher.Error != nil {
			return resTeacher.Error
		}

		return nil
	})
}
