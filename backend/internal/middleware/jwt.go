package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rizky-rahmad/resi-generator/backend/internal/services"
)

// Protected adalah JWT middleware factory.
// Menerima jwtSecret sebagai parameter agar tidak ada global state.
// Mengembalikan fiber.Handler yang bisa dipasang ke route manapun.
func Protected(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Step 1: Ambil header Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Token tidak ditemukan",
			})
		}

		// Step 2: Pastikan format header adalah "Bearer <token>"
		// Split hasilnya: ["Bearer", "eyJhbGci..."]
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Format token tidak valid. Gunakan: Bearer <token>",
			})
		}

		tokenString := parts[1]

		// Step 3: Parse dan verifikasi token
		// jwt.ParseWithClaims akan:
		// - decode token
		// - verifikasi signature menggunakan secret
		// - cek apakah token sudah expired
		token, err := jwt.ParseWithClaims(
			tokenString,
			&services.JWTClaims{}, // struct claims yang kita definisikan di auth_service
			func(token *jwt.Token) (interface{}, error) {
				// Validasi algoritma — pastikan algoritma yang dipakai adalah HS256
				// Ini mencegah "algorithm confusion attack" (none algorithm attack)
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fiber.NewError(
						fiber.StatusUnauthorized,
						"Algoritma token tidak valid",
					)
				}
				return []byte(jwtSecret), nil
			},
		)

		// Step 4: Handle error parsing
		if err != nil {
			// jwt library memberikan error spesifik yang bisa kita bedakan
			if strings.Contains(err.Error(), "expired") {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"success": false,
					"message": "Token sudah expired, silakan login ulang",
				})
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Token tidak valid",
			})
		}

		// Step 5: Ambil claims dari token yang sudah terverifikasi
		claims, ok := token.Claims.(*services.JWTClaims)
		if !ok || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Token tidak valid",
			})
		}

		// Step 6: Simpan claims ke Locals agar bisa diakses di handler
		// Key "user" adalah konvensi — bisa diganti apapun asal konsisten
		c.Locals("user", claims)

		// Step 7: Lanjutkan ke handler berikutnya
		return c.Next()
	}
}

// GetClaims adalah helper untuk mengambil claims dari Locals di handler.
// Dengan helper ini, handler tidak perlu type assertion manual setiap saat.
func GetClaims(c *fiber.Ctx) *services.JWTClaims {
	claims, ok := c.Locals("user").(*services.JWTClaims)
	if !ok {
		return nil
	}
	return claims
}
