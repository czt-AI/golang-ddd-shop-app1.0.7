package v1

import (
	"encoding/json"
	"net/http"
	"shop-app/api/v1/dto"
	"shop-app/api/v1/service"
	"shop-app/infrastructure/response"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) RecordPayment(w http.ResponseWriter, r *http.Request) {
	var req dto.RecordPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ERROR(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.RecordPayment(context.Background(), &req); err != nil {
		response.ERROR(w, "Error recording payment", http.StatusInternalServerError)
		return
	}

	response.SUCCESS(w, "Payment recorded successfully", http.StatusOK)
}

func (h *PaymentHandler) ConfirmPayment(w http.ResponseWriter, r *http.Request) {
	var req dto.ConfirmPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ERROR(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.ConfirmPayment(context.Background(), req.TransactionID); err != nil {
		response.ERROR(w, "Error confirming payment", http.StatusInternalServerError)
		return
	}

	response.SUCCESS(w, "Payment confirmed successfully", http.StatusOK)
}

func (h *PaymentHandler) RefundPayment(w http.ResponseWriter, r *http.Request) {
	var req dto.RefundPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ERROR(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.RefundPayment(context.Background(), req.TransactionID, req.Amount); err != nil {
		response.ERROR(w, "Error refunding payment", http.StatusInternalServerError)
		return
	}

	response.SUCCESS(w, "Payment refunded successfully", http.StatusOK)
}