package usermiddleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func JWTgenerate(email, role string) (string, error) {
	var loginkey = []byte("secretkeyjwt")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(loginkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func UserAuthorization(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header)
		if r.Header["Token"] == nil {
			json.NewEncoder(w).Encode("Error in token no token found")
			return
		}
		var loginkey = []byte("secretkeyjwt")
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("error in parsing jwt")
			}
			return loginkey, nil
		})

		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {

				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {

				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return
			}
		}
		json.NewEncoder(w).Encode(err)
	}
}
