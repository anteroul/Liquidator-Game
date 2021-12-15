// (C) Uljas Antero Lindell 2021
// Version 0.7 Alpha

package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

const screenWidth = 1280
const screenHeight = 720
const PlayerMaxLife = 3
const MaxBullets = 30
const MaxEnemies = 8

var cWeapon = 0 // 0 = AR-15, 1 = Galil, 2 = Barrett, 3 = Groza, 4 = M60
var currentFrame int32 = 0
var enemyFrame int32 = 0
var kills = 0
var killsRequired int
var wave = 1
var enableFullScreen bool
var inShop = false
var firingRateCounter = 0
var framesCounter = 0
var reloadCounter = 0
var score = 0
var money = 0
var sfxDeath rl.Sound
var sfxGroza rl.Sound
var sfxRifle rl.Sound
var sfxSniper rl.Sound

type Player struct {
	position  rl.Vector2
	lives     int
	speed     float32
	reloading bool
}

type Enemy struct {
	position rl.Vector2
	armed    bool
	active   bool
	speed    float32
}

type Bullet struct {
	rec    rl.Rectangle
	speed  rl.Vector2
	active bool
	Color  rl.Color
}

type Gun struct {
	name           string
	automatic      bool
	armourPiercing bool
	inInventory    bool
	firingRate     int // Low number = fast rate of fire
	ammo           int
	maxAmmo        int
	gunIcon        rl.Texture2D
	price          int
}

type Button struct {
	size rl.Vector2
	icon *rl.Texture2D
	gun  *Gun
}

type Game struct {
	gameOver          bool
	pause             bool
	player            Player
	enemy             [MaxEnemies]Enemy
	bullet            [MaxBullets]Bullet
	gun               [5]Gun
	button            [4]Button
	char              rl.Texture2D
	dead              rl.Texture2D
	heart             rl.Texture2D
	bg                rl.Texture2D
	shopScreen        rl.Texture2D
	deathScreen       rl.Texture2D
	enemyTexture      rl.Texture2D
	armedEnemyTexture rl.Texture2D
	splatter          rl.Texture2D
	bulletTex         rl.Texture2D
	armalite          rl.Texture2D
	barrett           rl.Texture2D
	galil             rl.Texture2D
	groza             rl.Texture2D
	machineGun        rl.Texture2D
	playerRec         rl.Rectangle
	enemyRec          rl.Rectangle
	splatterRec       rl.Rectangle
	barbedWire        rl.Rectangle
}

func GetEnemies() int {
	var enemies = 20 * wave
	return enemies
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

func main() {
	Setup() // Game Settings
	// Create window and set target FPS
	rl.InitWindow(screenWidth, screenHeight, "The Liquidator")
	rl.SetTargetFPS(int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())))

	if enableFullScreen {
		rl.ToggleFullscreen()
	}

	// Init game
	game := NewGame()
	game.heart = rl.LoadTexture("res/heart.png")
	game.char = rl.LoadTexture("res/soldier.png")
	game.dead = rl.LoadTexture("res/dead_soldier.png")
	game.bg = rl.LoadTexture("res/background.png")
	game.shopScreen = rl.LoadTexture("res/shop_screen.png")
	game.enemyTexture = rl.LoadTexture("res/enemy.png")
	game.armedEnemyTexture = rl.LoadTexture("res/enemy_soldier.png")
	game.splatter = rl.LoadTexture("res/splatter.png")
	game.deathScreen = rl.LoadTexture("res/red_screen.png")
	game.bulletTex = rl.LoadTexture("res/bullet.png")
	game.armalite = rl.LoadTexture("res/armalite.png")
	game.barrett = rl.LoadTexture("res/barrett.png")
	game.galil = rl.LoadTexture("res/galil.png")
	game.groza = rl.LoadTexture("res/groza.png")
	game.machineGun = rl.LoadTexture("res/m60.png")
	game.playerRec = rl.Rectangle{Width: float32(game.char.Width / 4), Height: float32(game.char.Height)}
	game.enemyRec = rl.Rectangle{Width: float32(game.enemyTexture.Width / 4), Height: float32(game.enemyTexture.Height)}
	game.splatterRec = rl.Rectangle{Width: float32(game.splatter.Width / 4), Height: float32(game.splatter.Height)}
	game.barbedWire = rl.Rectangle{X: 0, Y: 100, Width: screenWidth, Height: 80}

	game.gun[0] = Gun{"AR-15", false, false, true, 6, 30, 30, game.armalite, 0}
	game.gun[1] = Gun{"Galil", true, false, false, 6, 30, 30, game.galil, 3000}
	game.gun[2] = Gun{"Barrett", false, true, false, 6, 20, 20, game.barrett, 5000}
	game.gun[3] = Gun{"Groza", true, true, false, 4, 20, 20, game.groza, 12500}
	game.gun[4] = Gun{"M60", true, false, false, 8, 100, 100, game.machineGun, 25000}

	game.button[0] = Button{size: rl.Vector2{X: screenWidth / 5, Y: screenHeight / 7}, icon: &game.galil, gun: &game.gun[1]}
	game.button[1] = Button{size: rl.Vector2{X: screenWidth / 5, Y: screenHeight / 7}, icon: &game.barrett, gun: &game.gun[2]}
	game.button[2] = Button{size: rl.Vector2{X: screenWidth / 5, Y: screenHeight / 7}, icon: &game.groza, gun: &game.gun[3]}
	game.button[3] = Button{size: rl.Vector2{X: screenWidth / 5, Y: screenHeight / 7}, icon: &game.machineGun, gun: &game.gun[4]}

	for i := 0; i < 5; i++ {
		game.gun[i].ammo = game.gun[i].maxAmmo
	}

	rl.InitAudioDevice()
	sfxRifle = rl.LoadSound("res/sounds/rifle.wav")
	sfxDeath = rl.LoadSound("res/sounds/loro.mp3")
	sfxGroza = rl.LoadSound("res/sounds/groza.wav")
	sfxSniper = rl.LoadSound("res/sounds/sniper.wav")

	killsRequired = GetEnemies()

	for !rl.WindowShouldClose() { // Game loop
		// Reset game when the game is over
		if game.gameOver {
			if rl.IsKeyPressed(rl.KeyEnter) {
				Reset(&game)
			}
		}

		game.update() // Keep the game running

	}
	game.deInit() // De-initialize everything and close window
}

