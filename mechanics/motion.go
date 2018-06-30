package mechanics

import (
	"github.com/gandrin/ASharedJourney/shared"
	"github.com/gandrin/ASharedJourney/supervisor"
	"github.com/gandrin/ASharedJourney/tiles"
	"log"
)



//move function recives as input the data from a player direction channel
func (gm *Mechanics) Move(playDir *supervisor.PlayerDirections) *tiles.World {
	//log.Printf("Move called")

	var nextPos1 shared.Position //next position for player 1 with current direction
	var nextPos2 shared.Position // same for player 2

	//Player 1
	//check next position based on motions
	nextPos1 = playDir.Player1.Next(gm.Player1.Pos)
	//check if player can go on tile
	if gm.move_player(gm.Player1.PType, nextPos1){
		//check for movables here
		gm.Player1.Pos = nextPos1
	}
	//check if player has triggered an event
	//todo newMotion.Player1.PlayerEvent = gm.check_player_event(gm.Player1.PType, nextPos1, &newMotion.Player1)

	//Player 2
	nextPos2 = playDir.Player2.Next(gm.Player2.Pos)
	//check if player can go on tile
	if gm.move_player(gm.Player2.PType, nextPos1){
		//check for movables here
		gm.Player2.Pos = nextPos2
	}
	//todo check if player has triggered an event
	//newMotion.Player2.PlayerEvent = gm.check_player_event(gm.Player2.PType, nextPos2, &newMotion.Player2)



	//log debug
	log.Print("Motion player 1 ", gm.Player1.Pos,"Motion player 2 ", gm.Player2.Pos)

	//update map
	gm.world.Players[0].Position.X = float64(gm.Player1.Pos.X) * 16
	gm.world.Players[0].Position.Y = float64(gm.Player1.Pos.Y) * 16
	if(len(gm.world.Players)>1){
		gm.world.Players[1].Position.X = float64(gm.Player2.Pos.X) * 16
		gm.world.Players[1].Position.X = float64(gm.Player2.Pos.Y) * 16
	}


	return gm.copyToNewWorld()
}

//move player if hitmap permits
func (gm *Mechanics) move_player(ptype PlayerType, nextPos shared.Position) bool{
	//log.Print("Moving player ", nextPos , " legth of hitmap ",len(gm.hitMap),":",len(gm.hitMap[0]))
	//check if can move
	var hitVal = gm.hitMap[nextPos.X][nextPos.Y]
	//log.Printf("hit values ", hitVal)
	if ptype.can_walk(hitVal) {
		//can move according to hit map
		return true
	}
	return true
}

//check if player has triggered an event
func (gm *Mechanics) check_player_event(ptype PlayerType, nextPos shared.Position, playerMotion *playerMechanics) *Event {
	var nEvent *Event
	var eventType *EventType
	//check if we have triggered an event
	eventType = gm.eventMap[nextPos.X][nextPos.Y]
	if eventType != nil {
		//potencially have an event
		//check if our player can trigger it
		nEvent = ptype.trigger_event(eventType) // + dir +
	}
	return nEvent
}

func (gm *Mechanics) copyToNewWorld() * tiles.World{
	var newWorld *tiles.World =new(tiles.World)
	//copy player locations

	//copy world
	//make a copy of world : todo check if doesn't fail
	//this will have to be updated
	newWorld.BackgroundTiles = make([]tiles.SpriteWithPosition, len(gm.world.BackgroundTiles))
	newWorld.Movables = make([]tiles.SpriteWithPosition, len(gm.world.Movables))
	newWorld.Players = make([]tiles.SpriteWithPosition, len(gm.world.Players))
	copy(newWorld.BackgroundTiles, gm.world.BackgroundTiles)
	copy(newWorld.Movables ,gm.world.Movables)
	copy(newWorld.Players, gm.world.Players)

	return newWorld
}