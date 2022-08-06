package ports

import "github.com/huseyinbabal/microservices/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
	GetOrder(id int64) (domain.Order, error)
}
