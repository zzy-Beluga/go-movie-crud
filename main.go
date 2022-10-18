package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"Id"`
	Isbn     string    `json:"Isbn"`
	Title    string    `json:"Title"`
	Director *Director `json:"Director"`
}

type Director struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetent-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	lid, _ := strconv.Atoi(movies[len(movies)-1].Id)
	movie.Id = strconv.Itoa(lid + 1)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{Id: "1", Isbn: "1234", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Wick"}})
	movies = append(movies, Movie{Id: "2", Isbn: "2234", Title: "Movie Two", Director: &Director{Firstname: "Jason", Lastname: "Mory"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server Starting at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
