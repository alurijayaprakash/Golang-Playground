// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	gorm.Model
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	Email     string `json:"email"`
// }

// var DB *gorm.DB
// var err error

// const DNS = "root:root@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

// func InitialMigration() {
// 	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("Cannot connect to DB")
// 	}
// 	DB.AutoMigrate(&User{})
// }

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// }

// func GetUserbyId(w http.ResponseWriter, r *http.Request) {

// }

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var user User
// 	json.NewDecoder(r.Body).Decode(&user)
// 	DB.Create(&user)
// 	json.NewEncoder(w).Encode(user)

// }

// func UpdateUserbyId(w http.ResponseWriter, r *http.Request) {

// }

// func DeleteUserbyId(w http.ResponseWriter, r *http.Request) {

// }
