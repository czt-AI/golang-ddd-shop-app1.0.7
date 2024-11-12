package v1

import (
	"encoding/json"
	"net/http"
	"shop-app/api/v1/dto"
	"shop-app/api/v1/service"
	"shop-app/infrastructure/response"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ERROR(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.RegisterUser(context.Background(), &req); err != nil {
		response.ERROR(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	response.SUCCESS(w, "User registered successfully", http.StatusOK)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ERROR(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.service.LoginUser(context.Background(), req.Username, req.Password)
	if err != nil {
		response.ERROR(w, "Error logging in", http.StatusInternalServerError)
		return
	}

	response.JSON(w, user, http.StatusOK)
}