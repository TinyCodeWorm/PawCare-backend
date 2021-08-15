package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/olivere/elastic/v7"
)

func getFoods(w http.ResponseWriter, email string) ([]Food, error) {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("owner_email", email))
	searchResult, err := readFromES(query, FOOD_INDEX)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	var food Food
	var allFood []Food

	for _, item := range searchResult.Each(reflect.TypeOf(food)) {
		p := item.(Food)
		allFood = append(allFood, p)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("Failed to read food from Elasticsearch %v.\n", err)
		return nil, err
	}

	return allFood, nil
}