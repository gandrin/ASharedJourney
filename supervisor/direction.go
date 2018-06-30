package supervisor

import "log"

//call motion
func (mode GameMode)Move() *PlayerDirections{
	var newDir *PlayerDirections = new(PlayerDirections)
	if(mode == OnePlayer){
		newDir.Player1 = key()
		newDir.mirror()
	}else{
		log.Fatal("Unknown mode")
	}
	return newDir
}

//mirror motion of player 1 onto direction of player 2
func (dir * PlayerDirections) mirror(){
	if dir.Player1.X!= 0{
		dir.Player2.X = dir.Player1.X*(-1)
	}
	if dir.Player1.Y!= 0{
		dir.Player2.Y = dir.Player1.Y*(-1)
	}
}