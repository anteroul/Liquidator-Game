package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HiScore struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func ReadHiScores() []HiScore {
	response, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject []HiScore
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}

func SubmitNewHiScore(name string, score int) {
	hiScores := ReadHiScores()
	data := append(hiScores, HiScore{name, score})
	jsonData, _ := json.Marshal(data)
	response, _ := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(jsonData))
	var res map[string]interface{}
	json.NewDecoder(response.Body).Decode(&res)
	fmt.Println(res["json"])
}
