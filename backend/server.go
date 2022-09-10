package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type HiScores struct {
	HiScores [10]HiScore `json:"hiscore"`
}

type HiScore struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func GetHiScores() HiScores {
	file, _ := ioutil.ReadFile("./backend/leaderboards.json")
	data := HiScores{}

	_ = json.Unmarshal([]byte(file), &data)

	return data
}

func SubmitScore(name string, score int) {
	leaderboards := GetHiScores()

	for i := 0; i < len(leaderboards.HiScores); i++ {
		if score > leaderboards.HiScores[i].Score {
			leaderboards.HiScores[i] = HiScore{name, score}
			break
		}
	}

	file, _ := json.MarshalIndent(leaderboards.HiScores, "", " ")
	_ = ioutil.WriteFile("./backend/leaderboards.json", file, 0644)
	fmt.Println("Score submitted successfully!")
}
