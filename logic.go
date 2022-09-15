package main

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) update() {
	if !inShop {

		if !g.pause {

			framesCounter++

			if kills >= killsRequired && g.player.lives > 0 {
				inShop = true
			}

			if framesCounter%4 == 0 || framesCounter == 0 {
				enemyFrame++
				if g.player.reloading {
					reloadCounter++
				}
				if enemyFrame > 4 {
					enemyFrame = 0
				}
			}

			updatePlayerLogic(g)
			updateBullets(g) // Update shooting logic
			updateEnemy(g)   // Update enemy logic
			// Update controls if player is alive
			if !g.gameOver {
				keyCallback(g) // Game controls
			}
		}
	} else {
		kills = 0
		enterShopScreen(g)
		if rl.IsKeyPressed(rl.KeyEnter) {
			exitShopScreen(g)
			inShop = false
		}
	}
	draw(g) // Display graphics
}
