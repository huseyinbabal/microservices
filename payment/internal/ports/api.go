package ports

import "github.com/huseyinbabal/microservices/payment/internal/application/core/domain"

type APIPort interface {
	Charge(payment domain.Payment) (domain.Payment, error)
}
