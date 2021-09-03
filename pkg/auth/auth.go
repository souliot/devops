package auth

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"time"

	jwt "github.com/golang-jwt/jwt"
	"golang.org/x/crypto/scrypt"
)

const pwHashBytes = 64

type Claims struct {
	Appid string `json:"appid"`
	// recommended having
	jwt.StandardClaims
}

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func GenerateSalt() (salt string, err error) {
	buf := make([]byte, pwHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", buf), nil
}

func GeneratePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, pwHashBytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h), nil
}
func To_md5(encode string) (decode string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(encode))
	cipherStr := md5Ctx.Sum(nil)
	return string(base64Encode(cipherStr))
}

func CreateTokenWithExp(appid string, secret string, expireToken int64) (string, int64) {
	claims := Claims{
		appid,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    appid,
		},
	}

	// Create the token using your claims
	c_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, _ := c_token.SignedString([]byte(secret))

	return signedToken, expireToken
}

func Create_token(appid string, secret string, exp int64) (string, int64) {
	expireToken := time.Now().Add(time.Hour * time.Duration(exp)).Unix()
	claims := Claims{
		appid,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    appid,
		},
	}

	// Create the token using your claims
	c_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, _ := c_token.SignedString([]byte(secret))

	return signedToken, expireToken
}

func Token_auth(signedToken, secret string) (string, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		//fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
		//fmt.Println(reflect.TypeOf(claims.StandardClaims.ExpiresAt))
		return claims.Appid, err
	}
	return "", err
}
