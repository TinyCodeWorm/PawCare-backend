package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

const (
	USER_INDEX        = "pawcare_user"
	PET_INDEX         = "pet"
	FOOD_INDEX        = "food"
	REACTION_INDEX    = "reaction"
	PETREACTION_INDEX = "pet_reaction"
)

func readFromES(query elastic.Query, index string) (*elastic.SearchResult, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth(ES_USERNAME, ES_PASSWORD))
	if err != nil {
		return nil, err
	}

	searchResult, err := client.Search().
		Index(index).
		Query(query).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}

func saveToES(i interface{}, index string, id string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth(ES_USERNAME, ES_PASSWORD))
	if err != nil {
		return err
	}

	_, err = client.Index().
		Index(index).
		Id(id).
		BodyJson(i).
		Do(context.Background())
	return err
}

func deleteFromES(query elastic.Query, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth(ES_USERNAME, ES_PASSWORD))
	if err != nil {
		return err
	}

	_, err = client.DeleteByQuery().
		Index(index).
		Query(query).
		Pretty(true).
		Do(context.Background())

	return err
}

func setupDB() {

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth(ES_USERNAME, ES_PASSWORD))
	if err != nil {
		panic(err)
	}

	createIndex(client, USER_INDEX, userMapping)
	createIndex(client, PET_INDEX, petMapping)
	createIndex(client, FOOD_INDEX, foodMapping)
	createIndex(client, REACTION_INDEX, reactionMapping)
	createIndex(client, PETREACTION_INDEX, petreactionMpping)

	var arrayReaction = [6][2]string{
		{"No Reaction", "description"},
		{"Diarrhea", "description"},
		{"Ear Infection", "description"},
		{"Scratching", "description"},
		{"Sneezing", "description"},
		{"Vomiting", "description"},
	}

	for _, v := range arrayReaction {
		reaction := Reaction{
			Name:        v[0],
			Description: v[1],
		}
		saveToES(&reaction, REACTION_INDEX, "")
		fmt.Println(" add reaction ")
	}

}

func createIndex(client *elastic.Client, index string, mapping string) {
	index_exists, err := client.IndexExists(index).Do(context.Background())
	if err != nil {
		panic(err)
	}

	if index_exists {
		if index != REACTION_INDEX {
			return
		}

		fmt.Println(index + " index are existed,will be deleting..")
		deleteIndex(client, index)
	}

	_, err = client.CreateIndex(index).Body(mapping).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(index + " are created.")

}

func deleteIndex(client *elastic.Client, index string) {
	deleteIndex, err := client.DeleteIndex(index).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !deleteIndex.Acknowledged {
		// Not acknowledged
		fmt.Println("delete index is not Acknowledged.")
	}
}
