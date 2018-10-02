package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"github.com/gorilla/handlers"
	"os"
	"crypto/tls"
)

const (
	STATIC_DIR = "/static/"
)


func main() {

	//todo: remove it
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	router := mux.NewRouter()

	beerController := BeerController{}
	beerController.Initialize()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "."+STATIC_DIR+"index.html")
	})
	router.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))

	router.HandleFunc("/public/beers", beerController.GetAllBeerPublic).Methods("GET")
	router.HandleFunc("/secured/beers", beerController.GetAllBeer).Methods("GET")

	//
	// CORS + Logging handlers
	//
	allowedHeaders := handlers.AllowedHeaders([]string{"*"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	corsRouter := handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)
	corsLoggedRouter := handlers.LoggingHandler(os.Stdout, corsRouter)

	log.Fatal(http.ListenAndServe(":8080", corsLoggedRouter))
}