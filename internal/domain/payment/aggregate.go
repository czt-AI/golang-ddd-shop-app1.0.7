package payment

import (
	"context"
	"shop-app/domain/payment"
)

type PaymentAggregate struct {
	repo payment.PaymentRepository
}

func NewPaymentAggregate(repo payment.PaymentRepository) *PaymentAggregate {
	return &PaymentAggregate{repo: repo}
}

func (a *PaymentAggregate) RecordPayment(ctx context.Context, orderID string, paymentMethod string, amount float64) error {
	return a.repo.RecordPayment(ctx, orderID, paymentMethod, amount)
}

func (a *PaymentAggregate) ConfirmPayment(ctx context.Context, transactionID string) error {
	return a.repo.ConfirmPayment(ctx, transactionID)
}

func (a *PaymentAggregate) RefundPayment(ctx context.Context, transactionID string, amount float64) error {
	return a.repo.RefundPayment(ctx, transactionID, amount)
}