package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky-rahmad/resi-generator/backend/internal/dto"
	"github.com/rizky-rahmad/resi-generator/backend/internal/middleware"
	"github.com/rizky-rahmad/resi-generator/backend/internal/services"
)

type InvoiceHandler struct {
	invoiceService *services.InvoiceService
}

func NewInvoiceHandler(invoiceService *services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{invoiceService: invoiceService}
}

func (h *InvoiceHandler) CreateInvoice(c *fiber.Ctx) error {
	claims := middleware.GetClaims(c)
	if claims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}

	var req dto.CreateInvoiceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Format request tidak valid",
		})
	}

	if req.SenderName == "" || req.SenderAddress == "" ||
		req.ReceiverName == "" || req.ReceiverAddress == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Semua field pengirim dan penerima wajib diisi",
		})
	}

	if len(req.Items) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invoice harus memiliki minimal 1 item",
		})
	}

	for _, item := range req.Items {
		if item.ItemID == 0 || item.Quantity < 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Setiap item harus memiliki item_id dan quantity minimal 1",
			})
		}
	}

	invoice, err := h.invoiceService.CreateInvoice(req, claims.AdminID, claims.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Invoice berhasil dibuat",
		"data":    invoice,
	})
}
