package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type PlayerData struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

var hiScores []PlayerData

func getHiScores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hiScores)
}

func submitHiScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var score PlayerData
	_ = json.NewDecoder(r.Body).Decode(&score)
	hiScores = append(hiScores, score)
	json.NewEncoder(w).Encode(score)
}

func main() {
	r := mux.NewRouter()
	hiScores = append(hiScores,
		PlayerData{
			Name:  "valiant",
			Score: 98850,
		},
		PlayerData{
			Name:  "Joe",
			Score: 58430,
		},
		PlayerData{
			Name:  "Tyler",
			Score: 47690,
		},
		PlayerData{
			Name:  "Gabriel",
			Score: 45440,
		},
	)
	r.HandleFunc("/", getHiScores).Methods("GET")
	r.HandleFunc("/", submitHiScore).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
