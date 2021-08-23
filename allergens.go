package main

import (
	"fmt"
	"net/http"
)

//algorithm, first find food- that cause reactions and food has no reactions
// put the ingredients in two list, and filter the ingredients do not cause
// reactions, the left ingredients will be returned
func getPetAllergens(w http.ResponseWriter, useremail string) ([]string, error) {

	//find all petReactions
	// pets := getPets(w, email)
	petReactions, error := getPetReactions(w, useremail)
	allFood, err := getFoods(w, useremail)
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		fmt.Printf("Cannot get petReaction data from user")
	}

	if err != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		fmt.Printf("Cannot get food data from user")
	}

	var noAllergens []string
	var possibleAllergens []string

	for i := 0; i < len(petReactions); i++ {
		for j := 0; j < len(allFood); j++ {
			if petReactions[i].FoodName == allFood[j].FoodName {
				alg1 := allFood[j].Ingredient1
				alg2 := allFood[j].Ingredient2
				alg3 := allFood[j].Ingredient3
				alg4 := allFood[j].Ingredient4
				alg5 := allFood[j].Ingredient5
				alg6 := allFood[j].Ingredient6
				fmt.Printf(alg6)
				if petReactions[i].ReactionName == "No Reaction" {
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
