package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/olivere/elastic/v7"
)

func getAllReactions(w http.ResponseWriter) []Reaction {
	query := elastic.NewMatchQuery("name", "")
	query.ZeroTermsQuery("all")
	searchResult, err := readFromES(query, REACTION_INDEX)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Printf("Cannot get reaction data from index %v\n", err)
	}

	var reaction Reaction
	var reactions []Reaction

	for _, item := range searchResult.Each(reflect.TypeOf(reaction)) {
		p := item.(Reaction)
		reactions = append(reactions, p)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("Failed to read reactions from Elasticsearch %v.\n", err)
		return nil
	}

	return reactions
}
