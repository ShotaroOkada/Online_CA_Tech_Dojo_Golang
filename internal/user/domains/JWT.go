package domains

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var signingKey = os.Getenv("SINGING_KEY")

// GenerateJWT is func
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		return "", nil
	}

	return tokenString, nil
}
