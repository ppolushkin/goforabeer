package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hii!! Let's Go for a beer!")
		fmt.Fprint(w, "Hii!! Let's Go for a Beer!")
	})

	http.ListenAndServe(":8080", nil)
}