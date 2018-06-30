package mechanics

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/gandrin/ASharedJourney/supervisor"
	"github.com/gandrin/ASharedJourney/tiles"
)

func (gm *Mechanics) handlePlayerWon(nextPos1 pixel.Vec) {
	for _, val := range gm.world.WinStars {
		if val.Position.X == nextPos1.X && val.Position.Y == nextPos1.Y {
			log.Printf("you won!")
		}
	}
}

//move function recives as input the data from a player direction channel
func (gm *Mechanics) Move(playDir *supervisor.PlayerDirections) *tiles.World {
	//log.Printf("Move called")

	gm.movePlayer(&gm.world.Players[0], playDir.Player1.Next)
	gm.movePlayer(&gm.world.Players[1], playDir.Player2.Next)

	return gm.copyToNewWorld()
}

func (gm *Mechanics) movePlayer(player *tiles.SpriteWithPosition, getNextPosition func(pixel.Vec) pixel.Vec) {
	var canPlayerMove = true
	nextPlayerPosition := getNextPosition(player.Position)

	// Obstacles
	for _, obstacle := range gm.world.Obstacles {
		if obstacle.Position.X == nextPlayerPosition.X && obstacle.Position.Y == nextPlayerPosition.Y {
			canPlayerMove = false
		}
	}

	// Movables
	if canPlayerMove {
		for n, val := range gm.world.Movables {
			if val.Position.X == nextPlayerPosition.X && val.Position.Y == nextPlayerPosition.Y {
				// There's a movable in that position
				auxPos := getNextPosition(nextPlayerPosition)
				for _, val := range gm.world.Obstacles {
					if val.Position.X == auxPos.X && val.Position.Y == auxPos.Y {
						canPlayerMove = false
					}
				}
				for _, val := range gm.world.Movables {
					if val.Position.X == auxPos.X && val.Position.Y == auxPos.Y {
						canPlayerMove = false
					}
				}
				if canPlayerMove {
					gm.world.Movables[n].Position = auxPos
				}
			}
		}
	}

	if canPlayerMove {
		player.Position = nextPlayerPosition
		gm.handlePlayerWon(player.Position)
	}
}

func (gm *Mechanics) copyToNewWorld() *tiles.World {
	var newWorld *tiles.World = new(tiles.World)
	//copy player locations

	//copy world
	//make a copy of world : todo check if doesn't fail
	//this will have to be updated
	newWorld.BackgroundTiles = make([]tiles.SpriteWithPosition, len(gm.world.BackgroundTiles))
	newWorld.Movables = make([]tiles.SpriteWithPosition, len(gm.world.Movables))
	newWorld.Players = make([]tiles.SpriteWithPosition, len(gm.world.Players))
	newWorld.Obstacles = make([]tiles.SpriteWithPosition, len(gm.world.Obstacles))
	newWorld.Water = make([]tiles.SpriteWithPosition, len(gm.world.Water))
	newWorld.WinStars = make([]tiles.SpriteWithPosition, len(gm.world.WinStars))
	copy(newWorld.BackgroundTiles, gm.world.BackgroundTiles)
	copy(newWorld.Movables, gm.world.Movables)
	copy(newWorld.Players, gm.world.Players)
	copy(newWorld.Water, gm.world.Water)
	copy(newWorld.Obstacles, gm.world.Obstacles)
	copy(newWorld.WinStars, gm.world.WinStars)
	return newWorld
}
