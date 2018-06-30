package supervisor

import (
	"log"

	"github.com/gandrin/ASharedJourney/shared"
)

//call motion
func (mode GameMode) Move() *PlayerDirections {
	var newDir *PlayerDirections = new(PlayerDirections)
	if mode == OnePlayer {
		newDir.Player1 = key()
		newDir.mirror()
	} else {
		log.Fatal("Unknown mode")
	}
	return newDir
}

//mirror motion of player 1 onto direction of player 2
func (dir *PlayerDirections) mirror() {
	if dir.Player1.X != 0 {
		dir.Player2.X = dir.Player1.X * (-1)
	}
	if dir.Player1.Y != 0 {
		dir.Player2.Y = dir.Player1.Y * (-1)
	}
}

//calculate next position based on direction
func (dir Direction) Next(currentPos shared.Position) shared.Position {
	currentPos.X += dir.X
	currentPos.Y += dir.Y
	//log.Printf("Calculated next position ", currentPos)
	return currentPos
}
