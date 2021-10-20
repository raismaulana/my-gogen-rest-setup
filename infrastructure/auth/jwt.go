package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type CustomClaims struct {
	Activated bool
	Email     string
	ID        int64
	Role      string
	jwt.StandardClaims
}

type GenerateTokenRequest struct {
	Audience  string
	Issuer    string
	Subject   string
	ID        int64
	Email     string
	Activated bool
	Role      string
	SecretKey string
}

func GenerateToken(req GenerateTokenRequest) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		Activated: req.Activated,
		Email:     req.Email,
		ID:        req.ID,
		Role:      req.Role,
		StandardClaims: jwt.StandardClaims{
			Audience:  req.Audience,
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Id:        uuid.NewString(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    req.Issuer,
			NotBefore: time.Now().Unix(),
			Subject:   req.Subject,
		},
	})

	tokenString, err := token.SignedString([]byte(req.SecretKey))
	if err != nil {
		return "", err

	}
	return tokenString, nil
}

func VerifyToken(encodedToken string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
