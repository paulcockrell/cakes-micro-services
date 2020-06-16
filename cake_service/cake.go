package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Cake - definition of a cake record
type Cake struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Comment   string `json:"comment"`
	ImageURL  string `json:"image_url"`
	YumFactor int    `json:"yum_factor"`
}

// Cakes - In-memory store for our cakes
var Cakes []Cake

func findHighestCakeID() int {
	maxID := Cakes[0].ID
	for _, cake := range Cakes {
		if cake.ID > maxID {
			maxID = cake.ID
		}
	}

	return maxID
}

func allCakes(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Cakes)
}

func getCake(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, cake := range Cakes {
		if cake.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cake)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Cake not found"})
}

func newCake(w http.ResponseWriter, r *http.Request) {
	cake := Cake{}
	err := json.NewDecoder(r.Body).Decode(&cake)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid cake data"})
		return
	}

	cake.ID = findHighestCakeID() + 1
	Cakes = append(Cakes, cake)
	json.NewEncoder(w).Encode(cake)
}

func updateCake(w http.ResponseWriter, r *http.Request) {
	cake := Cake{}
	err := json.NewDecoder(r.Body).Decode(&cake)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid cake data"})
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	cake.ID = id

	for index, c := range Cakes {
		if c.ID == cake.ID {
			Cakes = append(Cakes[:index], Cakes[index+1:]...)
			Cakes = append(Cakes, cake)
			json.NewEncoder(w).Encode(cake)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Cake not found"})
}

func deleteCake(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, cake := range Cakes {
		if cake.ID == id {
			Cakes = append(Cakes[:index], Cakes[index+1:]...)
			json.NewEncoder(w).Encode(cake)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Cake not found"})
}
