package users

// User struct
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// UserLoginRequest Login Request
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
