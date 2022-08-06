package ports

import "github.com/huseyinbabal/microservices/order/internal/application/core/domain"

type DBPort interface {
	Get(id int64) (domain.Order, error)
	Save(*domain.Order) error
}
