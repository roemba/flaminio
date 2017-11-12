package flaminio

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"errors"
)

const (
	STATUS_SUCCESS = "success"
	STATUS_FAIL = "failed"
)

type UserCredentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Status string `json:"status"`
	Data interface{} `json:"data"`
}

func createNewToken(user User, c *gin.Context) (tokenString string) {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := jwt.StandardClaims {
		ExpiresAt: time.Now().Add(time.Hour * 10).Unix(), //10 hours from now
		IssuedAt: time.Now().Unix(),
		NotBefore: time.Now().Add(time.Minute * -2).Unix(), //two minutes ago
		Subject: user.UUID.String(),
	}

	token.Claims = claims

	tokenString, err := token.SignedString(signKey)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("error while signing the token"))
		fmt.Fprintln(c.Writer, "Error while signing the token")
		fatal(err)
		return
	}

	return tokenString
}

func getUserFromContext(c *gin.Context) (user User) {
	value, exists := c.Get("user")
	if !exists {
		c.AbortWithError(http.StatusInternalServerError, errors.New("error extracting the key"))
		fmt.Fprintln(c.Writer, "Error extracting the key")
		fatal(errors.New("could not load user from context"))
		return
	}
	return value.(User)
}

func LoginHandler(c *gin.Context) {
	var userInput UserCredentials

	err := json.NewDecoder(c.Request.Body).Decode(&userInput)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("error in request"))
		fmt.Fprint(c.Writer, "Error in request")
		return
	}

	user, recordNotFound := getUser(userInput.Email)
	if recordNotFound || !checkPasswordHash(userInput.Password, user.Password) {
		c.AbortWithError(http.StatusForbidden, errors.New("error logging in"))
		fmt.Fprint(c.Writer, "Invalid credentials")
		return
	}

	c.Writer.Header().Set("Authorization", "Bearer " + createNewToken(user, c))
}

//func ProtectedHandler(c *gin.Context) {
//	user := getUserFromContext(c)
//
//	log.Println(user.Email)
//	response := Response{STATUS_SUCCESS,"Gained access to protected resource"}
//	JsonResponse(response, c.Writer)
//}

func UserHandler(c *gin.Context) {
	user := getUserFromContext(c)

	JsonResponse(Response{STATUS_SUCCESS,user}, c.Writer)
}

func RefreshHandler(c *gin.Context) {
	user := getUserFromContext(c)

	c.Writer.Header().Set("Authorization", "Bearer " + createNewToken(user, c))
}
