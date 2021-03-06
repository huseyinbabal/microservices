package grpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/huseyinbabal/microservices-proto/golang/payment"
	"github.com/huseyinbabal/microservices/payment/internal/application/core/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.Charge(newPayment)
	err = errors.New("invalid billing address")
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &payment.CreatePaymentResponse{PaymentId: result.ID}, nil
}
