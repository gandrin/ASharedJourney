package main

import (
	"time"

	"github.com/gandrin/ASharedJourney/shared"

	"github.com/gandrin/ASharedJourney/supervisor"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/gandrin/ASharedJourney/tiles"
	"golang.org/x/image/colornames"
)

const frameRate = 60

func updatePlayer(win *pixelgl.Window, sprite *pixel.Sprite, playerDirectionChannel chan *supervisor.PlayerDirections, playerPosition *pixel.Vec) {
	go func(playerDirection chan *supervisor.PlayerDirections) {
		for true {
			newPlayerDirection := <-playerDirection
			playerNewPosition := pixel.V(
				playerPosition.X+float64(newPlayerDirection.Player1.X*16),
				playerPosition.Y+float64(newPlayerDirection.Player1.Y*16),
			)
			playerPosition.X = playerNewPosition.X
			playerPosition.Y = playerNewPosition.Y
		}
	}(playerDirectionChannel)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "A Shared Journey",
		Bounds: pixel.R(0, 0, 500, 500),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.White)
	shared.Win = win

	spritesheet, tilesFrames, world := tiles.GenerateMap()

	sprite := pixel.NewSprite(spritesheet, tilesFrames[203])

	fps := time.Tick(time.Second / frameRate)

	playerDirectionChannel := supervisor.Start(supervisor.OnePlayer)

	playerNewPosition := world.Players[0].Position
	updatePlayer(win, sprite, playerDirectionChannel, &playerNewPosition)
	for !win.Closed() {
		supervisor.Sup.Play()
		tiles.DrawMap(world.BackgroundTiles)
		world.Players[0].Sprite.Draw(win, pixel.IM.Moved(playerNewPosition))
		win.Update()
		<-fps
	}
}

func main() {
	pixelgl.Run(run)
}
