package main

// //  with out DB this rest api build
// // ref : https://www.youtube.com/watch?v=TkbhQQS3m_o

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// type Movie struct {
// 	ID       string    `json:"id"`
// 	Isbn     string    `json:"isbn"`
// 	Title    string    `json:"title"`
// 	Director *Director `json:"director"` // Director pointer here pointing to Director struct below
// }

// type Director struct {
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// }

// var movies []Movie // list of movies

// func getMovies(w http.ResponseWriter, r *http.Request) { // http.Request should be dereferencer
// 	fmt.Println("Entered getMovies")

// 	w.Header().Set("Content-Type", "application/json")
// 	// before sending we need to encode response w
// 	json.NewEncoder(w).Encode(movies) // sending process
// }

// func deleteMovie(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Entered deleteMovie")

// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range movies {
// 		if item.ID == params["id"] {
// 			movies = append(movies[:index], movies[index+1:]...)
// 			json.NewEncoder(w).Encode(item) // printing deleted item
// 			break
// 		}
// 	}

// }

// func getMoviebyId(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Entered getMoviebyId")

// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for _, item := range movies {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			break
// 		}
// 	}
// }

// // not working
// func createMovie(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Entered createMovie")
// 	w.Header().Set("Content-Type", "application/json")
// 	var movie Movie
// 	// we need to decode that request body and item
// 	_ = json.NewDecoder(r.Body).Decode(&movie) // must pass pointer
// 	movie.ID = strconv.Itoa(rand.Intn(1000000))
// 	// we saved ID as string, strconv help us to convert int to string
// 	movies = append(movies, movie)
// 	fmt.Println(movie)
// 	json.NewEncoder(w).Encode(movie)
// }

// // func updateMovie(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Println("Entered updateMovie")
// // 	w.Header().Set("Content-Type", "application/json")
// // 	var New_Movie Movie
// // 	_ = json.NewDecoder(r.Body).Decode(&New_Movie)
// // 	// based on id we are updating values in the movie
// // 	for index, item := range movies {
// // 		if item.ID == New_Movie.ID {
// // 			item.Title = New_Movie.Title
// // 			item.Isbn = New_Movie.Isbn

// // 			item.Director.FirstName = New_Movie.Director.FirstName
// // 			item.Director.LastName = New_Movie.Director.LastName

// // 			movies[index] = item // update item ie. movie in same index // imp
// // 			json.NewEncoder(w).Encode(item)
// // 			break
// // 		}
// // 	}

// // }

// func updateMovie(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Entered updateMovie")
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	// loop over the movies by range
// 	// delete the movie with id (sent by the paras of the postman)
// 	// add new movie (sent by the body of the postman)

// 	for index, item := range movies {
// 		if item.ID == params["id"] {
// 			movies = append(movies[:index], movies[index+1:]...)
// 			var New_Movie Movie
// 			json.NewDecoder(r.Body).Decode(&New_Movie)
// 			New_Movie.ID = params["id"]
// 			movies = append(movies, New_Movie)
// 			json.NewEncoder(w).Encode(New_Movie)
// 			return
// 		}
// 	}

// }

// func main() {
// 	var r = mux.NewRouter() // router by mux
// 	fmt.Println("Hello")

// 	//sample movies added to list/struct
// 	movies = append(movies, Movie{ID: "1", Isbn: "23456", Title: "Movie One", Director: &Director{FirstName: "Jhon", LastName: "Doe"}})
// 	movies = append(movies, Movie{ID: "2", Isbn: "34667", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})
// 	movies = append(movies, Movie{ID: "3", Isbn: "67342", Title: "Movie Three", Director: &Director{FirstName: "Devi", LastName: "DSP"}})

// 	// routers
// 	r.HandleFunc("/movies", getMovies).Methods("GET")
// 	r.HandleFunc("/movies/{id}", getMoviebyId).Methods("GET")
// 	r.HandleFunc("/movies", createMovie).Methods("POST")
// 	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
// 	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

// 	fmt.Println("Starting server at port 8080...")
// 	log.Fatal(http.ListenAndServe(":8080", r))

// }
