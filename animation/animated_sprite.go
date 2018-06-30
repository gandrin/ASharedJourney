package animation

import (
<<<<<<< HEAD
	"github.com/gandrin/ASharedJourney/shared"
	"github.com/faiface/pixel"
	"github.com/gandrin/ASharedJourney/mechanics"
=======
>>>>>>> 22a5795a8c766bfe8864c04217a9f695b4a6d0c1
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
