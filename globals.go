package main

import rl "github.com/gen2brain/raylib-go/raylib"

const screenWidth = 1280
const screenHeight = 720
const PlayerMaxLife = 5
const MaxBullets = 30
const MaxEnemies = 8
const Guns = 4

var cWeapon = 0 // 0 = AR-15, 1 = Galil, 2 = Barrett, 3 = Groza, 4 = M60
var currentFrame int32 = 0
var enemyFrame int32 = 0
var kills = 0
var killsRequired int
var wave = 1
var enableFullScreen bool
var inShop = false
var displayLeaderboards = false
var firingRateCounter = 0
var framesCounter = 0
var reloadCounter = 0
var score = 0
var money = 0
var sfxDeath rl.Sound
var sfxGroza rl.Sound
var sfxRifle rl.Sound
var sfxSniper rl.Sound
var sfxReload rl.Sound
var username string

type Player struct {
	position  rl.Vector2
	lives     int
	speed     float32
	reloading bool
}

type Enemy struct {
	position rl.Vector2
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

type Game struct {
	gameOver     bool
	pause        bool
	player       Player
	enemy        [MaxEnemies]Enemy
	bullet       [MaxBullets]Bullet
	gun          [5]Gun
	button       [4]Button
	char         rl.Texture2D
	dead         rl.Texture2D
	heart        rl.Texture2D
	bg           rl.Texture2D
	shopScreen   rl.Texture2D
	enemyTexture rl.Texture2D
	splatter     rl.Texture2D
	explosion    rl.Texture2D
	bulletTex    rl.Texture2D
	armalite     rl.Texture2D
	barrett      rl.Texture2D
	galil        rl.Texture2D
	groza        rl.Texture2D
	machineGun   rl.Texture2D
	playerRec    rl.Rectangle
	enemyRec     rl.Rectangle
	splatterRec  rl.Rectangle
	barbedWire   rl.Rectangle
	explosionRec rl.Rectangle
}
