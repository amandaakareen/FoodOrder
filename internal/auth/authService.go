package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ProvideJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

var ErrTokenInvalid = errors.New("token inválido")

type AuthService struct {
	secret []byte
}

func NewAuthService(secret string) *AuthService {
	return &AuthService{
		secret: []byte(secret),
	}
}

func (a *AuthService) GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString(a.secret)
}

func (a *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenInvalid
		}
		return a.secret, nil
	})

	if err != nil || !token.Valid {
		return nil, ErrTokenInvalid
	}

	return token, nil
}
