package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/rizky-rahmad/resi-generator/backend/internal/database"
	"github.com/rizky-rahmad/resi-generator/backend/internal/dto"
	"github.com/rizky-rahmad/resi-generator/backend/internal/models"
	"gorm.io/gorm"
)

type InvoiceService struct{}

func NewInvoiceService() *InvoiceService {
	return &InvoiceService{}
}

// CreateInvoice adalah core business logic pembuatan invoice.
// Mengimplementasikan dua requirement utama:
// 1. Zero-Trust: harga diambil dari DB, bukan dari request frontend
// 2. ACID Transaction: insert header + detail dalam satu transaksi atomik
func (s *InvoiceService) CreateInvoice(
	req dto.CreateInvoiceRequest,
	adminID uint,
	adminName string,
) (*dto.InvoiceResponse, error) {

	// === ZERO-TRUST VALIDATION ===
	// Kumpulkan semua item_id dari request
	itemIDs := make([]uint, len(req.Items))
	for i, item := range req.Items {
		itemIDs[i] = item.ItemID
	}

	// Query SEMUA item yang dibutuhkan dalam SATU query (efisien)
	var masterItems []models.Item
	result := database.DB.Where("id IN ?", itemIDs).Find(&masterItems)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data master item")
	}

	// Buat map item_id → item untuk lookup O(1)
	itemMap := make(map[uint]models.Item)
	for _, item := range masterItems {
		itemMap[item.ID] = item
	}

	// Validasi: pastikan semua item yang diminta ada di DB
	for _, reqItem := range req.Items {
		if _, exists := itemMap[reqItem.ItemID]; !exists {
			return nil, fmt.Errorf("item dengan ID %d tidak ditemukan", reqItem.ItemID)
		}
	}

	// === HITUNG ULANG TOTAL (Zero-Trust) ===
	// Backend yang hitung — tidak percaya angka dari frontend
	var grandTotal float64
	detailsData := make([]models.InvoiceDetail, len(req.Items))

	for i, reqItem := range req.Items {
		masterItem := itemMap[reqItem.ItemID]
		subtotal := masterItem.Price * float64(reqItem.Quantity)
		grandTotal += subtotal

		detailsData[i] = models.InvoiceDetail{
			ItemID:   reqItem.ItemID,
			Quantity: reqItem.Quantity,
			Price:    masterItem.Price, // harga snapshot dari DB, bukan dari request
			Subtotal: subtotal,
		}
	}

	// === GENERATE INVOICE NUMBER ===
	// Format: INV-20260410-000001
	// Timestamp + counter untuk memastikan uniqueness
	invoiceNumber := generateInvoiceNumber()

	// === DB TRANSACTION (ACID) ===
	// db.Transaction() akan otomatis:
	// - COMMIT jika function return nil
	// - ROLLBACK jika function return error
	var invoiceResponse *dto.InvoiceResponse

	err := database.DB.Transaction(func(tx *gorm.DB) error {

		// Step 1: Insert header invoice
		invoice := models.Invoice{
			InvoiceNumber:   invoiceNumber,
			SenderName:      req.SenderName,
			SenderAddress:   req.SenderAddress,
			ReceiverName:    req.ReceiverName,
			ReceiverAddress: req.ReceiverAddress,
			TotalAmount:     grandTotal,
			CreatedBy:       adminID,
		}

		if err := tx.Create(&invoice).Error; err != nil {
			// Return error → otomatis trigger ROLLBACK
			return fmt.Errorf("gagal menyimpan header invoice: %w", err)
		}

		// Step 2: Set invoice_id ke semua detail, lalu insert
		for i := range detailsData {
			detailsData[i].InvoiceID = invoice.ID
		}

		if err := tx.Create(&detailsData).Error; err != nil {
			// Return error → otomatis trigger ROLLBACK
			// Header yang sudah di-insert di Step 1 juga ikut di-rollback
			return fmt.Errorf("gagal menyimpan detail invoice: %w", err)
		}

		// Step 3: Bangun response dari data yang sudah disimpan
		detailResponses := make([]dto.InvoiceDetailResponse, len(detailsData))
		for i, detail := range detailsData {
			masterItem := itemMap[detail.ItemID]
			detailResponses[i] = dto.InvoiceDetailResponse{
				ID:       detail.ID,
				ItemID:   detail.ItemID,
				ItemCode: masterItem.Code,
				ItemName: masterItem.Name,
				Quantity: detail.Quantity,
				Price:    detail.Price,
				Subtotal: detail.Subtotal,
			}
		}

		invoiceResponse = &dto.InvoiceResponse{
			ID:              invoice.ID,
			InvoiceNumber:   invoice.InvoiceNumber,
			SenderName:      invoice.SenderName,
			SenderAddress:   invoice.SenderAddress,
			ReceiverName:    invoice.ReceiverName,
			ReceiverAddress: invoice.ReceiverAddress,
			TotalAmount:     invoice.TotalAmount,
			CreatedBy:       invoice.CreatedBy,
			CreatedByName:   adminName,
			CreatedAt:       invoice.CreatedAt,
			Details:         detailResponses,
		}

		// Return nil → otomatis COMMIT
		return nil
	})

	if err != nil {
		return nil, err
	}

	return invoiceResponse, nil
}

// generateInvoiceNumber membuat nomor invoice unik.
// Format: INV-YYYYMMDD-HHMMSS-mmm (sampai milidetik untuk menghindari collision)
func generateInvoiceNumber() string {
	now := time.Now()
	return fmt.Sprintf("INV-%s-%s",
		now.Format("20060102"),
		now.Format("150405")+fmt.Sprintf("%03d", now.Nanosecond()/1e6),
	)
}