func NewGame() (g Game) {
	money = 0
	score = 0
	g.gameOver = false
	g.Init()
	return
}

func Reset(game *Game) {
	kills = 0
	score = 0
	money = 0
	wave = 1
	killsRequired = GetEnemies()
	game.gameOver = false
	// Initialize player
	game.player.position = rl.NewVector2(float32(screenWidth)/2, 40)
	game.player.lives = PlayerMaxLife
	game.player.speed = 4.5
	// Initialize bullets
	for i := 0; i < MaxBullets; i++ {
		game.bullet[i] = Bullet{rec: rl.Rectangle{X: game.player.position.X + 40, Y: game.player.position.Y + 105, Width: 5, Height: 10}, speed: rl.Vector2{Y: 15}, active: false, Color: rl.Yellow}
	}
	// Initialize enemies
	for i := 0; i < MaxEnemies; i++ {
		game.enemy[i] = Enemy{rl.Vector2{X: float32(rl.GetRandomValue(0, screenWidth-100)), Y: float32(rl.GetRandomValue(screenHeight, screenHeight+1000))}, false, true, 3}
	}
	for i := 0; i < 5; i++ {
		game.gun[i].ammo = game.gun[i].maxAmmo
		game.gun[i].inInventory = false
	}
	game.gun[0].inInventory = true
	cWeapon = 0
}

func (g *Game) Init() {
	// Initialize player
	g.player.position = rl.NewVector2(float32(screenWidth)/2, 40)
	g.player.lives = PlayerMaxLife
	g.player.speed = 4.5
	// Initialize bullets
	for i := 0; i < MaxBullets; i++ {
		g.bullet[i] = Bullet{rec: rl.Rectangle{X: g.player.position.X + 40, Y: g.player.position.Y + 105, Width: 5, Height: 10}, speed: rl.Vector2{Y: 15}, active: false, Color: rl.Yellow}
	}
	// Initialize enemies
	for i := 0; i < MaxEnemies; i++ {
		g.enemy[i] = Enemy{rl.Vector2{X: float32(rl.GetRandomValue(0, screenWidth-100)), Y: float32(rl.GetRandomValue(screenHeight, screenHeight+1000))}, false, true, 3}
	}
}

