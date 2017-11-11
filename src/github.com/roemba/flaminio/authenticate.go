package flaminio

import (
	"io/ioutil"
	"github.com/dgrijalva/jwt-go"
	"crypto/rsa"
	"golang.org/x/crypto/bcrypt"
)

var (
	verifyKey *rsa.PublicKey
	signKey *rsa.PrivateKey
)

type CompleteClaims struct {
	Role int `json:"role"`
	jwt.StandardClaims
}

func initKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
