package models

import "time"

type Invoice struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	InvoiceNumber   string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"invoice_number"`
	SenderName      string    `gorm:"type:varchar(255);not null" json:"sender_name"`
	SenderAddress   string    `gorm:"type:text;not null" json:"sender_address"`
	ReceiverName    string    `gorm:"type:varchar(255);not null" json:"receiver_name"`
	ReceiverAddress string    `gorm:"type:text;not null" json:"receiver_address"`
	TotalAmount     float64   `gorm:"type:decimal(15,2);not null;default:0" json:"total_amount"`
	CreatedBy       uint      `gorm:"not null" json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	//Relasi
	Admin Admin `gorm:"foreignKey:CreatedBy" json:"admin,omitempty"`
	//relasi one to many
	Details []InvoiceDetail `gorm:"foreignKey:InvoiceID" json:"details,omitempty"`
}
