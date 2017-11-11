package flaminio

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
	"encoding/json"
	"strings"
	"github.com/gin-gonic/gin"
	"errors"
	"log"
)


type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Data string `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}

func ProtectedHandler(c *gin.Context) {
	response := Response{"Gained access to protected resource"}
	JsonResponse(response, c.Writer)
}

func LoginHandler(c *gin.Context) {
	var user UserCredentials

	err := json.NewDecoder(c.Request.Body).Decode(&user)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("error in request"))
		fmt.Fprint(c.Writer, "Error in request")
		return
	}

	var duser User
	db.Where(&User{Username:user.Username}).First(&duser)
	log.Println(duser.ID)

	if strings.ToLower(user.Username) != "someone" {
		if user.Password != "p@ssword" {
			c.AbortWithError(http.StatusForbidden, errors.New("error logging in"))
			fmt.Fprint(c.Writer, "Invalid credentials")
			return
		}
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := CompleteClaims{
		1,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
			IssuedAt: time.Now().Unix(),
			NotBefore: time.Now().Add(time.Minute * -2).Unix(),
			Id: "appelflapid",
		},
	}
	token.Claims = claims

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("error extracting the key"))
		fmt.Fprintln(c.Writer, "Error extracting the key")
		fatal(err)
	}

	tokenString, err := token.SignedString(signKey)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("error while signing the token"))
		fmt.Fprintln(c.Writer, "Error while signing the token")
		fatal(err)
	}

	response := Token{tokenString}
	JsonResponse(response, c.Writer)
}
