package mechanics

import (
	"github.com/faiface/pixel"
	"github.com/gandrin/ASharedJourney/shared"
	"github.com/gandrin/ASharedJourney/supervisor"
	"github.com/gandrin/ASharedJourney/tiles"
)

//move function recives as input the data from a player direction channel
func (gm *Mechanics) Move(playDir *supervisor.PlayerDirections) *tiles.World {
	//log.Printf("Move called")

	var nextPos1 pixel.Vec //next position for player 1 with current direction
	var auxPos pixel.Vec

	var canPlayer1Move bool = true
	//Player 1
	//check next position based on motions

	nextPos1 = playDir.Player1.Next(gm.world.Players[0].Position)

	// Obstacles
	for _, val := range gm.world.Obstacles {
		if val.Position.X == nextPos1.X && val.Position.Y == nextPos1.Y {
			canPlayer1Move = false
		}
	}

	// Movables
	if canPlayer1Move {
		for n, val := range gm.world.Movables {
			if val.Position.X == nextPos1.X && val.Position.Y == nextPos1.Y {
				// There's a mouvable in that position

				auxPos = playDir.Player1.Next(nextPos1)
				for _, val := range gm.world.Obstacles {
					if val.Position.X == auxPos.X && val.Position.Y == auxPos.Y {
						canPlayer1Move = false
					}
				}
				for _, val := range gm.world.Movables {
					if val.Position.X == auxPos.X && val.Position.Y == auxPos.Y {
						canPlayer1Move = false
					}
				}
				//
				if canPlayer1Move {
					gm.world.Movables[n].Position = auxPos
				}
			}
		}
	}

	if canPlayer1Move {
		gm.world.Players[0].Position = nextPos1
	}

	return gm.copyToNewWorld()
}

//move player if hitmap permits
func (gm *Mechanics) move_player(ptype PlayerType, nextPos shared.Position) bool {
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

func (gm *Mechanics) copyToNewWorld() *tiles.World {
	var newWorld *tiles.World = new(tiles.World)
	//copy player locations

	//copy world
	//make a copy of world : todo check if doesn't fail
	//this will have to be updated
	newWorld.BackgroundTiles = make([]tiles.SpriteWithPosition, len(gm.world.BackgroundTiles))
	newWorld.Movables = make([]tiles.SpriteWithPosition, len(gm.world.Movables))
	newWorld.Players = make([]tiles.SpriteWithPosition, len(gm.world.Players))
	copy(newWorld.BackgroundTiles, gm.world.BackgroundTiles)
	copy(newWorld.Movables, gm.world.Movables)
	copy(newWorld.Players, gm.world.Players)
	copy(newWorld.Water, gm.world.Water)

	return newWorld
}
