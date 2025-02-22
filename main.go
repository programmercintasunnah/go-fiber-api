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
	"fmt"
	database "go-fiber-api/db"
	"go-fiber-api/repositories"
	"go-fiber-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Initialize the database
	database.Init()
	if database.DB == nil {
		panic("Database connection failed")
	}
	fmt.Println("Database initialized successfully")
	// Initialize repositories after DB initialization
	repositories.Init() // Panggil fungsi init() untuk repositories jika diperlukan

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
