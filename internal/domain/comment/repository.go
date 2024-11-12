package comment

import (
	"context"
	"shop-app/internal/dao"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment *dao.Comment) error
	GetComments(ctx context.Context, productID int64, page, limit int) ([]*dao.Comment, error)
	GetComment(ctx context.Context, commentID int64) (*dao.Comment, error)
	DeleteComment(ctx context.Context, commentID int64) error
}

type commentRepository struct {
	db *dao.DB
}

func NewCommentRepository(db *dao.DB) *commentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) CreateComment(ctx context.Context, comment *dao.Comment) error {
	return dao.CreateComment(r.db, comment)
}

func (r *commentRepository) GetComments(ctx context.Context, productID int64, page, limit int) ([]*dao.Comment, error) {
	return dao.GetComments(r.db, productID, page, limit)
}

func (r *commentRepository) GetComment(ctx context.Context, commentID int64) (*dao.Comment, error) {
	return dao.GetComment(r.db, commentID)
}

func (r *commentRepository) DeleteComment(ctx context.Context, commentID int64) error {
	return dao.DeleteComment(r.db, commentID)
}