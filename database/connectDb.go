package database

import (
	"Rest-API/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	// 	config.AppConfig.DBHost,
	// 	config.AppConfig.DBUsername,
	// 	config.AppConfig.DBPassword,
	// 	config.AppConfig.DBName,
	// 	config.AppConfig.DBPort,
	// 	config.AppConfig.SSLMode,
	// )

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.AppConfig.DBUsername,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
	if err != nil {
		return nil, err
	}
	return db, nil

}
