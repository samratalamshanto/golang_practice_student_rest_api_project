package models

import "gorm.io/gorm"

type Role string

const (
	Admin      Role = "admin"
	User       Role = "user"
	SuperAdmin Role = "superadmin"
)

type Users struct {
	gorm.Model
	Username string `gorm:"unique; size:100; not null" json:"username"`
	Password string `gorm:"not null; size:512" json:"password"`
	Email    string `gorm:"unique; size:255" json:"email"`
	Role     Role   `gorm:"size:50; default:'user' " json:"role"`
	CommonFields
}

func (Users) TableName() string {
	return "pp_users"
}
