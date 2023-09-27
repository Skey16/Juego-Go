// game/generate_bullets.go
package game

import (
	"github.com/hajimehoshi/ebiten/v2" // Asegúrate de que esta línea esté presente
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "log"
    "math/rand"
    "time"
)

func generateBullets() {
	for {
		bulletImage, _, err := ebitenutil.NewImageFromFile("assets/bullet.png")
		if err != nil {
			log.Fatal(err)
		}

		numBullets := 4
		for i := 0; i < numBullets; i++ {
			bullet := Bullet{
				X:     0,
				Y:     float64(rand.Intn(ScreenHeight)),
				Image: ebiten.NewImageFromImage(bulletImage),
			}

			BulletsMutex.Lock()
			Bullets = append(Bullets, bullet)
			BulletsMutex.Unlock()
		}

		time.Sleep(1 * time.Second)
	}
}
