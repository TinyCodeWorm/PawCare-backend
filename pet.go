package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/olivere/elastic/v7"
)

func getPetReactions(w http.ResponseWriter, email string) ([]PetReaction, error) {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("owner_email", email))
	searchResult, err := readFromES(query, PETREACTION_INDEX)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	var petrea PetReaction
	var petreas []PetReaction

	for _, item := range searchResult.Each(reflect.TypeOf(petrea)) {
		p := item.(PetReaction)
		petreas = append(petreas, p)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("Failed to read food from Elasticsearch %v.\n", err)
		return nil, err
	}

	return petreas, nil
}

func uploadPetRea(w http.ResponseWriter, petrea Petrea, email string) {

	var reas = petrea.Reactions
	var espetrea esPetReaction

	for _, rea := range reas {
		espetrea.ReactionName = rea
		espetrea.OwnerEmail = email
		espetrea.FoodName = petrea.FoodName
		espetrea.ReactionDate = petrea.ReactionDate
		espetrea.PetName = "temp pet name"
		//need a get pet function to add the pet name

		if err := saveToES(espetrea, PETREACTION_INDEX, ""); err != nil {
			http.Error(w, "Cannot save pet reaction data from client", http.StatusBadRequest)
			fmt.Printf("Cannot save pet reaction  data from client %v\n", err)
		}

		fmt.Println("pet reaction : " + rea + " is saved successfully.")
	}

}
