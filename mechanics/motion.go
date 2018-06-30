package mechanics

import (
	"github.com/faiface/pixel"
	"github.com/gandrin/ASharedJourney/supervisor"
	"github.com/gandrin/ASharedJourney/tiles"
)

func (gm *Mechanics) handlePlayerWon(nextPos1 pixel.Vec) {
	for _, winStartTile := range gm.world.WinStars {
		if winStartTile.Position.X == nextPos1.X && winStartTile.Position.Y == nextPos1.Y {
			gm.world.Players[0].HasWon = true
		}
	}
}

//move function recives as input the data from a player direction channel
func (gm *Mechanics) Move(playDir *supervisor.PlayerDirections) *tiles.World {
	//log.Printf("Move called")

	if gm.world.Players[0].HasWon {
<<<<<<< HEAD
		gm.world = tiles.GenerateMap("forest") // TODO next level
=======
		gm.world = tiles.NextLevel()
>>>>>>> Switch to next level on death
	}

	if gm.world.Players[0].InTheWater && gm.world.Players[1].InTheWater {
		gm.world = tiles.GenerateMap("forest") // TODO same game RESTART
	}
	if playDir.Player1.X != 0 || playDir.Player1.Y != 0 {
		gm.movePlayer(&gm.world.Players[0], playDir.Player1.Next)
		gm.movePlayer(&gm.world.Players[1], playDir.Player2.Next)
	}

	return gm.copyToNewWorld()
}

func (gm *Mechanics) movePlayer(player *tiles.SpriteWithPosition, getNextPosition func(pixel.Vec) pixel.Vec) {
	var canPlayerMove = true
	nextPlayerPosition := getNextPosition(player.Position)

	// In the water
	if player.InTheWater {
		return
	}

	/// In the hole
	if player.InTheHole {
		player.InTheHole = false
		return
	}

	// Obstacles
	for _, obstacle := range gm.world.Obstacles {
		if obstacle.Position.X == nextPlayerPosition.X && obstacle.Position.Y == nextPlayerPosition.Y {
			canPlayerMove = false
		}
	}

	// Movables
	if canPlayerMove {
		for n, mov := range gm.world.Movables {
			if mov.Position.X == nextPlayerPosition.X && mov.Position.Y == nextPlayerPosition.Y {
				// There's a movable in that position
				auxPos := getNextPosition(nextPlayerPosition)
				for _, obstacleTile := range gm.world.Obstacles {
					if obstacleTile.Position.X == auxPos.X && obstacleTile.Position.Y == auxPos.Y {
						canPlayerMove = false
					}
				}
				for _, obstacleTile := range gm.world.Movables {
					if obstacleTile.Position.X == auxPos.X && obstacleTile.Position.Y == auxPos.Y {
						canPlayerMove = false
					}
				}
				for _, winStarTile := range gm.world.WinStars {
					if winStarTile.Position.X == auxPos.X && winStarTile.Position.Y == auxPos.Y {
						gm.world.Players[0].HasWon = true
					}
				}
				if canPlayerMove {
					gm.world.Movables[n].Position = auxPos
				}
				for h, holeTile := range gm.world.Holes {
					if holeTile.Position.X == auxPos.X && holeTile.Position.Y == auxPos.Y {
						// remove both obj (hole and movable)
						gm.world.Movables[n].Position.X = -100
						gm.world.Holes[h].Position.X = -100
					}
				}
			}
		}
	}

	if canPlayerMove {
		player.Position = nextPlayerPosition

		// Water
		for _, waterTile := range gm.world.Water {
			if waterTile.Position.X == nextPlayerPosition.X && waterTile.Position.Y == nextPlayerPosition.Y {
				player.InTheWater = true
			}
		}

		// Hole
		for _, holeTile := range gm.world.Holes {
			if holeTile.Position.X == nextPlayerPosition.X && holeTile.Position.Y == nextPlayerPosition.Y {
				player.InTheHole = true
			}
		}
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
	newWorld.Holes = make([]tiles.SpriteWithPosition, len(gm.world.Holes))
	newWorld.WinStars = make([]tiles.SpriteWithPosition, len(gm.world.WinStars))
	copy(newWorld.BackgroundTiles, gm.world.BackgroundTiles)
	copy(newWorld.Movables, gm.world.Movables)
	copy(newWorld.Players, gm.world.Players)
	copy(newWorld.WinStars, gm.world.WinStars)
	copy(newWorld.Water, gm.world.Water)
	copy(newWorld.Obstacles, gm.world.Obstacles)
	copy(newWorld.Holes, gm.world.Holes)
	return newWorld
}
