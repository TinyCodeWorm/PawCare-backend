package main

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/olivere/elastic/v7"
)

func savePet(myESPet *esPet, file multipart.File) error {

	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("name", myESPet.Name))
	query.Must(elastic.NewTermQuery("owner_email", myESPet.OwnerEmail))

	searchResult, err := readFromES(query, PET_INDEX)
	if err != nil {
		return err
	}

	if searchResult.TotalHits() > 0 {
		return fmt.Errorf("the pet name exists: %s", myESPet.Name)
	}

	if file != nil {
		medialink, err := saveToGCS(file, myESPet.PetID)
		if err != nil {
			return err
		}
		myESPet.Photourl = medialink
	} else {
		if myESPet.Type == "Dog" || myESPet.Type == "dog" {
			myESPet.Photourl = "https://storage.googleapis.com/pawcare-bucket/IMG_7343.JPG"
		} else {
			myESPet.Photourl = "https://storage.googleapis.com/pawcare-bucket/IMG_7342.JPG"
		}

	}
	return saveToES(myESPet, PET_INDEX, myESPet.PetID)
}

func deletPet(useremail string, petname string) error {

	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("owner_email", useremail))
	query.Must(elastic.NewTermQuery("name", petname))

	return deleteFromES(query, PET_INDEX)

}

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

	pets, err := getPets(w, email)
	if err != nil {
		return
	}
	if pets == nil {
		http.Error(w, "no pet exists", http.StatusBadRequest)
		return
	}

	firstPet := pets[0]

	for _, rea := range reas {
		espetrea.ReactionName = rea
		espetrea.OwnerEmail = email
		espetrea.FoodName = petrea.FoodName
		espetrea.ReactionDate = petrea.ReactionDate
		espetrea.PetName = firstPet.Name

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
