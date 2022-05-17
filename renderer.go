package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

func draw(g *Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.DarkGray)
	if !inShop {
		rl.DrawTexture(g.bg, 0, 0, rl.White) // Draw background
		// Draw player character
		if !g.gameOver {
			rl.DrawTextureRec(g.char, g.playerRec, g.player.position, rl.White)
		} else {
			rl.DrawTexture(g.dead, int32(g.player.position.X-50), int32(g.player.position.Y-60), rl.White)
		}
		// Draw enemies
		for i := 0; i < MaxEnemies; i++ {
			if g.enemy[i].active {
				switch g.enemy[i].armed {
				case true:
					rl.DrawTextureRec(g.armedEnemyTexture, g.enemyRec, g.enemy[i].position, rl.White)
					break
				case false:
					rl.DrawTextureRec(g.enemyTexture, g.enemyRec, g.enemy[i].position, rl.White)
					break
				}
			} else {
				rl.DrawTextureRec(g.splatter, g.splatterRec, g.enemy[i].position, rl.White)
			}
		}
		// Draw bullets
		for i := 0; i < MaxBullets; i++ {
			if g.bullet[i].active {
				rl.DrawRectangleRec(g.bullet[i].rec, g.bullet[i].Color)
				if g.bullet[i].rec.Y < g.player.position.Y+140 {
					rl.DrawCircle(int32(g.player.position.X+41.5), int32(g.player.position.Y+112.5), 15, rl.Orange)
				}
			}
		}
		// Draw ammo
		rl.DrawTexture(g.bulletTex, screenWidth-55, screenHeight-100, rl.RayWhite)
		if !g.player.reloading {
			if g.gun[cWeapon].ammo == 0 {
				rl.DrawText(strconv.Itoa(g.gun[cWeapon].ammo), screenWidth-35, screenHeight-30, 20, rl.Red)
			} else {
				rl.DrawText(strconv.Itoa(g.gun[cWeapon].ammo), screenWidth-35, screenHeight-30, 20, rl.Black)
			}
		} else {
			rl.DrawText("reloading", screenWidth-100, screenHeight-30, 20, rl.Red)
		}

		rl.DrawTexture(g.gun[cWeapon].gunIcon, screenWidth-325, screenHeight-100, rl.White)

		// Game Over screen
		if g.gameOver {
			//rl.DrawTexture(g.deathScreen, 0, 0, rl.White)
			rl.DrawText("Mission Failed!", screenWidth/2-280, screenHeight/2, 80, rl.Black)
			rl.DrawText("Press Enter to retry", screenWidth/2-220, screenHeight/2+100, 40, rl.Violet)
		}

		// Draw hearts
		for i := 0; i <= g.player.lives; i++ {
			switch i {
			case 1:
				rl.DrawTexture(g.heart, screenWidth-150, 0, rl.RayWhite)
				break
			case 2:
				rl.DrawTexture(g.heart, screenWidth-100, 0, rl.RayWhite)
				break
			case 3:
				rl.DrawTexture(g.heart, screenWidth-50, 0, rl.RayWhite)
				break
			default:
				break
			}
		}

		if score < 0 {
			rl.DrawText(strconv.Itoa(score), 20, screenHeight-120, 40, rl.Maroon)
		} else {
			rl.DrawText(strconv.Itoa(score), 20, screenHeight-120, 40, rl.White)
		}

		rl.DrawText("Kills: "+strconv.Itoa(kills), 280, screenHeight-60, 40, rl.SkyBlue)
		rl.DrawText(" / "+strconv.Itoa(killsRequired), 420, screenHeight-60, 40, rl.SkyBlue)

	} else { // Draw shop screen
		rl.DrawTexture(g.shopScreen, 0, 0, rl.White)

		// Draw UI buttons and their functionality implemented (band-aid solution, I know)
		for i := 0; i < 4; i++ {
			if !g.gun[i+1].inInventory {
				rl.DrawRectangle(int32(screenWidth/4*i+40), screenHeight/3*2, int32(g.button[i].size.X), int32(g.button[i].size.Y+15), rl.DarkGray)
				rl.DrawTexture(*g.button[i].icon, int32(screenWidth/4*i+45), screenHeight/3*2, rl.Black)
				rl.DrawText(g.gun[i+1].name+" $"+strconv.Itoa(g.gun[i+1].price), int32(screenWidth/4*i+60), screenHeight*0.75, 25, rl.Green)
			}
		}
		rl.DrawText("Press enter to exit shop", screenWidth/3, screenHeight*0.9, 40, rl.SkyBlue)
	}

	rl.DrawText(strconv.Itoa(money)+"$", 20, screenHeight-60, 40, rl.Green)

	rl.DrawFPS(5, 0)
	rl.EndDrawing()
}
