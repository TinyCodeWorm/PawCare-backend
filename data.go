package main

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type Profile struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

var PhotoTypes = map[string]string{
	".jpeg": "image",
	".jpg":  "image",
	".gif":  "image",
	".png":  "image",
}

type Pet struct {
	Name     string `json:"name"`
	Photourl string `json:"photo"`
	Type     string `json:"type"`
	Weight   string `json:"weight"`
	AgeYear  string `json:"ageyear"`
	AgeMonth string `json:"agemonth"`
	Sex      string `json:"sex"`
	Breed    string `json:"breed"`
}

type esPet struct {
	PetID      string `json:"pet_id"`
	OwnerEmail string `json:"owner_email"`
	Name       string `json:"name"`
	Photourl   string `json:"photo"`
	Type       string `json:"type"`
	Weight     string `json:"weight"`
	AgeYear    string `json:"ageyear"`
	AgeMonth   string `json:"agemonth"`
	Sex        string `json:"sex"`
	Breed      string `json:"breed"`
}

func toEsPet(pet *Pet) esPet {
	return esPet{
		PetID:      "",
		OwnerEmail: "",
		Name:       pet.Name,
		Photourl:   pet.Photourl,
		Type:       pet.Type,
		Weight:     pet.Weight,
		AgeYear:    pet.AgeYear,
		AgeMonth:   pet.AgeMonth,
		Sex:        pet.Sex,
		Breed:      pet.Breed,
	}
}

func toPet(espet *esPet) Pet {
	return Pet{
		Photourl: espet.Photourl,
		Type:     espet.Type,
		Name:     espet.Name,
		Weight:   espet.Weight,
		AgeYear:  espet.AgeYear,
		AgeMonth: espet.AgeMonth,
		Sex:      espet.Sex,
		Breed:    espet.Breed,
	}
}

type Food struct {
	FoodName    string `json:"name"`
	Brand       string `json:"brand"`
	Ingredient1 string `json:"ingredient1"`
	Ingredient2 string `json:"ingredient2"`
	Ingredient3 string `json:"ingredient3"`
	Ingredient4 string `json:"ingredient4"`
	Ingredient5 string `json:"ingredient5"`
	Ingredient6 string `json:"ingredient6"`
}

type esFood struct {
	OwnerEmail  string `json:"owner_email"`
	FoodName    string `json:"name"`
	Brand       string `json:"brand"`
	Ingredient1 string `json:"ingredient1"`
	Ingredient2 string `json:"ingredient2"`
	Ingredient3 string `json:"ingredient3"`
	Ingredient4 string `json:"ingredient4"`
	Ingredient5 string `json:"ingredient5"`
	Ingredient6 string `json:"ingredient6"`
}

func toEsFood(food *Food) esFood {
	return esFood{
		OwnerEmail:  "",
		FoodName:    food.FoodName,
		Brand:       food.Brand,
		Ingredient1: food.Ingredient1,
		Ingredient2: food.Ingredient2,
		Ingredient3: food.Ingredient3,
		Ingredient4: food.Ingredient4,
		Ingredient5: food.Ingredient5,
		Ingredient6: food.Ingredient6,
	}
}

func toFood(esfood *esFood) Food {
	return Food{
		FoodName:    esfood.FoodName,
		Brand:       esfood.Brand,
		Ingredient1: esfood.Ingredient1,
		Ingredient2: esfood.Ingredient2,
		Ingredient3: esfood.Ingredient3,
		Ingredient4: esfood.Ingredient4,
		Ingredient5: esfood.Ingredient5,
		Ingredient6: esfood.Ingredient6,
	}
}

type Reaction struct {
	Name        string `json:"reaction_name"`
	Description string `json:"reaction_description"`
}

type Breed struct {
	Species string `json:"animal_specie"`
}

type esBreed struct {
	Species string `json:"animal_specie"`
	Name    string `json:"breed_name"`
}

type esPetReaction struct {
	OwnerEmail   string `json:"owner_email"`
	PetName      string `json:"pet_name"`
	FoodName     string `json:"food_name"`
	ReactionDate string `json:"reaction_date"`
	ReactionName string `json:"reaction_name"`
}

type PetReaction struct {
	FoodName     string `json:"food_name"`
	ReactionDate string `json:"reaction_date"`
	ReactionName string `json:"reaction_name"`
}

type Petrea struct {
	ReactionDate string   `json:"reaction_date"`
	FoodName     string   `json:"food_name"`
	Reactions    []string `json:"reaction_name"`
}

const userMapping = `{
	"mappings": {
		"properties": {
			"email":       { "type": "keyword" },
			"firstname":   { "type": "keyword" },
			"lastname":    { "type": "keyword" },
			"password":    { "type": "keyword" }
		}
	}
}`

const petMapping = `{
	"mappings": {
		"properties": {
			"pet_id":   	 { "type": "keyword" },
			"owner_email":   { "type": "keyword" },
			"name":          { "type": "keyword" },
			"photo":         { "type": "keyword", "index": false },
			"type":          { "type": "keyword"},
			"weight":        { "type": "keyword" },
			"ageyear":       { "type": "keyword" },
			"agemonth":      { "type": "keyword" },
			"sex":           { "type": "keyword"},
			"breed":         { "type": "keyword"}
		}
	}
}`

const foodMapping = `{
	"mappings": {
		"properties": {
			"owner_email":   { "type": "keyword" },
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
			"name":         { "type": "text" },
			"description":  { "type": "keyword" }
		}
	}
}`

const breedMapping = `{
	"mappings": {
		"properties": {
			"animal_specie":       { "type": "keyword" },
			"breed_name":          { "type": "keyword" }
		}
	}
}`

const petreactionMpping = `{
	"mappings": {
		"properties": {
			"owner_email":                   { "type": "keyword" },
			"pet_name":                      { "type": "keyword" },
			"food_name":                     { "type": "keyword" },
			"reaction_date":                 { "type": "date", "format": "yyyy-MM-dd" },
			"reaction_name":                 { "type": "keyword" }
		}
	}
}`
