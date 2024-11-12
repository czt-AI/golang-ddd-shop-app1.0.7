package payment

import (
	"context"
	"shop-app/api/v1/dto"
	"shop-app/domain/payment"
)

type PaymentUsecase interface {
	RecordPayment(ctx context.Context, req *dto.RecordPaymentRequest) error
	ConfirmPayment(ctx context.Context, transactionID string) error
	RefundPayment(ctx context.Context, transactionID string, amount float64) error
}

type paymentUsecase struct {
	repo payment.PaymentRepository
}

func NewPaymentUsecase(repo payment.PaymentRepository) *paymentUsecase {
	return &paymentUsecase{repo: repo}
}

func (u *paymentUsecase) RecordPayment(ctx context.Context, req *dto.RecordPaymentRequest) error {
	return u.repo.RecordPayment(ctx, req.OrderID, req.PaymentMethod, req.Amount)
}

func (u *paymentUsecase) ConfirmPayment(ctx context.Context, transactionID string) error {
	return u.repo.ConfirmPayment(ctx, transactionID)
}

func (u *paymentUsecase) RefundPayment(ctx context.Context, transactionID string, amount float64) error {
	return u.repo.RefundPayment(ctx, transactionID, amount)
}