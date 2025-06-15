package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": []string{"admin", "user"},
	})
	tokenString, _ := token.SignedString([]byte("mysecret"))

	parsed, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("mysecret"), nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok || !parsed.Valid {
		fmt.Println("Invalid token")
		return
	}

	if claims.VerifyAudience("not-an-audience", false) {
		fmt.Println("⚠️ Audience bypassed! Token accepted")
	} else {
		fmt.Println("Audience rejected")
	}
}
