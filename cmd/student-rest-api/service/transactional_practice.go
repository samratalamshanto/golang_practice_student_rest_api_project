package service

import (
	"fmt"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/database"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/models"
	"gorm.io/gorm"
)

func TransactionalByManually() error {
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

func TransactionalByGorm() error {
	return database.DB.Transaction(func(tx *gorm.DB) error {

		// Set transaction-level timeout (60s) only for this transaction
		tx.Exec("SET statement_timeout = '60s'")

		// Set transaction isolation level only for this transaction
		tx.Exec("SET TRANSACTION ISOLATION LEVEL REPEATABLE READ")

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

func transactionWithPessimisticLock(db *gorm.DB) {

	err := db.Transaction(func(tx *gorm.DB) error {

		tx.Exec("SET statement_timeout = '60s'")
		tx.Exec("SET TRANSACTION ISOLATION LEVEL REPEATABLE READ")

		// Lock the row with "FOR UPDATE"--> pessimistic lock
		// SELECT * FROM student WHERE id = 1 FOR UPDATE;
		var student models.Student
		tx.Clauses(gorm.Expr("FOR UPDATE")).First(&student, 1)

		// Modify data
		student.Age += 10
		if err := tx.Save(&student).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("Transaction failed:", err)
	}
}

func transactionWithOptimisticLock(db *gorm.DB) {

	err := db.Transaction(func(tx *gorm.DB) error {

		tx.Exec("SET statement_timeout = '60s'")
		tx.Exec("SET TRANSACTION ISOLATION LEVEL REPEATABLE READ")

		// Lock the row with "version"--> pessimistic lock
		//need to add version in the struct like--> Version int `gorm:"version"`
		var student models.Student
		tx.First(&student, 1)

		//update student set age = age+10 where id=1
		db.Model(&models.Student{}).Update("age", gorm.Expr("age + ?", 10)).Update("name", "NewName").Where("id=?", student.ID)

		db.Model(&models.Student{}).
			Where("id = ?", student.ID).
			Updates(map[string]interface{}{
				"age":  gorm.Expr("age + ?", 10),
				"name": "New Name",
			})

		//SELECT * FROM users WHERE status = (SELECT status FROM orders WHERE id = 123);
		db.Where("status = ?", gorm.Expr("(SELECT status FROM orders WHERE id = ?)", 123)).Find(&models.Student{})

		// Modify data
		student.Age += 10
		if err := tx.Save(&student).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("Transaction failed:", err)
	}
}
