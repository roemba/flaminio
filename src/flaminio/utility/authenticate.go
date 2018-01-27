package utility

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

var (
	VerifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey
	IsDebug   bool
)

func Init(privKeyPath string, pubKeyPath string, isDebug bool) {
	IsDebug = isDebug
	initKeys(privKeyPath, pubKeyPath)
}

func initKeys(privKeyPath string, pubKeyPath string) {
	LogDebug("Attempting signing tokens with supplied public and private key.")
	LogDebug(fmt.Sprintf("Public key path: %s", pubKeyPath))
	LogDebug(fmt.Sprintf("Private key path: %s", privKeyPath))

	signBytes, err := ioutil.ReadFile(privKeyPath)
	LogFatal(err)

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	LogFatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	LogFatal(err)

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	LogFatal(err)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
