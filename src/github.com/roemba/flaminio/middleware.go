package flaminio

import (
	"net/http"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/gin-gonic/gin"
	"errors"
	"github.com/satori/go.uuid"
)

func ValidateTokenMiddleware(c *gin.Context){

	token, err := request.ParseFromRequestWithClaims(c.Request, request.AuthorizationHeaderExtractor, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

	if err == nil {
		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			UUID, err := uuid.FromString(claims.Subject)
			if err != nil {
				c.AbortWithError(http.StatusUnauthorized, errors.New("error converting subject of token to uuid"))
				fmt.Fprint(c.Writer, "Token integrity violated")
			}

			//Get user and store it in the context
			c.Set("user", getUserWithPermissions(UUID))
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
