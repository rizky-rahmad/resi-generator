package dto

// ItemResponse adalah struct untuk response data item.
// Kita tidak return seluruh model — hanya field yang dibutuhkan frontend.
type ItemResponse struct {
	ID    uint    `json:"id"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
