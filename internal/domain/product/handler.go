package product

import (
	"context"
	"encoding/json"
	"net/http"
	"shop-app/api/v1/product"
	"shop-app/domain/product"
	"shop-app/infrastructure/response"
)

type ProductHandler struct {
	service product.ProductService
}

func NewProductHandler(service product.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		response.ERROR(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProductByID(context.Background(), product.IDFromQuery(id))
	if err != nil {
		response.ERROR(w, "Error fetching product", http.StatusInternalServerError)
		return
	}

	response.JSON(w, product, http.StatusOK)
}

func (h *ProductHandler) GetProductList(w http.ResponseWriter, r *http.Request) {
	categoryID := r.URL.Query().Get("category_id")
	tag := r.URL.Query().Get("tag")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	products, err := h.service.GetProductList(context.Background(), categoryID, tag, page, limit)
	if err != nil {
		response.ERROR(w, "Error fetching product list", http.StatusInternalServerError)
		return
	}

	response.JSON(w, products, http.StatusOK)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req product.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ERROR(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product, err := h.service.CreateProduct(context.Background(), &req)
	if err != nil {
		response.ERROR(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	response.JSON(w, product, http.StatusCreated)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var req product.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ERROR(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product, err := h.service.UpdateProduct(context.Background(), &req)
	if err != nil {
		response.ERROR(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	response.JSON(w, product, http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		response.ERROR(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteProduct(context.Background(), product.IDFromQuery(id)); err != nil {
		response.ERROR(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

	response.SUCCESS(w, "Product deleted successfully", http.StatusOK)
}