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
			Name:  "xX_ProGamer12345",
			Score: 15850,
		},
		PlayerData{
			Name:  "Joe",
			Score: 12430,
		},
		PlayerData{
			Name:  "Tyler",
			Score: 9690,
		},
		PlayerData{
			Name:  "Gabriel",
			Score: 8440,
		},
		PlayerData{
			Name:  "Bob",
			Score: 7800,
		},
		PlayerData{
			Name:  "Ryder",
			Score: 7750,
		},
		PlayerData{
			Name:  "Jen",
			Score: 6800,
		},
		PlayerData{
			Name:  "Dan",
			Score: 5100,
		},
		PlayerData{
			Name:  "Sam",
			Score: 4950,
		},
		PlayerData{
			Name:  "Sam",
			Score: 3950,
		},
		PlayerData{
			Name:  "n00b",
			Score: 2950,
		},
	)
	r.HandleFunc("/", getHiScores).Methods("GET")
	r.HandleFunc("/", submitHiScore).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
