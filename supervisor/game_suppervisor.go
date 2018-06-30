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
	var nextMoveTotal int
	var nextMove *PlayerDirections
	for play := true; play; play = shared.Continue() {

		time.Sleep(shared.KeyPressedDelay_ms * time.Millisecond)
		//get the players key move
		nextMove = g.Mode.Move()
		nextMoveTotal = nextMove.sum()
		g.DirectionChannel <- nextMove

		//wait extra if a key was pressed
		if nextMoveTotal!= 0 {
			time.Sleep(shared.KeyWaitAfterPressed_ms* time.Millisecond)
		}
	}
}

//get the sum of the keys pressed by player 1 : check if a key was pressed
func (pd * PlayerDirections) sum() int{
	return pd.Player1.X + (pd.Player1.Y*10) //so we never get zero
}
