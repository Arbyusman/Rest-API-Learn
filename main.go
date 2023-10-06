package main

import (
	"Rest-API/config"
	"Rest-API/database"
	"Rest-API/routes"
	m "Rest-API/usecase/middleware"
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
	// database.Drop(db)
	// database.Migrate(db)

	m.LogMiddlewares(app)
	routes.Routes(app, db)

	fmt.Println(config.AppConfig.AppPort)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.AppConfig.AppPort)))

}
