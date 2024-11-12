package user

type UserRegistered struct {
	User *User `json:"user"`
}

type UserLoggedIn struct {
	User *User `json:"user"`
}

type UserUpdated struct {
	User *User `json:"user"`
}

type UserDeleted struct {
	UserID int64 `json:"user_id"`
}