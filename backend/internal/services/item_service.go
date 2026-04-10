package services

import (
	"github.com/rizky-rahmad/resi-generator/backend/internal/database"
	"github.com/rizky-rahmad/resi-generator/backend/internal/dto"
	"github.com/rizky-rahmad/resi-generator/backend/internal/models"
)

type ItemService struct{}

func NewItemService() *ItemService {
	return &ItemService{}
}

// GetItemsByCode mencari item berdasarkan kode.
// Menggunakan ILIKE (case-insensitive LIKE) agar pencarian lebih fleksibel.
// Contoh: query "brg" akan cocok dengan "BRG-001", "BRG-002", dst.
func (s *ItemService) GetItemsByCode(code string) ([]dto.ItemResponse, error) {
	var items []models.Item

	query := database.DB

	// Kalau ada query code, filter dengan ILIKE
	// Kalau kosong, return semua item
	if code != "" {
		// % di awal dan akhir = contains search
		// ILIKE = case-insensitive, khusus PostgreSQL
		query = query.Where("code ILIKE ?", "%"+code+"%")
	}

	// Limit 10 — mencegah response terlalu besar,
	// cukup untuk dropdown debounce di frontend
	result := query.Limit(10).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	// Convert dari []models.Item ke []dto.ItemResponse
	// Kita tidak return model langsung agar kontrak API tidak bergantung pada struktur DB
	response := make([]dto.ItemResponse, len(items))
	for i, item := range items {
		response[i] = dto.ItemResponse{
			ID:    item.ID,
			Code:  item.Code,
			Name:  item.Name,
			Price: item.Price,
		}
	}

	return response, nil
}
