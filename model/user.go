package model

import (
	"time"

	"gorm.io/gorm"
)

const USER_TYPE = "user"
const ADMIN_TYPE = "admin"
const ALL_TYPE = "all"

type Users struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100)" json:"name" `
	Email     string         `gorm:"type:varchar(100)" json:"email" `
	Role      string         `gorm:"type:varchar(100)" json:"role"`
	Address   string         `gorm:"type:text" json:"address"`
	Image     string         `gorm:"type:text" json:"image"`
	Phone     string         `gorm:"type:varchar(20)" json:"phone"`
	Password  string         `gorm:"type:varchar(200)" json:"password"`
	DeletedAt gorm.DeletedAt ` json:"deleted_at"`
	CreatedAt time.Time      ` json:"created_at"`
	UpdatedAt time.Time      ` json:"updated_at"`
}

// func generateUUId() {
// 	user{

// 		id := uuid.New().String()
// 	}
// }
