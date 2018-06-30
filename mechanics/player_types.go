package mechanics

import "log"

type PlayerType string

const (
	FOX      PlayerType = "fox"
	MOUSE    PlayerType = "mouse"
	BEE      PlayerType = "bee"
	ELEPHANT PlayerType = "elephant"
)

func (pt PlayerType) can_walk(tileType TileRules) bool{
	//check if this type of player can walk on this type of tile
	log.Print("Checking if player of type ",pt, " can walk of tile ", tileType)
	var retVal bool = true
	switch pt {
	case BEE:
		break
	case MOUSE:
		break
	case ELEPHANT:
		break
	case FOX:
		break
	}
	return retVal
}

//can block / modify event according to player
func (pt PlayerType) trigger_event(inialEvent *EventType) *Event{
	//check if this type of player can walk on this type of tile
	var retVal *Event
	switch pt {
	case BEE:
		break
	case MOUSE:
		break
	case ELEPHANT:
		break
	case FOX:
		break
	}
	return retVal
}