package payment

import (
	"context"
	"shop-app/internal/dao"
)

type PaymentRepository interface {
	RecordPayment(ctx context.Context, orderID string, paymentMethod string, amount float64) error
	ConfirmPayment(ctx context.Context, transactionID string) error
	RefundPayment(ctx context.Context, transactionID string, amount float64) error
}

type paymentRepository struct {
	db *dao.DB
}

func NewPaymentRepository(db *dao.DB) *paymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) RecordPayment(ctx context.Context, orderID string, paymentMethod string, amount float64) error {
	return dao.RecordPayment(r.db, orderID, paymentMethod, amount)
}

func (r *paymentRepository) ConfirmPayment(ctx context.Context, transactionID string) error {
	return dao.ConfirmPayment(r.db, transactionID)
}

func (r *paymentRepository) RefundPayment(ctx context.Context, transactionID string, amount float64) error {
	return dao.RefundPayment(r.db, transactionID, amount)
}