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
	database "go-fiber-api/db"
	"go-fiber-api/repositories"
	"go-fiber-api/routes"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect database
	database.ConnectDB()
	repositories.InitRepositories()

	app := fiber.New()

	// Middleware untuk CORS & Recover
	app.Use(cors.New())    // Izinkan akses API dari domain berbeda
	app.Use(recover.New()) // Tangani panic agar server tidak crash

	// Register Routes
	routes.AuthRoutes(app)
	routes.UserRoutes(app)
	routes.BookRoutes(app)
	routes.IktikafRoutes(app)

	// Start server
	app.Listen(":3000")
}
