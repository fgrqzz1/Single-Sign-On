package auth

import (
	"time"
)

type User struct {
	ID 			string
	Email 		string
	Password 	string
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

type Service interface {
	SignIn(email, password string) (TokenPair, error)
	SignUp(email, password string) (User, error)
	ValidateToken(accessToken string) (string, error)
}
