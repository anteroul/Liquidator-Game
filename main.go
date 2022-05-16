package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Main() {
	Setup() // Game Settings
	LaunchGame()
	if !rl.WindowShouldClose() && rl.IsWindowReady() {
		rl.CloseWindow()
		fmt.Errorf(string(rune(rl.LogError)))
		return
	}
	fmt.Println("Game ran and closed successfully.")
	rl.CloseWindow()
}
