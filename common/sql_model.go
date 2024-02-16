package common

import (
	"gorm.io/gorm"
	"time"
)

type SQLModel struct {
	UserCreatedId int64          `gorm:"type:INT(11);column:created_user" json:"created_user_id"`
	UserUpdatedId int64          `gorm:"type:INT(11);column:updated_user" json:"updated_user_id"`
	CreatedAt     time.Time      `gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"type:TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"type:TIMESTAMP NULL DEFAULT NULL" json:"deleted_at"`
}
