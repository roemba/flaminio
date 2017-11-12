package flaminio

import (
	"net/http"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/gin-gonic/gin"
	"errors"
	"log"
)

func ValidateTokenMiddleware(c *gin.Context){

	token, err := request.ParseFromRequestWithClaims(c.Request, request.AuthorizationHeaderExtractor, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

	if err == nil {
		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			//Get user and store it in the context
			var user User
			var permissions []Permission
			db.First(&user, User{StandardModel:StandardModel{UUID: claims.Subject}})
			db.Model(&user).Related(&permissions, "Permissions")
			user.Permissions = permissions

			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithError(http.StatusUnauthorized, errors.New("token is not valid"))
			fmt.Fprint(c.Writer, "Token is not valid")
		}
	} else {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized access to this resource"))
		fmt.Fprint(c.Writer, "Unauthorized access to this resource")
	}
	c.Abort()
}
