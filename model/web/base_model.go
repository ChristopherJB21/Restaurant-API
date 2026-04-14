package web

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel adalah struct kustom pengganti gorm.Model
type BaseModel struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// PENTING: Pindahkan hook BeforeCreate ke BaseModel
func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate UUID baru setiap kali data di-insert
	b.ID, _ = uuid.NewV7()
	return
}
