package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rizky-rahmad/resi-generator/backend/internal/database"
	"github.com/rizky-rahmad/resi-generator/backend/internal/dto"
	"github.com/rizky-rahmad/resi-generator/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// JWTClaims mendefinisikan payload yang akan disimpan di dalam token JWT.
// embed jwt.RegisteredClaims untuk dapat field standar seperti ExpiresAt.
type JWTClaims struct {
	AdminID  uint   `json:"admin_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

// AuthService menyimpan dependency yang dibutuhkan — dalam hal ini JWT secret.
type AuthService struct {
	jwtSecret string
}

func NewAuthService(jwtSecret string) *AuthService {
	return &AuthService{jwtSecret: jwtSecret}
}

// Login adalah core business logic auth:
// 1. Cari admin berdasarkan username
// 2. Verifikasi password dengan bcrypt
// 3. Generate JWT token
// Mengembalikan LoginResponse dan error.
func (s *AuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	// Step 1: Cari admin di database berdasarkan username
	var admin models.Admin
	result := database.DB.Where("username = ?", req.Username).First(&admin)
	if result.Error != nil {
		// Sengaja pakai pesan generic — jangan beritahu apakah username atau password yang salah.
		return nil, errors.New("username atau password salah")
	}

	// Step 2: Verifikasi password dengan bcrypt
	// bcrypt.CompareHashAndPassword membandingkan hash di DB dengan plaintext input.
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("username atau password salah")
	}

	// Step 3: Buat JWT token
	expiryHours := 24
	expiresAt := time.Now().Add(time.Duration(expiryHours) * time.Hour)

	// Isi payload (claims) token
	claims := JWTClaims{
		AdminID:  admin.ID,
		Username: admin.Username,
		Role:     admin.Role,
		Name:     admin.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "resi-generator", // identitas aplikasi yang issue token ini
		},
	}

	// Buat token dengan algoritma HS256 (HMAC-SHA256) — standard untuk kebanyakan API
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token dengan secret key → menghasilkan string token final
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, errors.New("gagal membuat token")
	}

	return &dto.LoginResponse{
		Token:     tokenString,
		Role:      admin.Role,
		Name:      admin.Name,
		ExpiresIn: expiryHours * 3600,
	}, nil
}
