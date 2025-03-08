package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null; size:150"`
	Description string `json:"description" gorm:"size:500"`
	CommonFields
}

func (Course) TableName() string {
	return "pp_course"
}
