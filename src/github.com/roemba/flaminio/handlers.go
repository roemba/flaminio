package flaminio

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"errors"
	"log"
)


type UserCredentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Data string `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}

func LoginHandler(c *gin.Context) {
	var userInput UserCredentials

	err := json.NewDecoder(c.Request.Body).Decode(&userInput)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("error in request"))
		fmt.Fprint(c.Writer, "Error in request")
		return
	}

	var user User
	if query := db.First(&user, User{Email: userInput.Email}); query.RecordNotFound() || !checkPasswordHash(userInput.Password, user.Password) {
		c.AbortWithError(http.StatusForbidden, errors.New("error logging in"))
		fmt.Fprint(c.Writer, "Invalid credentials")
		return
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := jwt.StandardClaims {
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(), //10 hours from now
			IssuedAt: time.Now().Unix(),
			NotBefore: time.Now().Add(time.Minute * -2).Unix(), //two minutes ago
			Subject: user.UUID,
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

func ProtectedHandler(c *gin.Context) {
	value, exists := c.Get("user")
	if !exists {
		c.AbortWithError(http.StatusInternalServerError, errors.New("error extracting the key"))
		fmt.Fprintln(c.Writer, "Error extracting the key")
		fatal(errors.New("could not load user from context"))
		return
	}
	user := value.(User)

	log.Println(user.Email)
	response := Response{"Gained access to protected resource"}
	JsonResponse(response, c.Writer)
}
