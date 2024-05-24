package main

import (
	db "demo_api/Config"
	models "demo_api/Models"
	routes "demo_api/Routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	fmt.Println("wellcome here...")
	db.Connect()
	app := fiber.New()
	app.Use(app)
	routes.Setup(app)
	app.Listen(":3000")
}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Product{},
	)
}
