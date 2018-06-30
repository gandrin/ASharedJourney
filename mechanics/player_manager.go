package mechanics

import "github.com/gandrin/ASharedJourney/shared"

//add data about a player used by the game mechanics
type playerManager struct {
	pos   shared.Position
	pType PlayerType
}

type playerMechanics struct {
	PType       PlayerType
	OldPos      shared.Position
	NewPos      shared.Position
	PlayerEvent *Event
}
