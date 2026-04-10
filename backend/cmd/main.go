package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rizky-rahmad/resi-generator/backend/internal/config"
	"github.com/rizky-rahmad/resi-generator/backend/internal/database"
	"github.com/rizky-rahmad/resi-generator/backend/internal/handlers"
	"github.com/rizky-rahmad/resi-generator/backend/internal/middleware"
	"github.com/rizky-rahmad/resi-generator/backend/internal/services"
)

func main() {
	// 1. Config
	cfg := config.Load()

	// 2. Database
	database.Connect(cfg)
	database.Migrate()
	database.Seed()

	// 3. Services
	authService := services.NewAuthService(cfg.JWTSecret)
	itemService := services.NewItemService()
	invoiceService := services.NewInvoiceService()

	// 4. Handlers
	authHandler := handlers.NewAuthHandler(authService)
	itemHandler := handlers.NewItemHandler(itemService)
	invoiceHandler := handlers.NewInvoiceHandler(invoiceService)

	// 5. Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})

	// 6. Global Middleware
	app.Use(recover.New())
	app.Use(fiberLogger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// 7. Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Server is running",
		})
	})

	// 8. Routes
	api := app.Group("/api")

	// Public routes
	api.Post("/login", authHandler.Login)
	api.Get("/items", itemHandler.GetItems)

	// Protected routes
	protected := api.Group("/", middleware.Protected(cfg.JWTSecret))
	protected.Post("/invoices", invoiceHandler.CreateInvoice)

	// 9. Start server
	log.Printf("Server berjalan di port %s (env: %s)", cfg.AppPort, cfg.AppEnv)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
