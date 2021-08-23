package main

import (
	"fmt"
	"net/http"
	"reflect"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/olivere/elastic/v7"
)

func getProfile(w http.ResponseWriter, r *http.Request) Profile {
	user := r.Context().Value("user")
	claims := user.(*jwt.Token).Claims
	useremail := claims.(jwt.MapClaims)["email"]

	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("email", useremail))
	searchResult, err := readFromES(query, USER_INDEX)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Printf("Cannot get user data from index %v\n", err)
	}

	var profile Profile
	for _, item := range searchResult.Each(reflect.TypeOf(profile)) {
		profile = item.(Profile)
		break
	}

	return profile
}
