package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	Setup() // Game Settings
	LaunchGame()
}

func Setup() {
	var gameShouldLaunch = false

	rl.InitWindow(800, 400, "Settings")
	rl.SetTargetFPS(60)

	for !gameShouldLaunch {
		if rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressed(rl.KeyDown) {
			if !enableFullScreen {
				enableFullScreen = true
			} else {
				enableFullScreen = false
			}
		}
		if rl.IsKeyPressed(rl.KeyEnter) {
			gameShouldLaunch = true
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawText("Fullscreen enabled", 300, 100, 20, rl.Black)
		rl.DrawText("Fullscreen disabled", 300, 200, 20, rl.Black)
		rl.DrawText("Press Enter to launch game", 250, 300, 20, rl.Magenta)

		switch enableFullScreen {
		case true:
			rl.DrawRectangle(200, 100, 20, 20, rl.Red)
			break
		case false:
			rl.DrawRectangle(200, 200, 20, 20, rl.Red)
			break
		}

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
