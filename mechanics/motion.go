package mechanics

import (
	"log"

	"github.com/gandrin/ASharedJourney/shared"
	"github.com/gandrin/ASharedJourney/supervisor"
)

//structure that will be passes on to animator
type Motion struct {
	Player1 playerMechanics
	Player2 playerMechanics
	//potencially add other events later not dependant directly on the player position
	//OtherEvents []Event
}

//move function recives as input the data from a player direction channel
func (gm *Mechanics) Move(playDir *supervisor.PlayerDirections) *Motion {

	var nextPos1 shared.Position //next position for player 1 with current direction
	var nextPos2 shared.Position // same for player 2

	var newMotion *Motion = new(Motion)
	//set old positions
	newMotion.Player1.OldPos = gm.Player1.pos
	newMotion.Player2.OldPos = gm.Player2.pos
	//set player type
	newMotion.Player1.PType = gm.Player1.pType
	newMotion.Player2.PType = gm.Player2.pType

	log.Print("Player old positions , p1 ", newMotion.Player1.OldPos, " p2 ",
		newMotion.Player2.OldPos)

	//Player 1
	//check next position based on motions
	nextPos1 = playDir.Player1.Next(newMotion.Player1.OldPos)
	gm.move_player(gm.Player1.pType, nextPos1, &newMotion.Player1)
	//check if player has triggered an event
	newMotion.Player1.PlayerEvent = gm.check_player_event(gm.Player1.pType, nextPos1, &newMotion.Player1)

	//Player 2
	nextPos2 = playDir.Player2.Next(newMotion.Player2.OldPos)
	gm.move_player(gm.Player2.pType, nextPos2, &newMotion.Player2)
	//check if player has triggered an event
	newMotion.Player2.PlayerEvent = gm.check_player_event(gm.Player2.pType, nextPos2, &newMotion.Player2)

	//log debug
	log.Printf("Motion player 1 ", newMotion.Player1)
	log.Printf("Motion player 2 ", newMotion.Player2)
	return newMotion
}

//move player if hitmap permits
func (gm *Mechanics) move_player(ptype PlayerType, nextPos shared.Position, playerMotion *playerMechanics) {
	//check if can move
	if ptype.can_walk(gm.hitMap[nextPos.X][nextPos.Y]) {
		//can move according to hit map
		playerMotion.NewPos = nextPos
	}
}

//check if player has triggered an event
func (gm *Mechanics) check_player_event(ptype PlayerType, nextPos shared.Position, playerMotion *playerMechanics) *Event {
	var nEvent *Event
	var eventType *eventType
	//check if we have triggered an event
	eventType = gm.eventMap[nextPos.X][nextPos.Y]
	if eventType != nil {
		//potencially have an event
		//check if our player can trigger it
		nEvent = ptype.trigger_event(eventType)
	}
	return nEvent
}
