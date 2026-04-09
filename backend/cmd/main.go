package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rizky-rahmad/resi-generator/backend/internal/config"
	"github.com/rizky-rahmad/resi-generator/backend/internal/database"
)

func main() {
	//load konfigurasi env
	cfg := config.Load()

	//koneksi ke db
	database.Connect(cfg)

	//migration seeder
	database.Migrate()
	database.Seed()

	//init fiber
	app := fiber.New(fiber.Config{
		//pesan error saat panic
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})

	//daftarkan middleware global
	app.Use(recover.New())     //tangkap panic
	app.Use(fiberLogger.New()) //log tiap request
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	//health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "SERVER IS RUNNING",
		})
	})

	//ROUTES

	//JALANKAN SERVER
	log.Printf("Server running from port: %s", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
