package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Name     string    `json:"name"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, value := range movies {
		if value.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, value := range movies {
		if value.ID == params["id"] {
			json.NewEncoder(w).Encode(value)
		}
	}

}

func createMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_=json.NewDecoder(r.Body).Decode(&movie)
	movie.ID=strconv.Itoa(rand.Intn(100000))
	movies=append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	var movie Movie
	for index,item:=range movies{
		if(item.ID==params["id"]){
			movies=append(movies[:index],movies[index+1:]...)
			_=json.NewDecoder(r.Body).Decode(&movie)
			movie.ID=params["id"]
			movies=append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{
		ID:    "1",
		ISBN:  "438227",
		Title: "THE LION KING",
		Director: &Director{
			Firstname: "NIKOLAS",
			Lastname:  "TESLA",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		ISBN:  "438228",
		Title: "SPIDER MAN",
		Director: &Director{
			Firstname: "MARTIN",
			Lastname:  "TESLA",
		},
	})
	movies = append(movies, Movie{
		ID:    "3",
		ISBN:  "438229",
		Title: "TITANIC",
		Director: &Director{
			Firstname: "FATRA",
			Lastname:  "TESLA",
		},
	})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/create", createMovie).Methods("POST")
	r.HandleFunc("movies/update/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/delete/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Server start at port:8080")
	http.ListenAndServe(":8080", r)
}
