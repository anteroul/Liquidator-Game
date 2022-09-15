package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
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
	sort.Slice(responseObject, func(i, j int) bool {
		return responseObject[i].Score > responseObject[j].Score
	})
	return responseObject
}

func SubmitNewHiScore(name string, score int) bool {
	hs := HiScore{name, score}
	leaderboards := ReadHiScores()
	if len(leaderboards) > 9 {
		if leaderboards[9].Score > hs.Score {
			fmt.Println("You didn't make it to the hi-scores.")
			return false
		}
	}
	data, _ := json.Marshal(hs)
	request, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return false
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return true
}
