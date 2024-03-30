package auth

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

type User struct {
	Id string `json:"id"`
}

func NewUser(id string) *User {
	return &User{
		Id: id,
	}
}
