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
	rmq  *messagequeue.RabbitMQClient
}

func NewPaymentService(repo payment.PaymentRepository, rmq *messagequeue.RabbitMQClient) *paymentService {
	return &paymentService{
		repo: repo,
		rmq:  rmq,
	}
}

func (s *paymentService) RecordPayment(ctx context.Context, orderID string, paymentMethod string, amount float64) error {
	// Record payment in the database
	if err := s.repo.RecordPayment(ctx, orderID, paymentMethod, amount); err != nil {
		return err
	}

	// Publish a message to a queue to handle further processing
	if err := s.rmq.Publish("payment", "payment_recorded", []byte(orderID)); err != nil {
		return err
	}

	return nil
}

func (s *paymentService) ConfirmPayment(ctx context.Context, transactionID string) error {
	// Confirm payment in the database
	if err := s.repo.ConfirmPayment(ctx, transactionID); err != nil {
		return err
	}

	// Publish a message to a queue to handle further processing
	if err := s.rmq.Publish("payment", "payment_confirmed", []byte(transactionID)); err != nil {
		return err
	}

	return nil
}

func (s *paymentService) RefundPayment(ctx context.Context, transactionID string, amount float64) error {
	// Refund payment in the database
	if err := s.repo.RefundPayment(ctx, transactionID, amount); err != nil {
		return err
	}

	// Publish a message to a queue to handle further processing
	if err := s.rmq.Publish("payment", "payment_refunded", []byte(transactionID)); err != nil {
		return err
	}

	return nil
}