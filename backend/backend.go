package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Scores struct {
	Scores []PlayerData `json:"scores"`
}

type PlayerData struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Open our jsonFile
	jsonFile, err := os.Open("db.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened db.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data Scores

	json.Unmarshal(byteValue, &data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
