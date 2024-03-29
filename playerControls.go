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
	var playerIsMoving = false
	// Movement:
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		playerIsMoving = true
		if g.player.position.X+80 < screenWidth {
			g.player.position.X += g.player.speed
		}
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		playerIsMoving = true
		if g.player.position.X > 0 {
			g.player.position.X -= g.player.speed
		}
	}
	if !playerIsMoving {
		g.playerRec.X = 0
	} else {
		updateCharRec(g)
	}

	// Semi-auto:
	if kills < killsRequired {

		if rl.IsKeyPressed(rl.KeySpace) || gamerMode && rl.IsMouseButtonPressed(0) {
			if g.gun[cWeapon].ammo > 0 && !g.player.reloading && !g.gun[cWeapon].automatic {
				initNewBullet(g)
				if !g.gun[cWeapon].armourPiercing {
					rl.PlaySoundMulti(sfxRifle)
				} else {
					rl.PlaySoundMulti(sfxSniper)
				}
			}
		}

		// Full-auto:
		if rl.IsKeyDown(rl.KeySpace) || gamerMode && rl.IsMouseButtonDown(0) {
			if g.gun[cWeapon].ammo > 0 && !g.player.reloading && g.gun[cWeapon].automatic {
				firingRateCounter++
				if firingRateCounter%g.gun[cWeapon].firingRate == 0 {
					initNewBullet(g)
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
	if enableSuicide {
		if rl.IsKeyPressed(rl.KeyEnd) {
			g.player.lives = 0
		}
	}
	if gamerMode {
		if rl.IsKeyPressed(rl.KeyE) && cWeapon < 4 {
			switchWeapon(g, cWeapon+1)
		}
		if rl.IsKeyPressed(rl.KeyQ) && cWeapon > 0 {
			switchWeapon(g, cWeapon-1)
		}
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
	// Change the controller layout to WASD and mouse.
	if rl.IsKeyPressed(rl.KeyF2) {
		if !gamerMode {
			gamerMode = true
		} else {
			gamerMode = false
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
