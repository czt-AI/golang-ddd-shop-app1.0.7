package v1

import (
	"encoding/json"
	"net/http"
	"shop-app/api/v1/dto"
	"shop-app/api/v1/service"
	"shop-app/infrastructure/response"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		response.ERROR(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	productResponse, err := h.service.GetProductByID(context.Background(), dto.IDFromQuery(id))
	if err != nil {
		response.ERROR(w, "Error fetching product", http.StatusInternalServerError)
		return
	}

	response.JSON(w, productResponse, http.StatusOK)
}

func (h *ProductHandler) GetProductList(w http.ResponseWriter, r *http.Request) {
	categoryID := r.URL.Query().Get("category_id")
	tag := r.URL.Query().Get("tag")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	productResponse, err := h.service.GetProductList(context.Background(), categoryID, tag, page, limit)
	if err != nil {
		response.ERROR(w, "Error fetching product list", http.StatusInternalServerError)
		return
	}

	response.JSON(w, productResponse, http.StatusOK)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ERROR(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	productResponse, err := h.service.CreateProduct(context.Background(), &req)
	if err != nil {
		response.ERROR(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	response.JSON(w, productResponse, http.StatusCreated)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ERROR(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	productResponse, err := h.service.UpdateProduct(context.Background(), &req)
	if err != nil {
		response.ERROR(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	response.JSON(w, productResponse, http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		response.ERROR(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteProduct(context.Background(), dto.IDFromQuery(id)); err != nil {
		response.ERROR(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

	response.SUCCESS(w, "Product deleted successfully", http.StatusOK)
}