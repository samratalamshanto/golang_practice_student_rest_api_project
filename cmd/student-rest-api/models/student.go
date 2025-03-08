package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null; size:100;"`
	Age     int    `json:"age" gorm:"type:int"`
	Roll    int    `json:"roll" gorm:"type:int"`
	Class   int    `json:"class" gorm:"type:int"`
	Section string `json:"section" gorm:"size:10"`
	Email   string `json:"email" gorm:"unique; size:255"`
	CommonFields
}

func (Student) TableName() string {
	return "pp_student"
}
