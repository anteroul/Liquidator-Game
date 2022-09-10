package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type HiScore struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

var leaderboards [10]HiScore

func SubmitScore(name string, score int) {
	var index = -1

	for i := 0; i < len(leaderboards); i++ {
		if score > leaderboards[i].Score {
			leaderboards[i] = HiScore{name, score}
			index = i
			break
		}
	}

	if index != -1 {
		file, _ := json.MarshalIndent(leaderboards[index], "", " ")
		_ = ioutil.WriteFile("./backend/leaderboards.json", file, 0644)
	}

	fmt.Println("Score submitted successfully")
}
