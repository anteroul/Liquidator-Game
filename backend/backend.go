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
			Score: 12850,
		},
		PlayerData{
			Name:  "Joe",
			Score: 8430,
		},
		PlayerData{
			Name:  "Tyler",
			Score: 7690,
		},
		PlayerData{
			Name:  "Gabriel",
			Score: 5440,
		},
		PlayerData{
			Name:  "Bob",
			Score: 4800,
		},
		PlayerData{
			Name:  "Ryder",
			Score: 4750,
		},
		PlayerData{
			Name:  "Jen",
			Score: 3800,
		},
		PlayerData{
			Name:  "Dan",
			Score: 2100,
		},
		PlayerData{
			Name:  "Sam",
			Score: 1950,
		},
	)
	r.HandleFunc("/", getHiScores).Methods("GET")
	r.HandleFunc("/", submitHiScore).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
