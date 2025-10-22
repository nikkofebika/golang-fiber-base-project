package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	IsAdmin  bool   `gorm:"default:false"`
	AuditCreatedModel
	AuditUpdatedModel
	AuditDeletedModel
}

// this is the way to define table name if we not follow the gorm convention
func (user *User) TableName() string {
	return "users"
}
