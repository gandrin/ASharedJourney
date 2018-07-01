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
	Mode             GameMode
}

var Sup *GameSupervisor

//Start inits the game and specify the game mode
func Start(gm GameMode) chan *PlayerDirections {
	Sup = new(GameSupervisor)
	Sup.Mode = gm
	Sup.DirectionChannel = make(chan *PlayerDirections, 1)
	return Sup.DirectionChannel
}

//Play launches game supervisor (should be lauched last)

func (g *GameSupervisor) Play() {
	var nextMove *PlayerDirections
	for play := true; play; play = shared.Continue() {

		time.Sleep(shared.KeyPressedDelay_ms * time.Millisecond)

		//get the players key move
		nextMove = g.Mode.Move()
		if( nextMove.Player1.X != 0 || nextMove.Player1.Y!= 0 ){
			//new move
			shared.AddAction()
		}

		g.DirectionChannel <- nextMove

	}
}
