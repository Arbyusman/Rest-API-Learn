package database

import (
	"Rest-API/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Users{},
	)

	if err != nil {
		panic(err)
	}
}

func Drop(db *gorm.DB) {
	err := db.Migrator().DropTable(
		&model.Users{},
	)
	if err != nil {
		panic(err)
	}
}
