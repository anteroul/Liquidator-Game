package main

import rl "github.com/gen2brain/raylib-go/raylib"

func updateBullets(g *Game) {
	for i := 0; i < MaxBullets; i++ {
		if g.bullet[i].active {
			g.bullet[i].rec.Y += g.bullet[i].speed.Y
			// Collision with enemy
			for j := 0; j < MaxEnemies; j++ {
				if g.enemy[j].active && !g.gameOver {
					if rl.CheckCollisionRecs(g.bullet[i].rec, rl.Rectangle{X: g.enemy[j].position.X, Y: g.enemy[j].position.Y, Width: 90, Height: 40}) {
						g.enemy[j].active = false
						if !g.gun[cWeapon].armourPiercing {
							g.bullet[i].active = false
						}
						tangoDown(g)
					}
					if g.enemy[j].position.Y < 100 {
						if rl.CheckCollisionPointRec(rl.Vector2{X: g.bullet[i].rec.X, Y: g.bullet[i].rec.Y - 75}, rl.Rectangle{X: g.enemy[j].position.X, Y: g.enemy[j].position.Y, Width: 90, Height: 40}) {
							g.enemy[j].active = false
							if !g.gun[cWeapon].armourPiercing {
								g.bullet[i].active = false
							}
							tangoDown(g)
						}
					}
				}
			}
			if g.bullet[i].rec.Y+g.bullet[i].rec.Height >= screenHeight {
				g.bullet[i].active = false
			}
		}
	}
}
