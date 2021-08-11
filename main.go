package main

import (
	"fmt"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("started-service")

	setupDB()

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	r := mux.NewRouter()

	r.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST", "OPTIONS")
	r.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST", "OPTIONS")
	r.Handle("/getprofile", jwtMiddleware.Handler(http.HandlerFunc(getprofileHandler))).Methods("GET", "OPTIONS")

	r.Handle("/getpets", jwtMiddleware.Handler(http.HandlerFunc(getpetsHandler))).Methods("GET", "OPTIONS")
	r.Handle("/uploadpet", jwtMiddleware.Handler(http.HandlerFunc(uploadpetHandler))).Methods("POST", "OPTIONS")

	r.Handle("/getfoods", jwtMiddleware.Handler(http.HandlerFunc(getfoodsHandler))).Methods("GET", "OPTIONS")
	r.Handle("/uploadfood", jwtMiddleware.Handler(http.HandlerFunc(uploadfoodHandler))).Methods("POST", "OPTIONS")

	r.Handle("/getpetreactions", jwtMiddleware.Handler(http.HandlerFunc(getpetreactionsHandler))).Methods("GET", "OPTIONS")
	r.Handle("/uploadpetreaction", jwtMiddleware.Handler(http.HandlerFunc(uploadpetreactionHandler))).Methods("POST", "OPTIONS")

	r.Handle("/getreactions", jwtMiddleware.Handler(http.HandlerFunc(getreactionsHandler))).Methods("GET", "OPTIONS")

	r.Handle("/getbreeds", jwtMiddleware.Handler(http.HandlerFunc(getbreedsHandler))).Methods("GET", "OPTIONS")

	r.Handle("/getallergens", jwtMiddleware.Handler(http.HandlerFunc(getallergensHandler))).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", r))
}
