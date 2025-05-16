package security

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

type TokenResult struct {
	Token     string
	ExpiresAt time.Time
}

// GenerateToken cria um JWT com expiração e retorna token + expires_at
func GenerateToken(userID int64) (*TokenResult, error) {
	now := time.Now().UTC()
	exp := now.Add(1 * time.Hour)

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   "user_token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &TokenResult{
		Token:     signed,
		ExpiresAt: exp,
	}, nil
}
