package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"shop-app/api/v1/dto"
	"shop-app/api/v1/service"
	"shop-app/infrastructure/response"
)

type CommentHandler struct {
	service service.CommentService
}

func NewCommentHandler(service service.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ERROR(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateComment(context.Background(), &req.Comment); err != nil {
		response.ERROR(w, "Error creating comment", http.StatusInternalServerError)
		return
	}

	response.SUCCESS(w, "Comment created successfully", http.StatusOK)
}

func (h *CommentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("product_id")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	comments, err := h.service.GetComments(context.Background(), comment.IDFromQuery(productID), page, limit)
	if err != nil {
		response.ERROR(w, "Error fetching comments", http.StatusInternalServerError)
		return
	}

	response.JSON(w, comments, http.StatusOK)
}

func (h *CommentHandler) GetComment(w http.ResponseWriter, r *http.Request) {
	commentID := r.URL.Query().Get("comment_id")
	if commentID == "" {
		response.ERROR(w, "Comment ID is required", http.StatusBadRequest)
		return
	}

	commentResponse, err := h.service.GetComment(context.Background(), comment.IDFromQuery(commentID))
	if err != nil {
		response.ERROR(w, "Error fetching comment", http.StatusInternalServerError)
		return
	}

	response.JSON(w, commentResponse, http.StatusOK)
}

func (h *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentID := r.URL.Query().Get("comment_id")
	if commentID == "" {
		response.ERROR(w, "Comment ID is required", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteComment(context.Background(), comment.IDFromQuery(commentID)); err != nil {
		response.ERROR(w, "Error deleting comment", http.StatusInternalServerError)
		return
	}

	response.SUCCESS(w, "Comment deleted successfully", http.StatusOK)
}