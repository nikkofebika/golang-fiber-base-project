package models

import (
	"gorm.io/gorm"
)

type UserDetail struct {
	gorm.Model
	UserId int
	// Name     string `gorm:"size:100;not null"`
	// Email    string `gorm:"size:100;uniqueIndex;not null"`
	// Password string `gorm:"size:100;not null"`
	// IsAdmin  bool   `gorm:"default:false"`
	// AuditCreatedModel
	// AuditUpdatedModel
	// AuditDeletedModel
}
