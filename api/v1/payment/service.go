package v1

import (
	"context"
	"shop-app/api/v1/dto"
	"shop-app/domain/payment"
)

type PaymentService interface {
	RecordPayment(ctx context.Context, req *dto.RecordPaymentRequest) error
	ConfirmPayment(ctx context.Context, transactionID string) error
	RefundPayment(ctx context.Context, transactionID string, amount float64) error
}

type paymentService struct {
	usecase payment.PaymentUsecase
}

func NewPaymentService(usecase payment.PaymentUsecase) *paymentService {
	return &paymentService{usecase: usecase}
}

func (s *paymentService) RecordPayment(ctx context.Context, req *dto.RecordPaymentRequest) error {
	return s.usecase.RecordPayment(ctx, req)
}

func (s *paymentService) ConfirmPayment(ctx context.Context, transactionID string) error {
	return s.usecase.ConfirmPayment(ctx, transactionID)
}

func (s *paymentService) RefundPayment(ctx context.Context, transactionID string, amount float64) error {
	return s.usecase.RefundPayment(ctx, transactionID, amount)
}