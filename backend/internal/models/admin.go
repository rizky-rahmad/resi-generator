package models

import "time"

type Admin struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Username  string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"` //"-" artinya tidak pernah kirim ke response
	Role      string    `gorm:"type:varchar(255);not null;default:'kerani'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
