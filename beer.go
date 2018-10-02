package main

import (
	"net/http"
	"encoding/json"
	verifier "github.com/okta/okta-jwt-verifier-golang"
	"fmt"
	"strings"
)

const (
	SPA_CLIENT_ID = "0oagfyv9iuDw0rDri0h7"
	CLIENT_SECRET = "4y-8ZMssEm0RDq_1XNOIT5fZIywZ4VfQwwiRf9Pe"
	ISSUER = "https://identity-np.swissre.com/oauth2/default"
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
func (c *BeerController) GetAllBeerPublic(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c.beers)
}

func (c *BeerController) GetAllBeer(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		return
	}

	fmt.Println("/secured/beers called")

	if !isAuthenticated(r) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - You are not authorized for this request"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c.beers)

}

func isAuthenticated(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")

	fmt.Println("authHeader " + authHeader) //todo: delete that

	if authHeader == "" {
		return false
	}

	tokenParts := strings.Split(authHeader, "Bearer ")
	bearerToken := tokenParts[1]

	fmt.Println("bearerToken " + bearerToken)

	tv := map[string]string{}
	tv["aud"] = "api://default"
	tv["cid"] = SPA_CLIENT_ID
	jv := verifier.JwtVerifier{
		Issuer:           ISSUER,
		ClaimsToValidate: tv,
	}

	_, err := jv.New().VerifyAccessToken(bearerToken)

	fmt.Printf("VerifyAccessToken error %s", err)
	if err != nil {
		return false
	}

	return true
}

