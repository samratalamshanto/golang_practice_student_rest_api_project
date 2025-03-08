package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name string `json:"name" gorm:"not null; size:100;"`
	Age  int    `json:"age" gorm:"type:int"`
	CommonFields
}

func (Teacher) TableName() string {
	return "pp_teacher"
}
