package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func init() {
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "1234",
		Title: "Movie 1",
		Director: &Director{
			FirstName: "John",
			LastName:  "Doe",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "1235",
		Title: "Movie 2",
		Director: &Director{
			FirstName: "Jane",
			LastName:  "Doe",
		},
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("starting server at port 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	addHeader(w)
	id := mux.Vars(r)["id"]

	for _, movie := range movies {
		if movie.ID == id {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	addHeader(w)
	id := mux.Vars(r)["id"]

	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	addHeader(w)
	json.NewEncoder(w).Encode(movies)
}

func addHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}
