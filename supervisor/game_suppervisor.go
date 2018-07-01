//retrieve player input and send the directions of motions to the mechanics
package supervisor

import (
	"time"

	"github.com/gandrin/ASharedJourney/shared"
)

type PlayerDirections struct {
	Player1 Direction
	Player2 Direction
}
type GameSupervisor struct {
	DirectionChannel chan *PlayerDirections
}

var Sup *GameSupervisor

//Start inits the game and specify the game mode
func Start() chan *PlayerDirections {
	Sup = new(GameSupervisor)
	Sup.DirectionChannel = make(chan *PlayerDirections, 1)
	return Sup.DirectionChannel
}

//Play launches game supervisor (should be lauched last)

func (gameSupervisor *GameSupervisor) Play() {
	var nextMove *PlayerDirections
	for play := true; play; play = shared.Continue() {

		time.Sleep(shared.KeyPressedDelay_ms * time.Millisecond)

		//get the players key move
		nextMove = Move()

		gameSupervisor.DirectionChannel <- nextMove

	}
}
