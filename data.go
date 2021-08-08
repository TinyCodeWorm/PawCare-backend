package main

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

var PhotoTypes = map[string]string{
	".jpeg": "image",
	".jpg":  "image",
	".gif":  "image",
	".png":  "image",
}

type Pet struct {
	Photourl string `json:"photourl"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Weight   string `json:"weight"`
	AgeYear  string `json:"ageyear"`
	AgeMonth string `json:"agemonth"`
	Sex      string `json:"sex"`
	Breed    string `json:"breed"`
}

type Food struct {
	FoodName    string `json:"foodname"`
	Brand       string `json:"brand"`
	Ingredient1 string `json:"ingredient1"`
	Ingredient2 string `json:"ingredient2"`
	Ingredient3 string `json:"ingredient3"`
	Ingredient4 string `json:"ingredient4"`
	Ingredient5 string `json:"ingredient5"`
	Ingredient6 string `json:"ingredient6"`
}

type Reaction struct {
	Name        string `json:"reaction_name"`
	Description string `json:"reaction_description"`
}

const userMapping = `{
	"mappings": {
		"properties": {
			"email":       { "type": "keyword" },
			"firstname":   { "type": "keyword", "index": false },
			"lastname":    { "type": "keyword", "index": false },
			"password":    { "type": "keyword" }
		}
	}
}`

const petMapping = `{
	"mappings": {
		"properties": {
			"owner_email":   { "type": "keyword" },
			"name":          { "type": "keyword" },
			"photo":         { "type": "keyword", "index": false },
			"type":          { "type": "keyword"},
			"weight":        { "type": "keyword", "index": false },
			"ageyear":       { "type": "keyword", "index": false },
			"agemonth":      { "type": "keyword", "index": false },
			"sex":           { "type": "keyword", "index": false },
			"breed":         { "type": "keyword", "index": false }
		}
	}
}`

const foodMapping = `{
	"mappings": {
		"properties": {
			"pet_id":         { "type": "integer" },
			"name":           { "type": "keyword" },
			"brand":          { "type": "keyword"},
			"ingredient1":    { "type": "keyword"},
			"ingredient2":    { "type": "keyword"},
			"ingredient3":    { "type": "keyword"},
			"ingredient4":    { "type": "keyword" },
			"ingredient5":    { "type": "keyword"},
			"ingredient6":    { "type": "keyword"}
		}
	}
}`

const reactionMapping = `{
	"mappings": {
		"properties": {
			"name":         { "type": "keyword" },
			"description":  { "type": "keyword" }
		}
	}
}`

const petreactionMpping = `{
	"mappings": {
		"properties": {
			"owner_email":                   { "type": "keyword" },
			"pet_name":                      { "type": "keyword" },
			"reaction_name":                 { "type": "keyword" },
			"food_name":                     { "type": "keyword" },
			"reaction_date":                 { "type": "date" }
		}
	}
}`
