package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky-rahmad/resi-generator/backend/internal/dto"
	"github.com/rizky-rahmad/resi-generator/backend/internal/services"
)

// AuthHandler menyimpan dependency ke AuthService.
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler adalah constructor AuthHandler.
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login menangani POST /api/login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	// Parse body request ke struct LoginRequest
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Format request tidak valid",
		})
	}

	// Validasi manual — field wajib tidak boleh kosong
	// (di step selanjutnya bisa diganti dengan library validator)
	if req.Username == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Username dan password wajib diisi",
		})
	}

	// Delegasikan ke service — handler tidak tahu cara kerja JWT atau bcrypt
	response, err := h.authService.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Login berhasil",
		"data":    response,
	})
}
