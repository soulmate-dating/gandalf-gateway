package auth

import "github.com/soulmate-dating/gandalf-gateway/internal/app/clients/auth"

// Credentials represents the credentials for a user.
type Credentials struct {
	Email    string `json:"email" binding:"required,email" example:"elon@mail.com"`
	Password string `json:"password" binding:"required,password" example:"password1234"`
}

// User represents a user.
type User struct {
	Id           string `json:"id" binding:"required" example:"d2095501-4295-4cb2-b616-94cd2dc5bfb1"`
	AccessToken  string `json:"access_token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTM5NzY2MDIsImlzcyI6ImF1dGgtc2VydmljZSIsIklkIjoiZDIwOTU1MDEtNDI5NS00Y2IyLWI2MTYtOTRjZDJkYzViZmIxIiwiRW1haWwiOiJtYXhlIn0.y-rHZyYh7i1q0gSqKeRPBBbl-xjfpTu7MOEQzEFozX4"`
	RefreshToken string `json:"refresh_token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU5NjM4MDIsImlzcyI6ImF1dGgtc2VydmljZSIsIklkIjoiZDIwOTU1MDEtNDI5NS00Y2IyLWI2MTYtOTRjZDJkYzViZmIxIiwiRW1haWwiOiJtYXhlIn0.MqRPlXpIU2WKd5t6U5V5yeJQUoC0E_9w8Qa7WPGSgZM"`
}

func NewUser(r *auth.TokenResponse) *User {
	return &User{
		Id:           r.Id,
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
	}
}
