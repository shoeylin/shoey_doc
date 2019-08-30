package main

import (
	"SinoPacBackstage/models"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"
)

func main() {

	token, err := getToken("shoey")
	if err != nil {
		fmt.Println(err)
	}

	b, userid := checkJWT(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b, userid)

}

const AuthToken = "shoey@tw"

func getToken(userID string) (string, error) {

	//Header
	token := jwt.New(jwt.SigningMethodHS256)
	//Payload
	claims := make(jwt.MapClaims)
	claims["iss"] = userID
	timeNow := time.Now()
	claims["exp"] = timeNow.Add(time.Hour * time.Duration(5)).Unix()
	claims["iat"] = timeNow.Unix()
	token.Claims = claims
	//Signature

	tokenString, err := token.SignedString([]byte(AuthToken))

	beego.Info("login userID:", userID, "tokenString:", tokenString, "err:", err)

	return tokenString, err
}

func checkJWT(tokenString string) (bool, string) {

	//後門
	if tokenString == "abcde" {
		return true, "token開後門測試"
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(models.AuthTokenSinpoac), nil
	})

	if err != nil {
		return false, ""
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if user, mapOk := claims["iss"]; mapOk { //userID
			return true, user.(string)
		} else if id, mapOk := claims["password"]; mapOk {
			return true, id.(string)
		}
		return true, ""
	}
	return false, err.Error()
}
