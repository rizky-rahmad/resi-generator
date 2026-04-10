package dto

import "time"

// InvoiceItemRequest adalah satu baris item yang dikirim frontend.
// Perhatikan: kita TIDAK terima price/subtotal dari frontend — Zero-Trust.
// Backend yang akan query harga asli dari DB.
type InvoiceItemRequest struct {
	ItemID   uint `json:"item_id" validate:"required"`
	Quantity int  `json:"quantity" validate:"required,min=1"`
}

// CreateInvoiceRequest adalah full payload dari frontend saat submit invoice.
type CreateInvoiceRequest struct {
	SenderName      string               `json:"sender_name" validate:"required"`
	SenderAddress   string               `json:"sender_address" validate:"required"`
	ReceiverName    string               `json:"receiver_name" validate:"required"`
	ReceiverAddress string               `json:"receiver_address" validate:"required"`
	Items           []InvoiceItemRequest `json:"items" validate:"required,min=1"`
}

// InvoiceDetailResponse adalah response untuk satu baris detail invoice.
type InvoiceDetailResponse struct {
	ID       uint    `json:"id"`
	ItemID   uint    `json:"item_id"`
	ItemCode string  `json:"item_code"`
	ItemName string  `json:"item_name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Subtotal float64 `json:"subtotal"`
}

// InvoiceResponse adalah full response setelah invoice berhasil dibuat.
type InvoiceResponse struct {
	ID              uint                    `json:"id"`
	InvoiceNumber   string                  `json:"invoice_number"`
	SenderName      string                  `json:"sender_name"`
	SenderAddress   string                  `json:"sender_address"`
	ReceiverName    string                  `json:"receiver_name"`
	ReceiverAddress string                  `json:"receiver_address"`
	TotalAmount     float64                 `json:"total_amount"`
	CreatedBy       uint                    `json:"created_by"`
	CreatedByName   string                  `json:"created_by_name"`
	CreatedAt       time.Time               `json:"created_at"`
	Details         []InvoiceDetailResponse `json:"details"`
}
