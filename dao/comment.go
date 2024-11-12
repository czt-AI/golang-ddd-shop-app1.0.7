package dao

import (
	"shop-app/domain/comment"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	CommentID    int64  `gorm:"index"`
	ProductID    int64  `gorm:"index"`
	UserID       int64  `gorm:"index"`
	Content      string `gorm:"type:text"`
	Rating       float64 `gorm:"index"`
	CreatedAt    string `gorm:"index"`
	UpdatedAt    string `gorm:"index"`
}

func (c *Comment) Create(db *gorm.DB) error {
	return db.Create(c).Error
}

func (c *Comment) GetComments(db *gorm.DB, productID int64, page, limit int) ([]*Comment, error) {
	var comments []*Comment
	query := db.Where("product_id = ?", productID)
	if page > 0 && limit > 0 {
		query = query.Limit(limit).Offset((page - 1) * limit)
	}
	err := query.Find(&comments).Error
	return comments, err
}

func (c *Comment) GetComment(db *gorm.DB, commentID int64) (*Comment, error) {
	var comment Comment
	err := db.Where("id = ?", commentID).First(&comment).Error
	return &comment, err
}

func (c *Comment) Delete(db *gorm.DB, commentID int64) error {
	return db.Delete(&Comment{CommentID: commentID}).Error
}