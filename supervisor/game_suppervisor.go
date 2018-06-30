//retrieve player input and send the directions of motions to the mechanics
package supervisor

import (
	"github.com/ASharedJourney/shared"
	"time"
)

type PlayerDirections struct {
	Player1 Direction
	Player2 Direction
}
type GameSupervisor struct {
	DirectionChannel chan *PlayerDirections
	Mode GameMode

}
var Sup * GameSupervisor

//Init the game and specify the game mode
func Start(gm GameMode) chan *PlayerDirections{
	Sup = new(GameSupervisor)
	Sup.Mode = gm
	Sup.DirectionChannel = make(chan *PlayerDirections, 1)
	return Sup.DirectionChannel
}

//launch game supervisor ( should be lauched last
func (g * GameSupervisor) Play() {
	for play := true; play; play = shared.Continue() {
		time.Sleep(FrameDelay_ms* time.Millisecond)
		g.DirectionChannel <- g.Mode.Move()
	}
}


