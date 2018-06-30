package mechanics

import (
	"log"

	"github.com/gandrin/ASharedJourney/shared"
)

//in game events
type Event struct {
	//fill with all the necessary elements to trigger events
	Dialog string

	//eg : to move objects
	//ObjectToMove []ObjectMotion

	Pos shared.Position
}

//list of events
type EventType string

const (
	eventHello EventType = "Hello"
	eventBy    EventType = "By"
)

//build a new event according to the events type
func NewEvent(newEventType EventType, ePos shared.Position) *Event {
	var newEvent = new(Event)
	newEvent.Pos = ePos
	//todo compleat with event mechanics struct implementation
	switch newEventType {
	case eventHello:
		newEvent.Dialog = string(eventHello)
		break
	case eventBy:
		newEvent.Dialog = string(eventBy)
		break
	default:
		log.Fatal("Unhandled event ", string(newEventType))
		break
	}
	return newEvent
}
