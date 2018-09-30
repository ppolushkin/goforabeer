package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

type BeerController struct {
	beers []Beer
}

type Beer struct {
	// the id for beer
	//
	// required: true
	// min: 1
	ID uint64 `json:"id"`
	// the name for beer
	// required: true
	// min length: 3
	Name string `json:"name"`
}

func (c *BeerController) Initialize() {
	c.beers = append(c.beers, Beer{ID: 1, Name: "Guiness extra cold"})
	c.beers = append(c.beers, Beer{ID: 2, Name: "Nevskoe svetloe"})
	c.beers = append(c.beers, Beer{ID: 3, Name: "Baltika"})
}

// swagger:operation GET /beers beers getAllBeer
// ---
// summary: Returns all beer in the bar
func (c *BeerController) GetAllBeer(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c.beers)
}

// swagger:operation GET /beers/{id} beers getBeerById
// ---
// summary: Returns beer by id
func (c *BeerController) GetBeerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := parseId(vars)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Beer ID")
		return
	}

	for _, item := range c.beers {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Beer{})
}

func parseId(vars map[string]string) (uint64, error) {
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	return id, err
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
