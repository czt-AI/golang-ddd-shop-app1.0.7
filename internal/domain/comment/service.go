package comment

import (
	"context"
	"shop-app/domain/comment"
)

type CommentService interface {
	CreateComment(ctx context.Context, comment *comment.Comment) error
	GetComments(ctx context.Context, productID int64, page, limit int) ([]*comment.Comment, error)
	GetComment(ctx context.Context, commentID int64) (*comment.Comment, error)
	DeleteComment(ctx context.Context, commentID int64) error
}

type commentService struct {
	repo comment.CommentRepository
}

func NewCommentService(repo comment.CommentRepository) *commentService {
	return &commentService{repo: repo}
}

func (s *commentService) CreateComment(ctx context.Context, comment *comment.Comment) error {
	// 创建评论前，检查用户是否可以评论
	if !canComment(ctx, comment.ProductID, comment.UserID) {
		return errors.New("user cannot comment on this product")
	}
	return s.repo.CreateComment(ctx, comment)
}

func (s *commentService) GetComments(ctx context.Context, productID int64, page, limit int) ([]*comment.Comment, error) {
	return s.repo.GetComments(ctx, productID, page, limit)
}

func (s *commentService) GetComment(ctx context.Context, commentID int64) (*comment.Comment, error) {
	return s.repo.GetComment(ctx, commentID)
}

func (s *commentService) DeleteComment(ctx context.Context, commentID int64) error {
	// 删除评论前，检查用户是否可以删除评论
	if !canDeleteComment(ctx, commentID) {
		return errors.New("user cannot delete this comment")
	}
	return s.repo.DeleteComment(ctx, commentID)
}

// canComment 检查用户是否可以评论
func canComment(ctx context.Context, productID, userID int64) bool {
	// 这里定义用户评论规则
	// ...
	return canComment
}

// canDeleteComment 检查用户是否可以删除评论
func canDeleteComment(ctx context.Context, commentID int64) bool {
	// 这里定义用户删除评论规则
	// ...
	return canDelete
}