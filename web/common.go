package web

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetToken(user, passwd string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = user
	claims["passwd"] = passwd
	token.Claims = claims
	tokenstr, err := token.SignedString([]byte("sdfsdfsdfeswfesfdss33e"))
	if err != nil {
		return ""
	} else {
		return tokenstr
	}

}
func ParseToken(token string) (string, string) {
	tok, _ := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		return []byte("sdfsdfsdfeswfesfdss33e"), nil
	})
	if claims, ok := tok.Claims.(jwt.MapClaims); ok && tok.Valid {
		log.Println(ok)
		return claims["sub"].(string), claims["passwd"].(string)
	} else {
		log.Println(ok)
		return "", ""
	}
}
