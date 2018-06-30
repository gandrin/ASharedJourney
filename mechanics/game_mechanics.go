//defines the behavior of the game once we have defined the desired direction
package mechanics

import (
	"log"

	"github.com/gandrin/ASharedJourney/shared"
	"github.com/gandrin/ASharedJourney/supervisor"
	"fmt"
	"time"
	"github.com/gandrin/ASharedJourney/tiles"
)

type Mechanics struct {
	//player data
	Player1 PlayerManager
	Player2 PlayerManager

	hitMap [][]TileRules
	//location of event that can be trigged on the map
	eventMap [][]*EventType

	dynamicObject [][]*Object

	world tiles.World
	//communication channel to animator
	toAnime chan *tiles.World

	//communication channel from supervisor
	fromSuper chan *supervisor.PlayerDirections



	//all data relative to game status ( score , nb actions , ect ... ) is in game_status : call by func
}
//game mechanincs stringleton
var Mecha *Mechanics

//initialise the game mechanics structure
func Start(fromSup chan *supervisor.PlayerDirections,
	p1 PlayerManager, p2 PlayerManager, baseWorld tiles.World) chan *tiles.World {

	Mecha = new(Mechanics)
	//build return channel to animator
	var toAnim chan *tiles.World
	toAnim = make(chan *tiles.World, 1)

	Mecha.toAnime = toAnim
	Mecha.fromSuper = fromSup
	Mecha.world = baseWorld

	//load initial player positions + type
	Mecha.Player1 = p1
	Mecha.Player2 = p2

	//log.Print("Mecanics loaded")
	return Mecha.toAnime
}
//synchronisation objects
func (m *Mechanics) muxChannel() *supervisor.PlayerDirections {
	var nextMotion *supervisor.PlayerDirections
	select {
	case m, ok := <-m.fromSuper:
		if ok {
			fmt.Printf("Motion was read.")
			nextMotion = m
		} else {
			fmt.Println("Channel closed!")
			log.Fatal()
		}
	default:
		fmt.Println("No player direction mecha is faster than supervisor ")
		//set motion to default values
		nextMotion = new(supervisor.PlayerDirections)
		nextMotion.Player1.X = 0
		nextMotion.Player1.Y = 0
		nextMotion.Player2.X = 0
		nextMotion.Player2.Y = 0


	}
	return nextMotion
}

//call mechanics
func (m *Mechanics) Play() {

	for play := true; play; play = shared.Continue() {
		//delay to not call and overload cpu
		time.Sleep(shared.MechanicsRefreshDelay_ms * time.Millisecond)

		playDir := m.muxChannel()

		//log.Printf("Got direction ", playDir)
		m.toAnime <- m.Move(playDir)
	}
}
