package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

const (
	STATIC_DIR = "/static/"
)


func main() {
	router := mux.NewRouter()

	beerController := BeerController{}
	beerController.Initialize()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "."+STATIC_DIR+"index.html")
	})
	router.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))

	router.HandleFunc("/public/beers", beerController.GetAllBeerPublic).Methods("GET")

	router.HandleFunc("/secured/beers", beerController.GetAllBeer).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}