package comment

import (
	"context"
	"shop-app/domain/comment"
)

type CommentAggregate struct {
	repo comment.CommentRepository
}

func NewCommentAggregate(repo comment.CommentRepository) *CommentAggregate {
	return &CommentAggregate{repo: repo}
}

func (a *CommentAggregate) CreateComment(ctx context.Context, comment *comment.Comment) error {
	return a.repo.CreateComment(ctx, comment)
}

func (a *CommentAggregate) GetComments(ctx context.Context, productID int64, page, limit int) ([]*comment.Comment, error) {
	return a.repo.GetComments(ctx, productID, page, limit)
}

func (a *CommentAggregate) GetComment(ctx context.Context, commentID int64) (*comment.Comment, error) {
	return a.repo.GetComment(ctx, commentID)
}

func (a *CommentAggregate) DeleteComment(ctx context.Context, commentID int64) error {
	return a.repo.DeleteComment(ctx, commentID)
}