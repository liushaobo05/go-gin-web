package model

import "time"

type Model struct {
	ID        uint       `gorm:"column:id;primary_key;auto_increment;"`
	CreatedAt time.Time  `gorm:"column:created_at;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;"`
}
