package model

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	Name string `json:"username"`
	jwt.StandardClaims
}