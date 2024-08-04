package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

func main() {
	mySigningKey := []byte("AllYourBase")
	c := MyClaims{
		Username: "shenyaoguan",
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-time.Minute)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			Issuer:    "shenyaougan",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := t.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	fmt.Println(s, err)

	token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		fmt.Errorf("%s", err)
	}
	fmt.Println(err)
	fmt.Println(token.Claims.(*MyClaims).Username)
}
