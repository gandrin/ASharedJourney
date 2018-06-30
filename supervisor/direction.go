package supervisor

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/gandrin/ASharedJourney/tiles"
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
	dir.Player2.X = dir.Player1.X
	if dir.Player1.Y != 0 {
		dir.Player2.Y = dir.Player1.Y * (-1)
	}
}

//calculate next position based on direction
func (dir Direction) Next(currentPos pixel.Vec) pixel.Vec {
	currentPos.X += float64(dir.X * tiles.TileSize)
	currentPos.Y += float64(dir.Y * tiles.TileSize)
	//log.Printf("Calculated next position ", currentPos)
	return currentPos
}
