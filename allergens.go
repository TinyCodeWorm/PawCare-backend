package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/olivere/elastic/v7"
)

//algorithm, first find food- that cause reactions and food has no reactions
// put the ingredients in two list, and filter the ingredients do not cause
// reactions, the left ingredients will be returned
func getPetAllergens(w http.ResponseWriter, useremail string) ([]string, error) {

	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("email", useremail))
	searchResult, error := readFromES(query, FOOD_INDEX)
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		fmt.Printf("Cannot get food data from index %v\n", error)
	}
	//find all petReactions
	// pets := getPets(w, email)
	petReactions, error := getPetReactions(w, useremail)
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		fmt.Printf("Cannot get petReaction data from user")
	}

	var noAllergens []string
	var possibleAllergens []string

	var food Food

	for i := 0; i < len(petReactions); i++ {
		for _, item := range searchResult.Each(reflect.TypeOf(food)) {
			if petReactions[i].FoodName == item.(Food).FoodName {
				alg1 := item.(Food).Ingredient1
				alg2 := item.(Food).Ingredient2
				alg3 := item.(Food).Ingredient3
				alg4 := item.(Food).Ingredient4
				alg5 := item.(Food).Ingredient5
				alg6 := item.(Food).Ingredient6
				if petReactions[i].ReactionName == "No Infection" {
					noAllergens = append(noAllergens, alg1, alg2, alg3, alg4, alg5, alg6)
				} else {
					possibleAllergens = append(possibleAllergens, alg1, alg2, alg3, alg4, alg5, alg6)
				}

			}
		}

	}

	// filter noAllergen
	return findAllergens(possibleAllergens, noAllergens), nil
}

func findAllergens(possibleAllergens, noAllergens []string) []string {
	var petAllergens []string
	for i := 0; i < len(possibleAllergens); i++ {
		for j := 0; i < len(noAllergens); j++ {
			if possibleAllergens[i] != noAllergens[j] {
				petAllergens = append(petAllergens, possibleAllergens[i])
			}
		}

	}
	return petAllergens
}
