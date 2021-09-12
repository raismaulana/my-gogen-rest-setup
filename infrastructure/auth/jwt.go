package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/raismaulana/my-gogen-rest-setup/infrastructure/envconfig"
)

type JWTToken struct {
	env *envconfig.EnvConfig
}

type CustomClaims struct {
	Activated bool
	Email     string
	ID        int64
	Role      string
	jwt.StandardClaims
}

type GenerateTokenRequest struct {
	Subject   string
	ID        int64
	Email     string
	Activated bool
	Role      string
}

func NewJWTToken(env *envconfig.EnvConfig) (*JWTToken, error) {
	if strings.TrimSpace(env.SecretKey) == "" {
		return nil, fmt.Errorf("JWT secret key must not empty")
	}

	return &JWTToken{
		env: env,
	}, nil

}

func (r *JWTToken) GenerateToken(req GenerateTokenRequest) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		Activated: req.Activated,
		Email:     req.Email,
		ID:        req.ID,
		Role:      req.Role,
		StandardClaims: jwt.StandardClaims{
			Audience:  r.env.AppBaseURL,
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Id:        uuid.NewString(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    r.env.AppBaseURL,
			NotBefore: time.Now().Unix(),
			Subject:   req.Subject,
		},
	})

	tokenString, err := token.SignedString([]byte(r.env.SecretKey))
	if err != nil {
		return "", err

	}
	return tokenString, nil
}

func (r *JWTToken) VerifyToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.env.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
