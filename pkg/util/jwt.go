package util

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte


type MyCustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	var token string
	var err error

	expiresAt := time.Now().Add(24 * time.Hour)

	claims := MyCustomClaims{
		Username: EncodingWithMD5(username),
		Password: EncodingWithMD5(password),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			Issuer:    "test",
		},
	}

	// example 
	/*
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	*/

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*MyCustomClaims, error){
	myToken, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := myToken.Claims.(*MyCustomClaims); ok && myToken.Valid {
		return claims,nil
	}

	return nil, err
}
