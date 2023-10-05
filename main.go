package main

import (
	"Rest-API/config"
	"Rest-API/database"
	"Rest-API/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()

	app := fiber.New()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	database.Migrate(db)

	routes.Routes(app, db)
	fmt.Println(config.AppConfig.AppPort)
	// var port = fmt.Sprintf(":%s", config.AppConfig.AppPort)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.AppConfig.AppPort)))

}
