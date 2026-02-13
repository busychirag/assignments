package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel replaces gorm.Model with uint64 ID to prevent overflow.
type BaseModel struct {
	ID        uint64         `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
