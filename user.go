package main

import (
	"fmt"
	"reflect"

	"github.com/olivere/elastic/v7"
)

func checkUser(username, password string) (bool, error) {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("email", username))
	query.Must(elastic.NewTermQuery("password", password))
	searchResult, err := readFromES(query, USER_INDEX)
	if err != nil {
		return false, err
	}

	var utype User
	for _, item := range searchResult.Each(reflect.TypeOf(utype)) {
		u := item.(User)
		if u.Password == password {
			fmt.Printf("Login as %s\n", username)
			return true, nil
		}
	}
	return false, nil
}

func addUser(user *User) (bool, error) {
	query := elastic.NewTermQuery("email", user.Email)
	searchResult, err := readFromES(query, USER_INDEX)
	if err != nil {
		return false, err
	}

	if searchResult.TotalHits() > 0 {
		return false, nil
	}

	err = saveToES(user, USER_INDEX, user.Email)
	if err != nil {
		return false, err
	}
	fmt.Printf("User is added: %s\n", user.Email)
	return true, nil
}
