package models

type CommonFields struct {
	CreatedByName string `json:"created_by_name" gorm:"column:created_by_name; size:100"`
	CreatedBy     int    `json:"created_by" gorm:"column:created_by;"`

	UpdatedByName string `json:"updated_by_name" gorm:"column:updated_by_name; size:100"`
	UpdatedBy     int    `json:"updated_by" gorm:"column:updated_by;"`

	DeletedByName string `json:"deleted_by_name" gorm:"column:deleted_by_name; size:100"`
	DeletedBy     int    `json:"deleted_by" gorm:"column:deleted_by;"`

	Active bool `json:"active" gorm:"column:active; default:true"`
}
