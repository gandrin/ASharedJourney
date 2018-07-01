package supervisor

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/gandrin/ASharedJourney/shared"
)

//transition Direction in x , y coord
type Direction struct {
	X int
	Y int
}

//get the key values that was pressed
//old directions
var prevKeyPressed pixelgl.Button
var interty int

func key() Direction {
	var pressed pixelgl.Button

	var newDir = Direction{
		X: 0,
		Y: 0,
	}

	//check if key was just pressed
	if shared.Win.Pressed(pixelgl.KeyLeft) {
		pressed = pixelgl.KeyLeft
		newDir.X = -1
		goto end
	} else if shared.Win.Pressed(pixelgl.KeyRight) {
		pressed = pixelgl.KeyRight
		newDir.X = 1
		goto end
	} else if shared.Win.Pressed(pixelgl.KeyDown) {
		newDir.X = 0
		newDir.Y = -1
		pressed = pixelgl.KeyDown
		goto end
	} else if shared.Win.Pressed(pixelgl.KeyUp) {
		newDir.X = 0
		newDir.Y = 1
		pressed = pixelgl.KeyUp
		goto end
	} else {
		//no key pressed
		prevKeyPressed = pixelgl.Key0 //default
		return newDir
	}

end:
	//check if key repressed
	if pressed == prevKeyPressed {
		//time penalty
		if interty == 5 {
			interty = 0
			return newDir
		} else {
			newDir = Direction{0, 0}
		}
		interty += 1
	} else {
		interty = 0
	}
	prevKeyPressed = pressed
	return newDir
}
