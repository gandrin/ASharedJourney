package supervisor

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/gandrin/ASharedJourney/shared"
)

type Event string

//get the key values that was pressed
func catchEvent() *Event {
	var event = new(Event)
	//check if key was just pressed
	if shared.Win.Pressed(pixelgl.KeyR) {
		*event = "RESTART"
		return event
	}
	return event
}
