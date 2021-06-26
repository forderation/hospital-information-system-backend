package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/forderation/hospital-information-system/db"
	"os"
)

func CreateToken(user db.User) (token string, err error) {
	secretKey := os.Getenv("SECRET_KEY")
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["exp"] = true

	return
}
