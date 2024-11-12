package user

type User struct {
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
	Password  string `json:"password"` // 应该是加密存储
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	Role      string `json:"role"` // 买家、卖家、管理员
	CreateTime string `json:"create_time"`
	LastLogin string `json:"last_login"`
}