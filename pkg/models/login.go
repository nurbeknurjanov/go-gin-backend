package models

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey      = "some key"
	refreshTokenTTL = 12 * time.Hour
	accessTokenTTL  = 10 * time.Hour
	//accessTokenTTL  = 10 * time.Second
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type CustomStandardClaims struct {
	jwt.StandardClaims
	User *User `json:"user"`
}

func GenerateTokens(u *User) *Tokens {
	accessToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomStandardClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix()},
		User: u,
	}).SignedString([]byte(signingKey))

	refreshToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomStandardClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix()},
		User: u,
	}).SignedString([]byte(signingKey))

	return &Tokens{
		accessToken,
		refreshToken,
	}
}

func GenerateAccessToken(u *User) string {
	accessToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomStandardClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix()},
		User: u,
	}).SignedString([]byte(signingKey))

	return accessToken
}

func ParseAccessToken(accessToken string) (*User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &CustomStandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomStandardClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.User, nil
}
