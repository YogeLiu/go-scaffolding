package jwt

import (
	"errors"
	"net/http"
	"scaffolding/middleware"
	"scaffolding/serializer"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")
		if authToken == "" {
			c.AbortWithStatusJSON(http.StatusOK, serializer.Response{Code: serializer.CodeCheckLogin, Msg: "need login"})
			return
		}

		claims, err := parserToken(authToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, serializer.Response{Code: serializer.CodeNoRightErr, Msg: "token is invalid"})
			return
		}

		if claims.UserId == 0 {
			c.AbortWithStatusJSON(http.StatusOK, serializer.Response{Code: serializer.CodeNoRightErr, Msg: "token is invalid"})
			return
		}

		middleware.SetLoginUserId(c, claims.UserId)
		c.Next()
	}
}

type CustomClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

var jwtKey = []byte("xxx-secret")
var duration = 24

func CreateToken(userId int) (string, error) {
	expire := time.Now().Add(time.Duration(duration) * time.Hour)
	claims := &CustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func parserToken(signToken string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(signToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("token is invalid")
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	return claims, nil
}
