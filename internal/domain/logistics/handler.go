package logistics

import (
	"context"
	"encoding/json"
	"net/http"
	"shop-app/api/v1/logistics"
	"shop-app/domain/logistics"
	"shop-app/infrastructure/response"
)

type LogisticsHandler struct {
	service logistics.LogisticsService
}

func NewLogisticsHandler(service logistics.LogisticsService) *LogisticsHandler {
	return &LogisticsHandler{service: service}
}

func (h *LogisticsHandler) CreateLogistics(w http.ResponseWriter, r *http.Request) {
	var req logistics.CreateLogisticsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ERROR(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateLogistics(context.Background(), &req.LogisticsOrder); err != nil {
		response.ERROR(w, "Error creating logistics", http.StatusInternalServerError)
		return
	}

	response.SUCCESS(w, "Logistics created successfully", http.StatusOK)
}

func (h *LogisticsHandler) UpdateLogisticsStatus(w http.ResponseWriter, r *http.Request) {
	var req logistics.UpdateLogisticsStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ERROR(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateLogisticsStatus(context.Background(), req.LogisticsID, req.Status, req.Description); err != nil {
		response.ERROR(w, "Error updating logistics status", http.StatusInternalServerError)
		return
	}

	response.SUCCESS(w, "Logistics status updated successfully", http.StatusOK)
}

func (h *LogisticsHandler) GetLogisticsDetails(w http.ResponseWriter, r *http.Request) {
	logisticsID := r.URL.Query().Get("logistics_id")
	if logisticsID == "" {
		response.ERROR(w, "Logistics ID is required", http.StatusBadRequest)
		return
	}

	logisticsOrder, err := h.service.GetLogisticsDetails(context.Background(), logisticsID)
	if err != nil {
		response.ERROR(w, "Error fetching logistics details", http.StatusInternalServerError)
		return
	}

	response.JSON(w, logisticsOrder, http.StatusOK)
}