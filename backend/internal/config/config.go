package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//config untuk simpan semua konfigurasi app dari env

type Config struct {
	AppPort string
	AppEnv  string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	DBDSN string

	JWTSecret      string
	JWTExpiryHours int
}

// helper getEnv = ambil env var, kalau tidak ada, pakai defaultvalue
// patter standar di go untuk env variabel dengan fallback
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// load untuk membaca .env
// dipanggil sekali ketika main running
func Load() *Config {

	//Di docker, env var sudah terinject langsung oleh docker compose
	//godotenv.Load() hanya berguna saat development local
	//pakai _ agar tidak error meskipun tidak ada.env
	_ = godotenv.Load()

	cfg := &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		AppEnv:  getEnv("APP_ENV", "development"),

		DBHost:     getEnv("DB_Host", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "resi_generator"),

		JWTSecret:      getEnv("JWT_SECRET", ""),
		JWTExpiryHours: 24,
	}

	//validasi: wajib ada JWT_SECRET
	if cfg.JWTSecret == "" {
		log.Fatal("FATAL: JWT_SECRET TIDAK BOLEH KOSONG")
	}

	//build DSN(connection string) dari komponen DB yang sudah di read
	cfg.DBDSN = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	return cfg

}
