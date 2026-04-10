package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky-rahmad/resi-generator/backend/internal/services"
)

type ItemHandler struct {
	itemService *services.ItemService
}

func NewItemHandler(itemService *services.ItemService) *ItemHandler {
	return &ItemHandler{itemService: itemService}
}

// GetItems menangani GET /api/items?code={kode}
func (h *ItemHandler) GetItems(c *fiber.Ctx) error {
	// Ambil query parameter "code"
	// Kalau tidak ada, default ke string kosong → return semua item
	code := c.Query("code", "")

	items, err := h.itemService.GetItemsByCode(code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mengambil data item",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Data item berhasil diambil",
		"data":    items,
	})
}
