
// game/update_score.go
package game

func updateScore() {
	for {
		select {
		case <-ScoreChan:
			Score++
		}
	}
}
