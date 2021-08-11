package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/olivere/elastic/v7"
)

func addBreedsData() {
	var arrayBreed = [12][2]string{
		{"Dog", "Retrievers"},
		{"Dog", "French Bulldogs"},
		{"Dog", "German Shepherd Dogs"},
		{"Dog", "Retrievers"},
		{"Dog", "Bulldogs"},
		{"Dog", "Poodles"},

		{"Cat", "Abyssinian Cat"},
		{"Cat", "American Bobtail Cat"},
		{"Cat", "American Curl Cat"},
		{"Cat", "American Shorthair Cat"},
		{"Cat", "American Wirehair Cat"},
		{"Cat", "Balinese-Javanese Cat"},
	}

	for _, v := range arrayBreed {
		breed := esBreed{
			Species: v[0],
			Name:    v[1],
		}
		saveToES(&breed, BREED_INDEX, "")
		fmt.Println(" add breed ")
	}
}

func getAllBreeds(w http.ResponseWriter, specie string) []esBreed {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("animal_specie", specie))
	searchResult, err := readFromES(query, BREED_INDEX)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	var breed esBreed

	var breeds []esBreed

	for _, item := range searchResult.Each(reflect.TypeOf(breed)) {
		p := item.(esBreed)
		breeds = append(breeds, p)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("Failed to read breeds from Elasticsearch %v.\n", err)
		return nil
	}

	return breeds
}
