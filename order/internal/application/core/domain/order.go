package domain

import (
	"time"
)

type Order struct {
	ID         int64       `json:"id"`
	CustomerID int64       `json:"customer_id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  int64       `json:"created_at"`
}

type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

func NewOrder(customerId int64, orderItems []OrderItem) Order {
	return Order{
		CreatedAt:  time.Now().Unix(),
		Status:     "Pending",
		CustomerID: customerId,
		OrderItems: orderItems,
	}
}
