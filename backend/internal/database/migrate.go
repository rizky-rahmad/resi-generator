package database

import (
	"log"

	"github.com/rizky-rahmad/resi-generator/backend/internal/models"
)

//migrate jalankan auto-migration untuk semua model.
//gorm akan buat tabel/kolom baru jika belum ada atau jika ada perubahan struct

func Migrate() {
	log.Println("Running DB Migration...")

	err := DB.AutoMigrate(
		&models.Admin{},
		&models.Item{},
		&models.Invoice{},
		&models.InvoiceDetail{},
	)

	if err != nil {
		log.Fatalf("FATAL: MIGRATION FAIL: %v", err)
	}

	log.Println("DB Migration Success!")

}
