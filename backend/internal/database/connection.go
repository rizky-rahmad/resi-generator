package database

import (
	"log"
	"time"

	"github.com/rizky-rahmad/resi-generator/backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// buat variabel db
// DB adalah var global yang menyimpan instance koneksi GORM
// Hanya ada satu instance di seluruh aplikasi (singleton Pattern)
var DB *gorm.DB

// buka koneksi ke PG dan konfigurasi connection pool
func Connect(cfg *config.Config) {
	var logLevel logger.LogLevel

	//tampilkan semua sql querydi terminal (untuk development)
	//hanya tampilkan erorr ketika prod
	if cfg.AppEnv == "development" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Error
	}

	var err error
	DB, err = gorm.Open(postgres.Open(cfg.DBDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		//kalau koneksi DB gagal, stop server
		log.Fatalf("FATAL: GAGAL KONEKSI KE DB: %v", err)
	}

	//ambil underlying *sql.DB dari GORM utk konfigurasi conn pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("FATAL: GAGAL MENGAMBIL INSTANCE SQL.DB: %v", err)
	}

	//konfigurasi connection pool

	//maksimal koneksi yang aktif
	sqlDB.SetMaxOpenConns(25)

	//maksimal koneksi idle yang tetap bisa siap pakai di pool
	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("DB Connection Success!")

}
