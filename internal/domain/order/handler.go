package order

import (
	"context"
	"encoding/json"
	"net/http"
	"shop-app/api/v1/order"
	"shop-app/domain/order"
	"shop-app/infrastructure/response"
)

type OrderHandler struct {
	service order.OrderService
}

func NewOrderHandler(service order.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req order.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ERROR(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateOrder(context.Background(), &req.Order); err != nil {
		response.ERROR(w, "Error creating order", http.StatusInternalServerError)
		return
	}

	response.SUCCESS(w, "Order created successfully", http.StatusOK)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("order_id")
	if orderID == "" {
		response.ERROR(w, "Order ID is required", http.StatusBadRequest)
		return
	}

	order, err := h.service.GetOrder(context.Background(), orderID)
	if err != nil {
		response.ERROR(w, "Error fetching order", http.StatusInternalServerError)
		return
	}

	response.JSON(w, order, http.StatusOK)
}