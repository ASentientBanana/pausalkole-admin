package models

import (
	"github.com/google/uuid"
)

type Users struct {
	BaseModel
	ID       uuid.UUID `gorm:"primaryKey;type:uuid"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`

	//Files    []Files   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
