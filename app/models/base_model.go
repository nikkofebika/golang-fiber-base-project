package models

type AuditCreatedModel struct {
	CreatedByID *uint `gorm:"column:created_by_id;index;default:null"`
}

type AuditUpdatedModel struct {
	UpdatedByID *uint `gorm:"column:updated_by_id;index;default:null"`
}

type AuditDeletedModel struct {
	DeletedByID *uint `gorm:"column:deleted_by_id;index;default:null"`
}
