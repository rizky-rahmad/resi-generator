package models

type InvoiceDetail struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	InvoiceID uint    `gorm:"not null;index" json:"invoice_id"`
	ItemID    uint    `gorm:"not null;index" json:"item_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"type:decimal(15,2);not null" json:"price"`    // snapshot harga
	Subtotal  float64 `gorm:"type:decimal(15,2);not null" json:"subtotal"` // price * quantity

	// Relasi: GORM bisa eager load item saat query detail
	Item Item `gorm:"foreignKey:ItemID" json:"item,omitempty"`
}
