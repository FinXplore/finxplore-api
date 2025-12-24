package utils

import (
	"time"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserCode string `json:"user_code"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// 1. CreateAccessToken (JWT) - Short Lived (e.g., 15 mins)
func CreateAccessToken(userCode, role, secret string) (string, error) {
	claims := Claims{
		UserCode: userCode,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // Short life
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "finxplore-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// 2. CreateRefreshToken (Opaque) - Long Lived (e.g., 7 days)
// We just generate a random UUID. We will validate it against Redis later.
func CreateRefreshToken() string {
	return uuid.New().String()
}

// VerifyToken checks if the token is valid and returns the claims
func VerifyToken(tokenString, secret string) (*Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method (Security Check)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract Claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}