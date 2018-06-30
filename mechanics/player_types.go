package mechanics

type PlayerType string

const (
	FOX      PlayerType = "fox"
	MOUSE    PlayerType = "mouse"
	BEE      PlayerType = "bee"
	ELEPHANT PlayerType = "elephant"
)

func (pt PlayerType) can_walk(tileType tileRules) bool{
	//check if this type of player can walk on this type of tile
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
func (pt PlayerType) trigger_event(inialEvent *eventType) *Event{
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