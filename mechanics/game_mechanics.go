//defines the behavior of the game once we have defined the desired direction
package mechanics

import (
	"log"

	"github.com/gandrin/ASharedJourney/shared"
	"github.com/gandrin/ASharedJourney/supervisor"
)

type Mechanics struct {
	//player data
	Player1 playerManager
	Player2 playerManager

	hitMap [][]tileRules
	//location of event that can be trigged on the map
	eventMap [][]*eventType

	//communication channel to animator
	toAnime chan *Motion

	//communication channel from supervisor
	fromSuper chan *supervisor.PlayerDirections

	//all data relative to game status ( score , nb actions , ect ... ) is in game_status : call by func
}

var Mecha *Mechanics

func Start(fromSup chan *supervisor.PlayerDirections,
	p1 playerManager, p2 playerManager,
	hitmap [][]tileRules,
	eventmap [][]*eventType) chan *Motion {
	Mecha = new(Mechanics)
	var toAnim chan *Motion
	toAnim = make(chan *Motion, 1)
	Mecha.toAnime = toAnim
	Mecha.fromSuper = fromSup

	//load initial player positions + type
	Mecha.Player1 = p1
	Mecha.Player2 = p2

	//load maps
	Mecha.hitMap = hitmap
	Mecha.eventMap = eventmap

	log.Print("Mecanics loaded")
	return Mecha.toAnime
}

func (m *Mechanics) Play() {

	for play := true; play; play = shared.Continue() {
		//wait for next deplacement
		playDir := <-m.fromSuper
		log.Printf("Got direction ", playDir)
		m.Move(playDir)
	}
}