func (g *Game) deInit() {
	// De-initialize everything and close window
	rl.UnloadTexture(g.bg)
	rl.UnloadTexture(g.shopScreen)
	rl.UnloadTexture(g.heart)
	rl.UnloadTexture(g.char)
	rl.UnloadTexture(g.dead)
	rl.UnloadTexture(g.enemyTexture)
	rl.UnloadTexture(g.armedEnemyTexture)
	rl.UnloadTexture(g.splatter)
	rl.UnloadTexture(g.deathScreen)
	rl.UnloadTexture(g.bulletTex)
	rl.UnloadTexture(g.armalite)
	rl.UnloadTexture(g.barrett)
	rl.UnloadTexture(g.galil)
	rl.UnloadTexture(g.groza)
	rl.UnloadTexture(g.machineGun)
	rl.UnloadSound(sfxDeath)
	rl.UnloadSound(sfxGroza)
	rl.UnloadSound(sfxRifle)
	rl.UnloadSound(sfxSniper)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func displayShopScreen(g *Game) {
	if rl.IsKeyPressed(rl.KeyEnter) {
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
}

func (g *Game) update() {

	if !inShop {

		framesCounter++

		if kills == killsRequired && g.player.lives > 0 {
			for i := 0; i < 5; i++ {
				g.gun[i].ammo = g.gun[i].maxAmmo
			}
			g.player.reloading = false
			kills = 0
			wave++
			killsRequired = GetEnemies()
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
		// Shoot logic
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
							tangoDown()
						}
						if g.enemy[j].position.Y < 100 {
							if rl.CheckCollisionPointRec(rl.Vector2{X: g.bullet[i].rec.X, Y: g.bullet[i].rec.Y - 75}, rl.Rectangle{X: g.enemy[j].position.X, Y: g.enemy[j].position.Y, Width: 90, Height: 40}) {
								g.enemy[j].active = false
								if !g.gun[cWeapon].armourPiercing {
									g.bullet[i].active = false
								}
								tangoDown()
							}
						}
					}
				}
				if g.bullet[i].rec.Y+g.bullet[i].rec.Height >= screenHeight {
					g.bullet[i].active = false
				}
			}
		}
		// Enemy behaviour
		// TODO: Shooting for armed enemies (armed enemies are currently disabled)
		for i := 0; i < MaxEnemies; i++ {
			if g.enemy[i].active {
				g.enemy[i].position.Y -= float32(rl.GetRandomValue(1, int32(g.enemy[i].speed)))
				// Crossing the border
				if g.enemy[i].position.Y < -120 {
					g.enemy[i].position.X = float32(rl.GetRandomValue(0, screenWidth-100))
					g.enemy[i].position.Y = float32(rl.GetRandomValue(screenHeight+200, screenHeight+1000))
					if !g.gameOver {
						score -= 500
					}
				}
				// Collision with player
				if !g.gameOver {
					if rl.CheckCollisionRecs(rl.Rectangle{X: g.enemy[i].position.X, Y: g.enemy[i].position.Y, Width: 90, Height: 40}, rl.Rectangle{X: g.player.position.X, Y: g.player.position.Y, Width: 90, Height: 40}) {
						g.enemy[i].active = false
						g.player.lives--
						tangoDown()
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
					if killsRequired-kills >= MaxEnemies {
						g.enemy[i].active = true
					}
				}
			}
			go updateEnemyRec(g, g.enemy[i])
		}
		// Update controls if player is alive
		if !g.gameOver {
			go keyCallback(g) // Game controls
		}
	} else {
		displayShopScreen(g)
	}

	draw(g) // Display graphics

}

func tangoDown() {
	rl.PlaySoundMulti(sfxDeath)
	score += 100
	if score >= 0 {
		money += 50
	}
	kills++
	enemyFrame = 0
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
					rl.PlaySoundMulti(sfxGroza)
				}
			}
		}
	}
	// Reload keys:
	if rl.IsKeyPressed(rl.KeyR) {
		if g.gun[cWeapon].ammo != g.gun[cWeapon].maxAmmo {
			g.player.reloading = true
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
}

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
			if onClickEvent(&rl.Rectangle{X: float32(int32(screenWidth/4*i + 40)), Y: screenHeight / 3 * 2, Width: float32(int32(g.button[i].size.X)), Height: float32(int32(g.button[i].size.Y + 15))}) {
				if money >= g.gun[i+1].price && !g.gun[i+1].inInventory {
					money -= g.gun[i+1].price
					g.gun[i+1].inInventory = true
				}
			}
		}
		rl.DrawText("Press enter to exit shop", screenWidth/3.5, screenHeight*0.9, 40, rl.SkyBlue)
	}

	rl.DrawText(strconv.Itoa(money)+"$", 20, screenHeight-60, 40, rl.Green)

	rl.DrawFPS(5, 0)
	rl.EndDrawing()
}

// An event listener
func onClickEvent(rectangle *rl.Rectangle) bool {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), *rectangle) && rl.IsMouseButtonPressed(0) {
		return true
	} else {
		return false
	}
}
