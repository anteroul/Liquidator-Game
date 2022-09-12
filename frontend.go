package main

import (
	"encoding/json"
	"io/ioutil"
)

type PlayerData struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func submitScore(name string, score int) {
	data := []PlayerData{
		{"xX_PussyMolester69", 38550},
		{"Gaymer666", 27830},
		{"asdfgh", 26650},
		{name, score},
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("./backend/db.json", file, 0644)
}
