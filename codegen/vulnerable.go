package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/text/encoding/unicode"
)

func createVulnerableToken() {
	// Using a vulnerable version of jwt-go
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString([]byte("secret"))
	fmt.Println(tokenString)

	// Using a vulnerable version of golang.org/x/text
	encoder := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	encoder.NewEncoder()
}
