package flaminio

import (
	"io/ioutil"
	"github.com/dgrijalva/jwt-go"
	"crypto/rsa"
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
