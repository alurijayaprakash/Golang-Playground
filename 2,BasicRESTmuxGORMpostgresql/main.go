package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// -------------------
type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var DB *gorm.DB
var err error

const DNS = "postgres://postgres:root@localhost:5432/godb"

// username : postgres
// pw : root
// db name : godb

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{})
}

// GET
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)

}

// GET
func GetUserbyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)

}

// POST
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)

}

// PUT  //doubt
func UpdateUserbyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)         // getting id
	var user User                 // to save user from db we are creating user struct
	DB.First(&user, params["id"]) // fetching value by id
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

// DELETE
func DeleteUserbyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("User Deleted")
}

// --------------

func initializeRouter() {
	var r = mux.NewRouter() // creating new router from mux pkg

	//routes
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUserbyId).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUserbyId).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUserbyId).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}

func main() {
	fmt.Println("App started at Port no : 8080 ....")
	InitialMigration()
	initializeRouter()

}
