package supervisor

import (
	"github.com/ASharedJourney/shared"
	"github.com/faiface/pixel/pixelgl"
)

//transition Direction in x , y coord
type Direction struct {
	X int
	Y int
}


func key() Direction {
	var newDir Direction = Direction{
		X:0,
		Y:0,
	}
	if shared.Win.Pressed(pixelgl.KeyLeft) {
		newDir.X = -1
	}
	if shared.Win.Pressed(pixelgl.KeyRight) {
		newDir.X = 1
	}
	if shared.Win.Pressed(pixelgl.KeyDown) {
		newDir.Y = 1
	}
	if shared.Win.Pressed(pixelgl.KeyUp) {
		newDir.Y = -1
	}
}
