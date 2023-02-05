package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"os/user"
	"time"
)

func GetEnemies(difficulty GameDifficulty) int {
	return int(difficulty) + wave*20
}

func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(128) == 1
}

func LaunchGame() {
	var cDifficulty = Normal
	var enableGore = true
	enableSuicide = false

	// Create window and set target FPS
	rl.InitWindow(screenWidth, screenHeight, "The Liquidator")
	rl.SetTargetFPS(60)

	if enableFullScreen {
		rl.ToggleFullscreen()
	}

	// Init game
	game := NewGame()
	game.heart = rl.LoadTexture("res/heart.png")
	game.char = rl.LoadTexture("res/soldier.png")
	game.dead = rl.LoadTexture("res/dead_soldier.png")
	game.bg = rl.LoadTexture("res/background.png")
	game.menuScreen = rl.LoadTexture("res/menu_screen.png")
	game.shopScreen = rl.LoadTexture("res/shop_screen.png")
	game.enemyTexture = rl.LoadTexture("res/enemy.png")
	game.explosion = rl.LoadTexture("res/explosion.png")
	game.bulletTex = rl.LoadTexture("res/bullet.png")
	game.armalite = rl.LoadTexture("res/armalite.png")
	game.barrett = rl.LoadTexture("res/barrett.png")
	game.galil = rl.LoadTexture("res/galil.png")
	game.groza = rl.LoadTexture("res/groza.png")
	game.machineGun = rl.LoadTexture("res/m60.png")
	game.splatter = rl.LoadTexture("res/splatter.png")
	game.blood = rl.LoadTexture("res/blood.png")

	game.playerRec = rl.Rectangle{Width: float32(game.char.Width / 4), Height: float32(game.char.Height / 2)}
	game.enemyRec = rl.Rectangle{Width: float32(game.enemyTexture.Width / 4), Height: float32(game.enemyTexture.Height)}
	game.splatterRec = rl.Rectangle{Width: float32(game.splatter.Width / 4), Height: float32(game.splatter.Height)}
	game.barbedWire = rl.Rectangle{X: 0, Y: 100, Width: screenWidth, Height: 80}
	game.explosionRec = rl.Rectangle{Width: float32(game.explosion.Width / 16), Height: float32(game.explosion.Height)}

	game.gun[0] = Gun{"AR-15", false, false, true, 6, 30, 30, game.armalite, 0}
	game.gun[1] = Gun{"Galil", true, false, false, 6, 30, 30, game.galil, 2000}
	game.gun[2] = Gun{"Barrett", false, true, false, 6, 20, 20, game.barrett, 5000}
	game.gun[3] = Gun{"Groza", true, true, false, 4, 40, 40, game.groza, 12500}
	game.gun[4] = Gun{"M60", true, false, false, 8, 100, 100, game.machineGun, 25000}

	game.button[0] = Button{size: rl.Vector2{X: screenWidth / 5, Y: screenHeight / 7}, icon: &game.galil, gun: &game.gun[1]}
	game.button[1] = Button{size: rl.Vector2{X: screenWidth / 5, Y: screenHeight / 7}, icon: &game.barrett, gun: &game.gun[2]}
	game.button[2] = Button{size: rl.Vector2{X: screenWidth / 5, Y: screenHeight / 7}, icon: &game.groza, gun: &game.gun[3]}
	game.button[3] = Button{size: rl.Vector2{X: screenWidth / 5, Y: screenHeight / 7}, icon: &game.machineGun, gun: &game.gun[4]}

	for i := 0; i < Guns; i++ {
		game.gun[i].ammo = game.gun[i].maxAmmo
	}

	game.difficulty = cDifficulty

	rl.InitAudioDevice()
	sfxRifle = rl.LoadSound("res/sounds/rifle.wav")
	sfxDeath = rl.LoadSound("res/sounds/loro.mp3")
	sfxGroza = rl.LoadSound("res/sounds/groza.wav")
	sfxSniper = rl.LoadSound("res/sounds/sniper.wav")
	sfxReload = rl.LoadSound("res/sounds/reload.mp3")

	if enableGore {
		game.splatter = game.blood
	}

	killsRequired = GetEnemies(game.difficulty)
	current, _ := user.Current()
	username = current.Username

	for !rl.WindowShouldClose() { // Game loop
		// Reset game when the game is over
		if game.gameOver {
			if rl.IsKeyPressed(rl.KeyEnter) {
				if !displayLeaderboards {
					SubmitNewHiScore(username, score)
					displayLeaderboards = true
				} else {
					Reset(&game)
					displayLeaderboards = false
				}
			}
		}
		specialKeyCallback(&game)
		game.update() // Keep the game running
	}
	game.deInit() // De-initialize everything and close window
}

func NewGame() (g Game) {
	money = 1000
	score = 0
	g.gameOver = false
	g.Init()
	return
}

func Reset(game *Game) {
	kills = 0
	score = 0
	money = 1000
	wave = 1
	killsRequired = GetEnemies(game.difficulty)
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
		game.enemy[i] = Enemy{rl.Vector2{X: float32(rl.GetRandomValue(0, screenWidth-100)), Y: float32(rl.GetRandomValue(screenHeight, screenHeight+1000))}, true, float32(game.difficulty + 1)}
	}
	for i := 0; i < Guns; i++ {
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
		g.enemy[i] = Enemy{rl.Vector2{X: float32(rl.GetRandomValue(0, screenWidth-100)), Y: float32(rl.GetRandomValue(screenHeight, screenHeight+1000))}, true, 3}
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
	rl.UnloadTexture(g.splatter)
	rl.UnloadTexture(g.explosion)
	rl.UnloadTexture(g.bulletTex)
	rl.UnloadTexture(g.armalite)
	rl.UnloadTexture(g.barrett)
	rl.UnloadTexture(g.galil)
	rl.UnloadTexture(g.groza)
	rl.UnloadTexture(g.machineGun)
	rl.UnloadTexture(g.menuScreen)
	rl.UnloadTexture(g.blood)
	rl.UnloadSound(sfxDeath)
	rl.UnloadSound(sfxGroza)
	rl.UnloadSound(sfxRifle)
	rl.UnloadSound(sfxSniper)
	rl.UnloadSound(sfxReload)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
