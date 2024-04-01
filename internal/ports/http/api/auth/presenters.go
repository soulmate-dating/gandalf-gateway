package auth

import "github.com/soulmate-dating/gandalf-gateway/internal/app/clients/auth"

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

type User struct {
	Id           string `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewUser(r *auth.TokenResponse) *User {
	return &User{
		Id:           r.Id,
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
	}
}
