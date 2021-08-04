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
	Date     string `json:"date"`
	Foodname string `json:"foodname"`
	Reaction string `json:"reaction"`
}
