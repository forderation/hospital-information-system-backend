package util

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
	"time"
)

type PayloadToken struct {
	UserId uint
	Exp    int64
}

func CreateToken(userID uint) (token string, err error) {
	secretKey := os.Getenv("SECRET_KEY")
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err = at.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return
}

func ExtractToken(c *gin.Context) string {
	bearToken := c.GetHeader("Authorization")
	token := strings.Split(bearToken, " ")
	if len(token) == 2 {
		return token[1]
	}
	return ""
}

func VerifyToken(c *gin.Context) (token *jwt.Token, err error) {
	tokenString := ExtractToken(c)
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(
				fmt.Sprintf("unexpected signing method: %v",
					token.Header["alg"],
				),
			)
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractPayload(c *gin.Context) (p *PayloadToken, err error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := strconv.ParseUint(
			fmt.Sprintf("%.f", claims["user_id"]),
			10, 64,
		)
		if err != nil {
			return nil, err
		}
		Exp, err := strconv.ParseInt(
			fmt.Sprintf("%.f", claims["exp"]),
			10, 64,
		)
		if time.Now().Unix() > Exp {
			return nil, errors.New("token is expired")
		}
		if err != nil {
			return nil, err
		}
		p = &PayloadToken{
			UserId: uint(userId),
			Exp:    Exp,
		}
		return p, nil
	}
	return nil, err
}
