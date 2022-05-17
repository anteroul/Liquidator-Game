package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Button struct {
	size rl.Vector2
	icon *rl.Texture2D
	gun  *Gun
}

// An event listener
func onClickEvent(rectangle *rl.Rectangle) bool {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), *rectangle) && rl.IsMouseButtonPressed(0) {
		return true
	} else {
		return false
	}
}
