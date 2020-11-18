package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// GenToken generates a JWT for the specified user.
func GenToken(log *log.Logger) error {
	privatePEM, err := ioutil.ReadFile("./private.pem")
	if err != nil {
		return errors.Wrap(err, "reading PEM private key file")
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return errors.Wrap(err, "parsing PEM into private key")
	}

	claims := struct {
		jwt.StandardClaims
		Kid string
		Roles []string
	}{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "class project",
			Subject:   "5cf37266-3473-4006-984f-9325122678b7",
			ExpiresAt: time.Now().Add(8760 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},			
		Roles: []string{"ADMIN"},
	}

	algorithm := "RS256"
	method := jwt.GetSigningMethod(algorithm)
	token := jwt.NewWithClaims(method, claims)
	token.Header["kid"] = "54bb2165-71e1-41a6-af3e-7da4a0e1e2c1"
	
	str, err := token.SignedString(privateKey)
	if err != nil {
		return errors.Wrap(err, "signing token")
	}

	fmt.Printf("-----BEGIN TOKEN-----\n%s\n-----END TOKEN-----\n", str)
	return nil
}