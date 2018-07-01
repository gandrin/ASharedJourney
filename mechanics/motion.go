package mechanics

import (
	"reflect"

	"github.com/faiface/pixel"
	"github.com/gandrin/ASharedJourney/music"
	"github.com/gandrin/ASharedJourney/supervisor"
	"github.com/gandrin/ASharedJourney/tiles"
	"github.com/gandrin/ASharedJourney/menu"
)

//move function recives as input the data from a player direction channel
func (gm *Mechanics) Move(playDir *supervisor.PlayerDirections) *tiles.World {
	//log.Printf("Move called")

	if gm.world.Players[0].HasWon && gm.world.Players[1].HasWon &&
		!reflect.DeepEqual(gm.world.Players[0].WinningPosition, gm.world.Players[1].WinningPosition) {
		music.Music.PlayEffect(music.SOUND_EFFECT_WIN_GAME)
		gm.world = tiles.NextLevel()
	}

	if gm.world.Players[0].InTheWater || gm.world.Players[1].InTheWater {
		menu.Menu(menu.DrownedGameImage, "Oops ....", pixel.V(300,150),true, music.SOUND_EFFECT_LOSE_GAME)
		gm.world = tiles.RestartLevel()
	}

	if playDir.Player1.X != 0 || playDir.Player1.Y != 0 {
		// Zz tile... c'est moche mais bon...
		gm.world.Holes[len(gm.world.Holes)-1].Position.X = -100
		gm.world.Holes[len(gm.world.Holes)-1].Position.Y = -100

		gm.movePlayer(&gm.world.Players[0], playDir.Player1.Next)
		gm.movePlayer(&gm.world.Players[1], playDir.Player2.Next)

	}

	return gm.copyToNewWorld()
}

func (gm *Mechanics) movePlayer(player *tiles.SpriteWithPosition, getNextPosition func(pixel.Vec) pixel.Vec) {
	var canPlayerMove = true
	nextPlayerPosition := getNextPosition(player.Position)

	/// In the hole
	if player.InTheHole {
		player.InTheHole = false
		gm.world.Holes[len(gm.world.Holes)-1].Position = player.Position
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
				movableNextPosition := getNextPosition(nextPlayerPosition)
				for _, playerTile := range gm.world.Players {
					if playerTile.Position.X == movableNextPosition.X &&
						playerTile.Position.Y == movableNextPosition.Y {
						canPlayerMove = false
					}
				}
				for _, obstacleTile := range gm.world.Obstacles {
					if obstacleTile.Position.X == movableNextPosition.X &&
						obstacleTile.Position.Y == movableNextPosition.Y {
						canPlayerMove = false
					}
				}
				for _, movableTile := range gm.world.Movables {
					if movableTile.Position.X == movableNextPosition.X &&
						movableTile.Position.Y == movableNextPosition.Y {
						canPlayerMove = false
					}
				}
				for _, winStarTile := range gm.world.WinStars {
					if winStarTile.Position.X == movableNextPosition.X && winStarTile.Position.Y == movableNextPosition.Y {
						player.HasWon = true
						player.WinningPosition = pixel.V(winStarTile.Position.X, winStarTile.Position.Y)
					}
				}
				if canPlayerMove {
					gm.world.Movables[n].Position = movableNextPosition
				}
				for h, holeTile := range gm.world.Holes {
					if holeTile.Position.X == movableNextPosition.X && holeTile.Position.Y == movableNextPosition.Y {
						// remove both obj (hole and movable)
						gm.world.Movables[n].Position.X = -100
						gm.world.Holes[h].Position.X = -100
						music.Music.PlayEffect(music.SOUND_EFFECT_SNORE)
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
				music.Music.PlayEffect(music.SOUND_EFFECT_WATER)
			}
		}

		// Hole
		for _, holeTile := range gm.world.Holes {
			if holeTile.Position.X == nextPlayerPosition.X && holeTile.Position.Y == nextPlayerPosition.Y {
				player.InTheHole = true
				gm.world.Holes[len(gm.world.Holes)-1].Position = nextPlayerPosition
			}
		}

		player.HasWon = false

		// Winning rule
		for _, winStarTile := range gm.world.WinStars {
			if winStarTile.Position.X == nextPlayerPosition.X && winStarTile.Position.Y == nextPlayerPosition.Y {
				player.HasWon = true
				player.WinningPosition = pixel.V(winStarTile.Position.X, winStarTile.Position.Y)
			}
		}
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
