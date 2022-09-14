package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HiScore struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func ReadHiScores() []HiScore {
	response, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var responseObject []HiScore
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}

func SubmitNewHiScore(name string, score int) {
	var hs = HiScore{name, score}
	data, _ := json.Marshal(hs)
	request, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
}
