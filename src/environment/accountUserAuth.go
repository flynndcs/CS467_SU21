package environment

import (
	fdbDriver "CS467_SU21/src/store/fdb"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type createAccountStruct struct {
	AccountName string
	Users       []string
}

type createUserGetTokenStruct struct {
	AccountName string
}

func Authenticate(gwmux runtime.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, status := r.BasicAuth()
		if r.Method == "POST" && r.URL.Path == "/createAccount" {
			createAccount(r)
			return
		} else if r.Method == "POST" && r.URL.Path == "/createUser" {
			createUser(r, user, pass, status)
			return
		} else if r.Method == "POST" && r.URL.Path == "/getToken" {
			getToken(r, user, pass, w)
			return
		} else {
			if !validate(r) {
				return
			}
		}
		gwmux.ServeHTTP(w, r)
	})
}

func createAccount(r *http.Request) {
	bodyDecoder := json.NewDecoder(r.Body)
	var body createAccountStruct
	err := bodyDecoder.Decode(&body)
	if err != nil {
		log.Println("could not unmarshal body: ", err)
	}
	accountName := body.AccountName
	users := []byte{}
	for _, v := range body.Users {
		users = append(users, []byte(v)...)
	}
	if !fdbDriver.CreateAccount(accountName, users) {
		log.Println("Account could not be created")
	} else {
		log.Printf("Account created: %s, users: %s", accountName, users)
	}
}

func createUser(r *http.Request, user string, pass string, status bool) {
	bodyDecoder := json.NewDecoder(r.Body)
	var body createUserGetTokenStruct
	err := bodyDecoder.Decode(&body)
	if err != nil {
		log.Println("could not unmarshal body: ", err)
	}
	accountName := body.AccountName
	if !fdbDriver.CreateUser(accountName, user, pass) || !status {
		log.Println("User could not be created")
	} else {
		log.Printf("User %v created for account %v", user, accountName)
	}
}

func getToken(r *http.Request, user string, pass string, w http.ResponseWriter) {
	bodyDecoder := json.NewDecoder(r.Body)
	var body createUserGetTokenStruct
	err := bodyDecoder.Decode(&body)
	if err != nil {
		log.Println("could not unmarshal body: ", err)
	}
	accountName := body.AccountName
	if !fdbDriver.CheckCredentials(accountName, user, pass) {
		log.Println("Unauthorized")
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user,
		"exp": time.Now().Add(1 * time.Hour).Unix(),
	})
	tokenString, tokenErr := token.SignedString(secret)
	if tokenErr != nil {
		log.Println("error generating jwt: ", tokenErr)
		return
	}
	w.Write([]byte(tokenString))
}

func validate(r *http.Request) (validated bool) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	parsedToken, parseErr := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", jwtToken.Header["alg"])
		}
		return secret, nil
	})
	if parseErr != nil {
		log.Println("Could not parse JWT: ", parseErr)
		return false
	}
	if _, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return true
	}
	return false
}
