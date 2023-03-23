package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var pl = fmt.Println
var pf = fmt.Printf

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var moviesArr []Movie

// error checker to handle errors
func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Response Headers are included with the data being sent back to the client
// to instruct the browser to do something

// A common requirement when creating HTTP servers is to be able to set headers on a response.
// Go offers great support for creating, reading, updating, and deleting headers.
// In the following example, suppose that the 250 server will send some JSON.
// By setting the Content-Type header, the server can inform the client that JSON data is being sent.
// Through the ResponseWriter, a handler function can add a header as follows:
// w.Header().Set("Content-Type","application/json")

func getMovies(w http.ResponseWriter, r *http.Request) {
	// here we are telling the client that a JSON value is being sent
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(moviesArr)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range moviesArr {
		if item.ID == params["id"] {
			moviesArr = append(moviesArr[:index], moviesArr[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range moviesArr {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	// decode the values send in the postman request and add it to the nwe movie variable
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100000))
	moviesArr = append(moviesArr, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	// steps -
	// 1. set the content type to json
	// 2. get the params
	// loop over the movies and -
	// 3. delete the id you want to update
	// 4. add the id sent in the API call

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range moviesArr {
		if item.ID == params["id"] {
			moviesArr = append(moviesArr[:index], moviesArr[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			moviesArr = append(moviesArr, movie)

			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	r := mux.NewRouter()

	moviesArr = append(moviesArr, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	moviesArr = append(moviesArr, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc(" /{id}", deleteMovie).Methods("DELETE")

	pf("Starting the server at port http://localhost:8000\n")
	err := http.ListenAndServe(":8000", r)
	errorCheck(err)
}
