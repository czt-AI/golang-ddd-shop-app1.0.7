package comment

import (
	"context"
	"shop-app/api/v1/dto"
	"shop-app/domain/comment"
)

type CommentUsecase interface {
	CreateComment(ctx context.Context, req *dto.CreateCommentRequest) error
	GetComments(ctx context.Context, productID int64, page, limit int) ([]*dto.CommentResponse, error)
	GetComment(ctx context.Context, commentID int64) (*dto.CommentResponse, error)
	DeleteComment(ctx context.Context, commentID int64) error
}

type commentUsecase struct {
	repo comment.CommentRepository
}

func NewCommentUsecase(repo comment.CommentRepository) *commentUsecase {
	return &commentUsecase{repo: repo}
}

func (u *commentUsecase) CreateComment(ctx context.Context, req *dto.CreateCommentRequest) error {
	return u.repo.CreateComment(ctx, &req.Comment)
}

func (u *commentUsecase) GetComments(ctx context.Context, productID int64, page, limit int) ([]*dto.CommentResponse, error) {
	dbComments, err := u.repo.GetComments(ctx, productID, page, limit)
	if err != nil {
		return nil, err
	}
	comments := make([]*dto.CommentResponse, 0, len(dbComments))
	for _, dbComment := range dbComments {
		comments = append(comments, dto.ConvertToCommentResponse(dbComment))
	}
	return comments, nil
}

func (u *commentUsecase) GetComment(ctx context.Context, commentID int64) (*dto.CommentResponse, error) {
	dbComment, err := u.repo.GetComment(ctx, commentID)
	if err != nil {
		return nil, err
	}
	return dto.ConvertToCommentResponse(dbComment), nil
}

func (u *commentUsecase) DeleteComment(ctx context.Context, commentID int64) error {
	return u.repo.DeleteComment(ctx, commentID)
}