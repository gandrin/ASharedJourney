//defines the behavior of the game once we have defined the desired direction
package mechanics

import (
	"fmt"
	"log"
	"time"

	"github.com/gandrin/ASharedJourney/shared"
	"github.com/gandrin/ASharedJourney/supervisor"
	"github.com/gandrin/ASharedJourney/tiles"
)

type Mechanics struct {
	world tiles.World
	//communication channel to animator
	toAnime chan *tiles.World
	//communication channel from supervisor
	gameEventChannel chan *supervisor.GameEvent
}

//game mechanincs stringleton
var Mecha *Mechanics

//initialise the game mechanics structure
func Start(
	gameEventChannel chan *supervisor.GameEvent,
	baseWorld tiles.World,
) chan *tiles.World {
	Mecha = new(Mechanics)
	//build return channel to animator
	var toAnim chan *tiles.World
	toAnim = make(chan *tiles.World, 1)

	Mecha.toAnime = toAnim
	Mecha.gameEventChannel = gameEventChannel
	Mecha.world = baseWorld

	//log.Print("Mecanics loaded")
	return Mecha.toAnime
}

//synchronisation objects
func (mechanics *Mechanics) muxChannel() *supervisor.GameEvent {
	select {
	case nextGameEvent, ok := <-mechanics.gameEventChannel:
		fmt.Println("HERE")
		fmt.Println(nextGameEvent)
		if !ok {
			fmt.Println("Channel  closed!")
			log.Fatal()
		}
		return nextGameEvent
	default:
		nextEvent := supervisor.Event("NONE")
		nextGameEvent := new(supervisor.GameEvent)
		nextGameEvent.PlayerDirections = new(supervisor.PlayerDirections)
		nextGameEvent.PlayerDirections.Player1.X = 0
		nextGameEvent.PlayerDirections.Player1.Y = 0
		nextGameEvent.PlayerDirections.Player2.X = 0
		nextGameEvent.PlayerDirections.Player2.Y = 0
		nextGameEvent.Event = &nextEvent
		fmt.Println("No player direction mecha is faster than supervisor ")
		return nextGameEvent
		//set motion to default values
	}
}

//call mechanics
func (mechanics *Mechanics) Play() {

	for play := true; play; play = shared.Continue() {
		//delay to not call and overload cpu
		time.Sleep(shared.MechanicsRefreshDelay_ms * time.Millisecond)

		gameEvent := mechanics.muxChannel()
		mechanics.toAnime <- mechanics.Move(gameEvent.PlayerDirections)
		mechanics.handleGameEvent(gameEvent.Event)
	}
}

func (mechanics *Mechanics) handleGameEvent(event *supervisor.Event) {
	switch *event {
	case "RESTART":
		mechanics.world = tiles.RestartLevel()
		break
	default:
		//No event
	}
}
