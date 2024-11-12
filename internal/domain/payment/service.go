package payment

import (
	"context"
	"shop-app/domain/payment"
)

type PaymentService interface {
	RecordPayment(ctx context.Context, orderID string, paymentMethod string, amount float64) error
	ConfirmPayment(ctx context.Context, transactionID string) error
	RefundPayment(ctx context.Context, transactionID string, amount float64) error
}

type paymentService struct {
	repo payment.PaymentRepository
}

func NewPaymentService(repo payment.PaymentRepository) *paymentService {
	return &paymentService{repo: repo}
}

func (s *paymentService) RecordPayment(ctx context.Context, orderID string, paymentMethod string, amount float64) error {
	// 记录支付前，检查订单状态是否为待支付
	order, err := s.repo.GetOrder(ctx, orderID)
	if err != nil {
		return err
	}
	if order.Status != "pending" {
		return errors.New("order status is not pending")
	}
	return s.repo.RecordPayment(ctx, orderID, paymentMethod, amount)
}

func (s *paymentService) ConfirmPayment(ctx context.Context, transactionID string) error {
	// 确认支付前，检查交易状态是否为待确认
	payment, err := s.repo.GetPayment(ctx, transactionID)
	if err != nil {
		return err
	}
	if payment.Status != "pending" {
		return errors.New("payment status is not pending")
	}
	return s.repo.ConfirmPayment(ctx, transactionID)
}

func (s *paymentService) RefundPayment(ctx context.Context, transactionID string, amount float64) error {
	// 退款前，检查支付状态是否为已支付
	payment, err := s.repo.GetPayment(ctx, transactionID)
	if err != nil {
		return err
	}
	if payment.Status != "succeeded" {
		return errors.New("payment status is not succeeded")
	}
	return s.repo.RefundPayment(ctx, transactionID, amount)
}