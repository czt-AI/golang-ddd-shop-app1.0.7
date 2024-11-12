package v1

import (
	"context"
	"shop-app/api/v1/dto"
	"shop-app/domain/comment"
)

type CommentService interface {
	CreateComment(ctx context.Context, req *dto.CreateCommentRequest) error
	GetComments(ctx context.Context, productID int64, page, limit int) ([]*dto.CommentResponse, error)
	GetComment(ctx context.Context, commentID int64) (*dto.CommentResponse, error)
	DeleteComment(ctx context.Context, commentID int64) error
}

type commentService struct {
	usecase comment.CommentUsecase
}

func NewCommentService(usecase comment.CommentUsecase) *commentService {
	return &commentService{usecase: usecase}
}

func (s *commentService) CreateComment(ctx context.Context, req *dto.CreateCommentRequest) error {
	return s.usecase.CreateComment(ctx, req)
}

func (s *commentService) GetComments(ctx context.Context, productID int64, page, limit int) ([]*dto.CommentResponse, error) {
	return s.usecase.GetComments(ctx, productID, page, limit)
}

func (s *commentService) GetComment(ctx context.Context, commentID int64) (*dto.CommentResponse, error) {
	return s.usecase.GetComment(ctx, commentID)
}

func (s *commentService) DeleteComment(ctx context.Context, commentID int64) error {
	return s.usecase.DeleteComment(ctx, commentID)
}