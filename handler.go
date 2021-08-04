package main

import (
	"net/http"
)

var mySigningKey = []byte("secret")

func signinHandler(w http.ResponseWriter, r *http.Request) {
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
}

func getprofileHandler(w http.ResponseWriter, r *http.Request) {
}

func getpetsHandler(w http.ResponseWriter, r *http.Request) {
}

func uploadpetHandler(w http.ResponseWriter, r *http.Request) {
}

func getfoodsHandler(w http.ResponseWriter, r *http.Request) {
}

func uploadfoodHandler(w http.ResponseWriter, r *http.Request) {
}

func getreactionsHandler(w http.ResponseWriter, r *http.Request) {
}

func uploadreactionHandler(w http.ResponseWriter, r *http.Request) {
}
