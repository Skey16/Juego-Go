// game/collision_detection.go
package game

import (
	"time"
)

func collisionDetection() {
	for {
		if !GameOver {
			BulletsMutex.Lock()
			for _, bullet := range Bullets {
				scaledPlayerSize := PlayerSize * PlayerScaling
				if float64(PlayerX) < bullet.X+float64(scaledPlayerSize) &&
					float64(PlayerX)+float64(scaledPlayerSize) > bullet.X &&
					float64(PlayerY) < bullet.Y+float64(scaledPlayerSize) &&
					float64(PlayerY)+float64(scaledPlayerSize) > bullet.Y {

					GameOverMutex.Lock()
					GameOver = true
					GameOverMutex.Unlock()
				}
			}
			BulletsMutex.Unlock()
		}
		time.Sleep(10 * time.Millisecond)
	}
}
