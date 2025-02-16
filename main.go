package main

// go-fiber-api/
// │── main.go            # File utama
// │── .env               # File konfigurasi environment
// │── database/          # Koneksi database
// │   └── db.go
// │── models/            # Struktur tabel/database
// │   └── user.go
// │── routes/            # Endpoint API
// │   ├── auth.go
// │   ├── user.go
// │── middleware/        # Middleware (Auth & Security)
// │   ├── jwt.go
// │   ├── limiter.go

import (
	"go-fiber-api/database"
	"go-fiber-api/models"
	"go-fiber-api/routes"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// Connect database
	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{})

	// Register Routes
	routes.AuthRoutes(app)
	routes.UserRoutes(app)

	// Start server
	app.Listen(":3000")
}
