package infra

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID
	Balance   int
	Limit     int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
