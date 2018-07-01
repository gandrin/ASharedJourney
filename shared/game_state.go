package shared



//struct that holds all data about the current game state
type gameState struct {
	//is the game to contrinue
	Playing  bool
	Level    int
	NbAction int
	Score    int
}

var gState gameState

func StartGame(newlevel int ){
	gState.Playing = true
	gState.Score = 0
	gState.Level = newlevel
	gState.NbAction = 0
}
func StopGame(){
	gState.Playing = false
}
func AddAction()  {
	gState.NbAction +=1
	//log.Print("Actions ",gState.NbAction)
}
func Continue() bool  {
	return gState.Playing
}