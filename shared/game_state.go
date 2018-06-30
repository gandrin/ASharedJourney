package shared

import "log"

//struct that holds all data about the current game state
type gameState struct {
	//is the game to contrinue
	playing bool
	level int
	nActions int
	score int
}

var gState gameState

func StartGame(newlevel int ){
	gState.playing = true
	gState.score = 0
	gState.level = newlevel
	gState.nActions = 0
}
func StopGame(){
	gState.playing = false
}
func AddAction()  {
	gState.nActions +=1
	log.Print("Actions ",gState.nActions)
}
func Continue() bool  {
	return gState.playing
}