package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"flaminio/database"
	"flaminio/utility"
	"encoding/json"
	"errors"
	"flaminio/models"
	"github.com/dgrijalva/jwt-go"
	"time"
	"database/sql"
)

func createNewToken(user models.User, c *gin.Context) (tokenString string) {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := jwt.StandardClaims {
		ExpiresAt: time.Now().Add(time.Hour * 10).Unix(), //10 hours from now
		IssuedAt: time.Now().Unix(),
		NotBefore: time.Now().Add(time.Minute * -2).Unix(), //two minutes ago
		Subject: user.UUID.String(),
	}

	token.Claims = claims

	tokenString, err := token.SignedString(utility.SignKey)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("error while signing the token: " + err.Error()))
		fmt.Fprintln(c.Writer, "Error while signing the token")
		return
	}

	return tokenString
}

type UserCredentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
func LoginHandler(c *gin.Context) {
	var userInput UserCredentials

	err := json.NewDecoder(c.Request.Body).Decode(&userInput)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("error in request"))
		fmt.Fprint(c.Writer, "Error in request")
		return
	}

	user, err := database.GetUserByEmail(userInput.Email)
	if err != sql.ErrNoRows {
		utility.Fatal(err)
	}

	if err == sql.ErrNoRows || !utility.CheckPasswordHash(userInput.Password, user.Password) {
		c.AbortWithError(http.StatusForbidden, errors.New("error logging in"))
		fmt.Fprint(c.Writer, "Invalid credentials")
		return
	}

	database.AddDatabaseLog(user.UUID, utility.GetUUIDFromMapSafely("Authentication", database.LogOperationTypeMap),"User logged in")

	c.Header("Authorization", "Bearer " + createNewToken(user, c))
}

func RefreshHandler(c *gin.Context) {
	c.Header("Authorization", "Bearer " + createNewToken(getUserFromContext(c), c))
}
