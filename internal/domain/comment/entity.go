package comment

type Comment struct {
	CommentID int64  `json:"comment_id"`
	ProductID int64  `json:"product_id"`
	UserID    int64  `json:"user_id"`
	Content   string `json:"content"`
	Rating    float64 `json:"rating"`
	CreatedAt string `json:"created_at"`
}