package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        string         `gorm:"primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string         `gorm:"type:varchar(100)" json:"name" `
	Email     string         `gorm:"type:varchar(100)" json:"email" `
	Alamat    string         `gorm:"type:text" json:"alamat"`
	NoHp      string         `gorm:"type:varchar(20)" json:"nomor_handphone"`
	Password  string         `gorm:"type:varchar" json:"password"`
	DeletedAt gorm.DeletedAt ` json:"deleted_at"`
	CreatedAt time.Time      ` json:"created_at"`
	UpdatedAt time.Time      ` json:"updated_at"`
}
