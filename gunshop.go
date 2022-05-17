package main

import rl "github.com/gen2brain/raylib-go/raylib"

func enterShopScreen(g *Game) {
	for i := 0; i < 4; i++ {
		if !g.gun[i+1].inInventory {
			if onClickEvent(&rl.Rectangle{X: float32(int32(screenWidth/4*i + 40)), Y: screenHeight / 3 * 2, Width: float32(int32(g.button[i].size.X)), Height: float32(int32(g.button[i].size.Y + 15))}) {
				if money >= g.gun[i+1].price && !g.gun[i+1].inInventory {
					money -= g.gun[i+1].price
					g.gun[i+1].inInventory = true
				}
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		go exitShopScreen(g)
	}
}

func exitShopScreen(g *Game) {
	g.player.reloading = false
	kills = 0
	wave++
	killsRequired = GetEnemies()
	g.player.position = rl.NewVector2(float32(screenWidth)/2, 40)
	// Initialize bullets
	for i := 0; i < MaxBullets; i++ {
		g.bullet[i] = Bullet{rec: rl.Rectangle{X: g.player.position.X + 40, Y: g.player.position.Y + 105, Width: 5, Height: 10}, speed: rl.Vector2{Y: 15}, active: false, Color: rl.Yellow}
	}
	// Initialize enemies
	for i := 0; i < MaxEnemies; i++ {
		g.enemy[i] = Enemy{rl.Vector2{X: float32(rl.GetRandomValue(0, screenWidth-100)), Y: float32(rl.GetRandomValue(screenHeight, screenHeight+1000))}, false, true, 3}
	}
	inShop = false
}
