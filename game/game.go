// game/game.go
package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"sync"
	
)

const (
	ScreenWidth   = 800
	ScreenHeight  = 600
	PlayerSize    = 30
	PlayerSpeed   = 6
	BulletSpeed   = 9
	PlayerScaling = 1.0
)

var (
	PlayerX, PlayerY float64 = (ScreenWidth - PlayerSize) / 2, (ScreenHeight - PlayerSize) / 2
	PlayerImage      *ebiten.Image

	Bullets       []Bullet
	BulletsMutex  sync.Mutex

	GameOver      bool
	GameOverMutex sync.Mutex

	Score     int
	ScoreChan = make(chan int)
)

type Bullet struct {
	X, Y  float64
	Image *ebiten.Image
}

type Game struct{}

func NewGame() *Game {
	initGame()
	return &Game{}
}

func initGame() {
	img, _, err := ebitenutil.NewImageFromFile("assets/Mario.png")
	if err != nil {
		log.Fatal(err)
	}
	PlayerImage = ebiten.NewImageFromImage(img)

	go generateBullets()
	go collisionDetection()
	go updateScore()
}

func (g *Game) Update() error {
	GameOverMutex.Lock()
	defer GameOverMutex.Unlock()

	if GameOver {
		if ebiten.IsKeyPressed(ebiten.KeyEscape) {
			GameOver = false
			Bullets = nil
			PlayerX, PlayerY = (ScreenWidth-PlayerSize)/2, (ScreenHeight-PlayerSize)/2
			Score = 0
		}
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) && PlayerY > 0 {
		PlayerY -= PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && PlayerY+PlayerSize < ScreenHeight {
		PlayerY += PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && PlayerX > 0 {
		PlayerX -= PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && PlayerX+PlayerSize < ScreenWidth {
		PlayerX += PlayerSpeed
	}

	BulletsMutex.Lock()
	newBullets := []Bullet{}
	for _, bullet := range Bullets {
		bullet.X += BulletSpeed
		if bullet.X < ScreenWidth {
			newBullets = append(newBullets, bullet)
		} else {
			ScoreChan <- 1
		}
	}
	Bullets = newBullets
	BulletsMutex.Unlock()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	scaleFactor := 4.0 // Escalar al 50% de su tamaño original
	op.GeoM.Scale(scaleFactor, scaleFactor)
	op.GeoM.Translate(PlayerX, PlayerY)

	screen.DrawImage(PlayerImage, op)

	BulletsMutex.Lock()
	for _, bullet := range Bullets {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(bullet.X, bullet.Y)
		screen.DrawImage(bullet.Image, op)
	}
	BulletsMutex.Unlock()

	if GameOver {
		ebitenutil.DebugPrint(screen, "Game Over - Press Esc to restart")
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", Score))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return ScreenWidth, ScreenHeight
}
