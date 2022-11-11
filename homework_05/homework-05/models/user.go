package models
import "github.com/dgrijalva/jwt-go"


type User struct {
	UserName string  `form:"username" json:"username" binding:"required"`
	Password string	`form:"password" json:"password" binding:"required"`
	PhoneNumber string `form:"phone" json:"phone" binding:"required"`
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
