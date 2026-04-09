package database

import (
	"log"

	"github.com/rizky-rahmad/resi-generator/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

func Seed() {
	log.Println("Running DB Seed...")

	seedAdmins()
	seedItems()

	log.Println("DB Seed Success!")

}

func seedAdmins() {
	admins := []models.Admin{
		{
			Name:     "Super Admin",
			Username: "admin",
			Password: hashPassword("admin123"),
			Role:     "admin",
		},
		{
			Name:     "Kerani",
			Username: "kerani",
			Password: hashPassword("kerani123"),
			Role:     "kerani",
		},
	}

	for _, admin := range admins {
		result := DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&admin)

		if result.RowsAffected == 0 {
			log.Printf("   [Admin] Skipped: %s (sudah ada)", admin.Username)
			continue
		}

		if result.Error != nil {
			log.Fatalf("FATAL: Gagal seed admin '%s': %v", admin.Username, result.Error)
		}

		log.Printf("   [Admin] Created: %s (role: %s)", admin.Username, admin.Role)
	}
}

func seedItems() {
	items := []models.Item{
		{Code: "BRG-001", Name: "Laptop ASUS VivoBook", Price: 8500000},
		{Code: "BRG-002", Name: "Mouse Wireless Logitech", Price: 350000},
		{Code: "BRG-003", Name: "Keyboard Mechanical Keychron", Price: 1200000},
		{Code: "BRG-004", Name: "Monitor LG 24 inch", Price: 3200000},
		{Code: "BRG-005", Name: "SSD External Samsung 1TB", Price: 1750000},
		{Code: "BRG-006", Name: "Webcam Logitech C920", Price: 1100000},
		{Code: "BRG-007", Name: "Headset Sony WH-1000XM5", Price: 4500000},
		{Code: "BRG-008", Name: "USB Hub 7 Port Anker", Price: 450000},
	}

	for _, item := range items {
		result := DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&item)

		if result.RowsAffected == 0 {
			log.Printf("   [Item] Skipped: %s (sudah ada)", item.Code)
			continue
		}

		if result.Error != nil {
			log.Fatalf("FATAL: Gagal seed item '%s': %v", item.Code, result.Error)
		}

		log.Printf("   [Item] Created: %s - %s", item.Code, item.Name)
	}
}

// hash password
func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("GAGAL HASH PASSWORD: %v", err)

	}

	return string(hash)
}
