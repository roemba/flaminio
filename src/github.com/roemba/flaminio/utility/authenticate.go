package utility

import (
	"io/ioutil"
	"github.com/dgrijalva/jwt-go"
	"crypto/rsa"
	"golang.org/x/crypto/bcrypt"
)

const (
	privKeyPath              = "app.rsa"
	pubKeyPath               = "app.rsa.pub"
)

var (
	VerifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey
)

func InitKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	Fatal(err)

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	Fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	Fatal(err)

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	Fatal(err)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
