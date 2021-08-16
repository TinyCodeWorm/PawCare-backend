package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/pborman/uuid"
)

var mySigningKey = []byte("secret")

func signinHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one signin request")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	//  Get User information from client
	decoder := json.NewDecoder(r.Body)
	var user User
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Cannot decode user data from client", http.StatusBadRequest)
		fmt.Printf("Cannot decode user data from client %v\n", err)
		return
	}

	exists, err := checkUser(user.Email, user.Password)
	if err != nil {
		http.Error(w, "Failed to read user from Elasticsearch", http.StatusInternalServerError)
		fmt.Printf("Failed to read user from Elasticsearch %v\n", err)
		return
	}

	if !exists {
		http.Error(w, "User doesn't exists or wrong password", http.StatusUnauthorized)
		fmt.Printf("User doesn't exists or wrong password\n")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		fmt.Printf("Failed to generate token %v\n", err)
		return
	}

	w.Write([]byte(tokenString))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one signup request")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var user User
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Cannot decode user data from client", http.StatusBadRequest)
		fmt.Printf("Cannot decode user data from client %v\n", err)
		return
	}

	if user.Email == "" || user.Password == "" || regexp.MustCompile(`^[a-z0-9]$`).MatchString(user.Email) {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		fmt.Printf("Invalid username or password\n")
		return
	}

	success, err := addUser(&user)
	if err != nil {
		http.Error(w, "Failed to save user to Elasticsearch", http.StatusInternalServerError)
		fmt.Printf("Failed to save user to Elasticsearch %v\n", err)
		return
	}

	if !success {
		http.Error(w, "User already exists", http.StatusBadRequest)
		fmt.Println("User already exists")
		return
	}
	fmt.Printf("user added successfully: %s.\n", user.Email)
}

func getprofileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one signup request")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	profile := getProfile(w, r)

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, "Failed to load my Profile", http.StatusInternalServerError)
		fmt.Printf("Failed to parse profile into JSON format %v.\n", err)
		return
	}
	w.Write(js)
}

func getpetsHandler(w http.ResponseWriter, r *http.Request) {
}

func uploadpetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one upload request")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	user := r.Context().Value("user")
	claims := user.(*jwt.Token).Claims
	useremail := claims.(jwt.MapClaims)["email"]

	myESPet := esPet{
		PetID:      uuid.New(),
		OwnerEmail: useremail.(string),
		Name:       r.FormValue("name"),
		Type:       r.FormValue("type"),
		Weight:     r.FormValue("weight"),
		AgeYear:    r.FormValue("ageyear"),
		AgeMonth:   r.FormValue("agemonth"),
		Sex:        r.FormValue("sex"),
		Breed:      r.FormValue("breed"),
	}

	file, header, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Media file is not available", http.StatusBadRequest)
		fmt.Printf("Media file is not available %v\n", err)
		return
	}

	suffix := filepath.Ext(header.Filename)
	if _, ok := PhotoTypes[suffix]; !ok {
		fmt.Printf("Photo format is not supported %v\n", err)
	}

	err = savePet(&myESPet, file)
	if err != nil {
		http.Error(w, "Failed to save post to GCS or Elasticsearch", http.StatusInternalServerError)
		fmt.Printf("Failed to save post to GCS or Elasticsearch %v\n", err)
		return
	}

	fmt.Println("Pet is saved successfully.")
}

func getfoodsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one getfoods request")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	user := r.Context().Value("user")
	claims := user.(*jwt.Token).Claims
	useremail := claims.(jwt.MapClaims)["email"].(string)

	allFood, err := getFoods(w, useremail)
	if err != nil {
		http.Error(w, "Failed to parse foods into JSON format", http.StatusInternalServerError)
		fmt.Printf("Failed to parse foods into JSON format %v.\n", err)
	}

	js, err := json.Marshal(allFood)
	if err != nil {
		http.Error(w, "Failed to parse foods into JSON format", http.StatusInternalServerError)
		fmt.Printf("Failed to parse foods into JSON format %v.\n", err)
		return
	}

	w.Write(js)
}

func uploadfoodHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one upload request")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	user := r.Context().Value("user")
	claims := user.(*jwt.Token).Claims
	useremail := claims.(jwt.MapClaims)["email"]

	var food esFood
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&food); err != nil {
		http.Error(w, "Cannot decode food data from client", http.StatusBadRequest)
		fmt.Printf("Cannot decode food data from client %v\n", err)
		return
	}

	food.OwnerEmail = useremail.(string)

	if err := saveToES(food, FOOD_INDEX, ""); err != nil {
		http.Error(w, "Cannot save food data from client", http.StatusBadRequest)
		fmt.Printf("Cannot save food data from client %v\n", err)
	}

	fmt.Println("food is saved successfully.")
}

func getreactionsHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received one getreactions request")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	reactions := getAllReactions(w)

	js, err := json.Marshal(reactions)
	if err != nil {
		http.Error(w, "Failed to parse reaction into JSON format", http.StatusInternalServerError)
		fmt.Printf("Failed to parse reaction into JSON format %v.\n", err)
		return
	}
	w.Write(js)

}

func getpetreactionsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one getprofile request")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	user := r.Context().Value("user")
	claims := user.(*jwt.Token).Claims
	useremail := claims.(jwt.MapClaims)["email"].(string)

	petreactions, err := getPetReactions(w, useremail)
	if err != nil {
		http.Error(w, "Failed to parse petreactions into JSON format", http.StatusInternalServerError)
		fmt.Printf("Failed to parse petreactions into JSON format %v.\n", err)
	}

	js, err := json.Marshal(petreactions)
	if err != nil {
		http.Error(w, "Failed to parse petreactions into JSON format", http.StatusInternalServerError)
		fmt.Printf("Failed to parse petreactions into JSON format %v.\n", err)
		return
	}

	w.Write(js)
}

func uploadpetreactionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one upload request")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	user := r.Context().Value("user")
	claims := user.(*jwt.Token).Claims
	useremail := claims.(jwt.MapClaims)["email"].(string)

	var petrea Petrea
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&petrea); err != nil {
		http.Error(w, "Cannot decode Pet Reaction data from client", http.StatusBadRequest)
		fmt.Printf("Cannot decode Pet Reaction data from client %v\n", err)
		return
	}

	uploadPetRea(w, petrea, useremail)

}

func getallergensHandler(w http.ResponseWriter, r *http.Request) {
}

func getbreedsHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received one getallbreeds request")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	//  Get User information from client
	decoder := json.NewDecoder(r.Body)
	var breed Breed
	if err := decoder.Decode(&breed); err != nil {
		http.Error(w, "Cannot decode breed data from client", http.StatusBadRequest)
		fmt.Printf("Cannot decode breed data from client %v\n", err)
		return
	}

	breeds := getAllBreeds(w, breed.Species)

	js, err := json.Marshal(breeds)
	if err != nil {
		http.Error(w, "Failed to parse breed into JSON format", http.StatusInternalServerError)
		fmt.Printf("Failed to parse breed into JSON format %v.\n", err)
		return
	}

	w.Write(js)

}
