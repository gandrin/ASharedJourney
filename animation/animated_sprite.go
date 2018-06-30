package animation

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/gandrin/ASharedJourney/mechanics"
	"github.com/gandrin/ASharedJourney/shared"
)

type AnimatedSprite struct {
	pos          shared.Position
	strite       pixel.Sprite
	currentEvent mechanics.Event
	frame        int
}

func (as *AnimatedSprite) Move() {
	//todo move animated sprite
	log.Print("Move called on animated sprite ", as.pos)
}
