package models

type Order struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`

	OrderID   string `json:"order_id"`
	ProductID uint   `json:"product_id"`

	CustomerName  string `json:"customer_name"`
	CustomerEmail string `json:"customer_email"`
	CustomerPhone string `json:"customer_phone"`

	Amount        float64 `json:"amount"`
	PaymentStatus string  `json:"payment_status"`
}

type CheckoutRequest struct {
	OrderID string `json:"order_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Amount  int64  `json:"amount"`
}
