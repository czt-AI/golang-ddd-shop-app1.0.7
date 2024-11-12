package comment

type CommentCreated struct {
	Comment *Comment `json:"comment"`
}

type CommentUpdated struct {
	Comment *Comment `json:"comment"`
}

type CommentDeleted struct {
	CommentID int64 `json:"comment_id"`
}