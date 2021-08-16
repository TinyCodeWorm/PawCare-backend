package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/olivere/elastic/v7"
)

// func addPetData() {
// 	var arrayPet = [2][10]string{
// 		{"dog1", "description", "ss", "dsds", "dsds", "ss", "dsds", "dsds", "a@b.com", "1"},
// 		{"dog2", "description", "ss", "dsds", "dsds", "ss", "dsds", "dsds", "a@b.com", "2"},
// 	}

// 	for _, v := range arrayPet {
// 		pet := esPet{

// 			Name:       v[0],
// 			Photourl:   v[1],
// 			Type:       v[2],
// 			Weight:     v[3],
// 			AgeYear:    v[4],
// 			AgeMonth:   v[5],
// 			Sex:        v[6],
// 			Breed:      v[7],
// 			OwnerEmail: v[8],
// 			PetID:      v[9],
// 		}
// 		saveToES(&pet, PET_INDEX, "")
// 		fmt.Println(" add pet ")
// 	}
// }

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
func getPets(w http.ResponseWriter, email string) ([]Pet, error) {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("owner_email", email))
	searchResult, err := readFromES(query, PET_INDEX)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	var pets Pet
	var allPets []Pet

	for _, item := range searchResult.Each(reflect.TypeOf(pets)) {
		p := item.(Pet)
		allPets = append(allPets, p)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("Failed to read pets from Elasticsearch %v.\n", err)
		return nil, err
	}

	return allPets, nil
}
