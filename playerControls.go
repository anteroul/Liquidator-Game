package main

import rl "github.com/gen2brain/raylib-go/raylib"

func updatePlayerLogic(g *Game) {
	// Reload logic
	if reloadCounter == 30 {
		g.player.reloading = false
		g.gun[cWeapon].ammo = g.gun[cWeapon].maxAmmo
		reloadCounter = 0
	}

	// Game over logic
	if g.player.lives <= 0 {
		g.gameOver = true
	}
}

func switchWeapon(g *Game, weaponIndex int) {
	if g.gun[weaponIndex].inInventory {
		g.player.reloading = false
		reloadCounter = 0
		cWeapon = weaponIndex
	}
}

func keyCallback(g *Game) {
	// Movement:
	if rl.IsKeyDown(rl.KeyRight) {
		if g.player.position.X+80 < screenWidth {
			g.player.position.X += g.player.speed
		}
		updateCharRec(g)
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		if g.player.position.X > 0 {
			g.player.position.X -= g.player.speed
		}
		updateCharRec(g)
	}
	if rl.IsKeyUp(rl.KeyLeft) && rl.IsKeyUp(rl.KeyRight) {
		g.playerRec.X = 0
	}
	// Semi-auto:
	if kills < killsRequired {
		if rl.IsKeyPressed(rl.KeySpace) {
			if g.gun[cWeapon].ammo > 0 && !g.player.reloading && !g.gun[cWeapon].automatic {
				for i := 0; i < MaxBullets; i++ {
					if !g.bullet[i].active {
						g.bullet[i].rec = rl.Rectangle{X: g.player.position.X + 40, Y: g.player.position.Y + 105, Width: 5, Height: 10}
						g.bullet[i].active = true
						g.gun[cWeapon].ammo--
						break
					}
				}
				if !g.gun[cWeapon].armourPiercing {
					rl.PlaySoundMulti(sfxRifle)
				} else {
					rl.PlaySoundMulti(sfxSniper)
				}
			}
		}
		// Full-auto:
		if rl.IsKeyDown(rl.KeySpace) {
			if g.gun[cWeapon].ammo > 0 && !g.player.reloading && g.gun[cWeapon].automatic {
				firingRateCounter++
				if firingRateCounter%g.gun[cWeapon].firingRate == 0 {
					for i := 0; i < MaxBullets; i++ {
						if !g.bullet[i].active {
							g.bullet[i].rec = rl.Rectangle{X: g.player.position.X + 40, Y: g.player.position.Y + 105, Width: 5, Height: 10}
							g.bullet[i].active = true
							g.gun[cWeapon].ammo--
							break
						}
					}
					if !g.gun[cWeapon].armourPiercing {
						rl.PlaySoundMulti(sfxRifle)
					} else {
						rl.SetSoundVolume(sfxGroza, 0.5)
						rl.PlaySoundMulti(sfxGroza)
					}
				}
			}
		}
	}
	// Reload keys:
	if rl.IsKeyPressed(rl.KeyR) {
		if g.gun[cWeapon].ammo != g.gun[cWeapon].maxAmmo {
			if !g.player.reloading {
				rl.PlaySound(sfxReload)
				g.player.reloading = true
			}
		}
	}
	// Weapon keys:
	if rl.IsKeyPressed(rl.KeyOne) {
		switchWeapon(g, 0) // Armalite
	}
	if rl.IsKeyPressed(rl.KeyTwo) {
		switchWeapon(g, 1) // Galil
	}
	if rl.IsKeyPressed(rl.KeyThree) {
		switchWeapon(g, 2) // Barrett
	}
	if rl.IsKeyPressed(rl.KeyFour) {
		switchWeapon(g, 3) // Groza
	}
	if rl.IsKeyPressed(rl.KeyFive) {
		switchWeapon(g, 4) // M60
	}
	if rl.IsKeyPressed(rl.KeyEnd) {
		g.player.lives = 0
	}
}

// Special keyboard events
func specialKeyCallback(g *Game) {
	// Pause game
	if rl.IsKeyPressed(rl.KeyP) {
		if !g.pause {
			g.pause = true
		} else {
			g.pause = false
		}
	}
	// Enable/Disable frame limiter
	if rl.IsKeyPressed(rl.KeyF1) {
		if rl.GetFPS() > 60 {
			rl.SetTargetFPS(60)
		} else {
			rl.SetTargetFPS(int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())))
		}
	}
}

func updateCharRec(g *Game) {
	// Update character sprites
	if framesCounter >= 5 {
		currentFrame++
		framesCounter = 0
		if currentFrame > 4 {
			currentFrame = 0
		}
		g.playerRec.X = float32(currentFrame * g.char.Width / 4)
	}
}
