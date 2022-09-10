package main

import rl "github.com/gen2brain/raylib-go/raylib"

func updateEnemy(g *Game) {
	// Enemy behaviour
	var isEnemiesSpawned = false
	// TODO: Shooting for armed enemies (armed enemies are currently disabled)
	for i := 0; i < MaxEnemies; i++ {
		if g.enemy[i].active {
			isEnemiesSpawned = true
			g.enemy[i].position.Y -= float32(rl.GetRandomValue(1, int32(g.enemy[i].speed)))
			// Crossing the border
			if g.enemy[i].position.Y < -120 {
				g.enemy[i].position.X = float32(rl.GetRandomValue(0, screenWidth-100))
				g.enemy[i].position.Y = float32(rl.GetRandomValue(screenHeight+200, screenHeight+1000))
				if !g.gameOver {
					score -= 500
					g.player.lives--
				}
			}
			// Collision with player
			if !g.gameOver {
				if rl.CheckCollisionRecs(rl.Rectangle{X: g.enemy[i].position.X, Y: g.enemy[i].position.Y, Width: 90, Height: 40}, rl.Rectangle{X: g.player.position.X, Y: g.player.position.Y, Width: 90, Height: 40}) {
					g.enemy[i].active = false
					g.player.lives--
					tangoDown(g)
				}
			}
			// Collision with barbed wire
			if rl.CheckCollisionRecs(rl.Rectangle{X: g.enemy[i].position.X, Y: g.enemy[i].position.Y, Width: 90, Height: 40}, g.barbedWire) {
				g.enemy[i].speed = 0.3
			} else {
				g.enemy[i].speed = 3
			}
		} else {
			g.splatterRec.X = float32(enemyFrame * g.enemyTexture.Width / 4)
			if enemyFrame >= 4 {
				g.enemy[i].position.X = float32(rl.GetRandomValue(0, screenWidth-100))
				g.enemy[i].position.Y = float32(rl.GetRandomValue(screenHeight+200, screenHeight+1000))
				if killsRequired-kills > MaxEnemies || !isEnemiesSpawned {
					g.enemy[i].active = true
					isEnemiesSpawned = true
				}
			}
		}
		updateEnemyRec(g, g.enemy[i])
	}
}

func tangoDown(g *Game) {
	rl.PlaySoundMulti(sfxDeath)
	score += 100
	if score >= 0 {
		money += 50
	}
	enemyFrame = 0
	if RandBool() && g.player.lives != PlayerMaxLife {
		g.player.lives++
	}
	kills++
}

func updateEnemyRec(g *Game, enemy Enemy) {
	if enemy.active {
		switch enemy.armed {
		case true:
			g.enemyRec.X = float32(enemyFrame * g.armedEnemyTexture.Width / 4)
			break
		case false:
			g.enemyRec.X = float32(enemyFrame * g.enemyTexture.Width / 4)
			break
		}
	}
}
