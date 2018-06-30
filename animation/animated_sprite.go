package animation

import (
	"github.com/gandrin/ASharedJourney/shared"
	"github.com/faiface/pixel"
	"github.com/gandrin/ASharedJourney/mechanics"
	"log"
)

type AnimatedSprite struct {
	pos shared.Position
	strite pixel.Sprite
	currentEvent mechanics.Event
	frame int
}

func (as * AnimatedSprite) Move(){
	//todo move animated sprite
	log.Print("Move called on animated sprite ",as.pos)
}
