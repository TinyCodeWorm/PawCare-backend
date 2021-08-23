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
	} else if err != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		fmt.Printf("Cannot get food data from user")
	}

	allergens := make(map[string]bool)
	noAllergens := make(map[string]bool)

	for i := 0; i < len(petReactions); i++ {
		for j := 0; j < len(allFood); j++ {
			alg1 := allFood[j].Ingredient1
			alg2 := allFood[j].Ingredient2
			alg3 := allFood[j].Ingredient3
			alg4 := allFood[j].Ingredient4
			alg5 := allFood[j].Ingredient5
			alg6 := allFood[j].Ingredient6
			if petReactions[i].FoodName == allFood[j].FoodName {
				
				if !allergens[alg1] {
					allergens[alg1] = true
				} 
				if !allergens[alg2] {
					allergens[alg2] = true
				}
				if !allergens[alg3] {
					allergens[alg3] = true
				}
				if !allergens[alg4] {
					allergens[alg4] = true
				}
				if !allergens[alg5] {
					allergens[alg5] = true
				} 
				if !allergens[alg6] {
					allergens[alg6] = true
				}

			} 
			if petReactions[i].ReactionName == "No Reaction" {
				if !noAllergens[alg1] {
					noAllergens[alg1] = true
				} 
				if !noAllergens[alg2] {
					noAllergens[alg2] = true
				}
				if !noAllergens[alg3] {
					noAllergens[alg3] = true
				}
				if !noAllergens[alg4] {
					noAllergens[alg4] = true
				}
				if !noAllergens[alg5]
					noAllergens[alg5] = true
				} 
				if !noAllergens[alg6] {
					noAllergens[alg6] = true
				}
			}
		}

	}
	petAllergens []string
	// filter noAllergen
	for allergen, item : range allergens {
		if !noAllergens[allergen] {
			petAllergens = append(petAllergens, allergen)
		}	

	}
		
	return petAllergens, nil
}


