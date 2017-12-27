package flaminio

import (
	"net/http"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/gin-gonic/gin"
	"errors"
	"github.com/satori/go.uuid"
	"flaminio/database"
	"flaminio/utility"
	"flaminio/models"
	"database/sql"
)

func validateTokenMiddleware(c *gin.Context){

	token, err := request.ParseFromRequestWithClaims(c.Request, request.AuthorizationHeaderExtractor, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				c.AbortWithError(http.StatusUnauthorized, errors.New("unexpected signing method"))
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return utility.VerifyKey, nil
		})

	if err == nil {
		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			if database.IsTokenBlacklisted(token) {
				c.AbortWithError(http.StatusUnauthorized, errors.New("token is on blacklist"))
				fmt.Fprint(c.Writer, "Token is on blacklist")
				return
			}

			UUID, err := uuid.FromString(claims.Subject)
			if err != nil {
				c.AbortWithError(http.StatusUnauthorized, errors.New("error converting subject of token to uuid"))
				fmt.Fprint(c.Writer, "Token integrity violated")
				return
			}

			//Get user and store it in the context
			var user = models.User{
				StandardModel: models.StandardModel{
					UUID: UUID,
				},
			}
			user, err = database.GetUserByUUID(user)
			if err == sql.ErrNoRows {
				c.AbortWithError(http.StatusUnauthorized, errors.New("user does not exist"))
				fmt.Fprint(c.Writer, "User does not exist")
				return
			}
			utility.Fatal(err)
			user, err = database.GetPermissionsForUser(user)
			utility.Fatal(err)

			c.Set("user", user)
			c.Set("token", token)
			c.Next()
		} else {
			c.AbortWithError(http.StatusUnauthorized, errors.New("token is not valid"))
			fmt.Fprint(c.Writer, "Token is not valid")
			return
		}
	} else {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized access to this resource"))
		fmt.Fprint(c.Writer, "Unauthorized access to this resource")
		return
	}
	c.Abort()
}

func checkMediaTypeHeaderIsJson(c *gin.Context) {
	if !(c.GetHeader("Content-Type") == "application/json") {
		c.AbortWithError(http.StatusUnsupportedMediaType, errors.New("invalid media type"))
		fmt.Fprint(c.Writer, "Invalid media type")
		return
	}
	c.Next()
}