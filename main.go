package main

import (
	"fmt"
	"time"

	"github.com/gandrin/ASharedJourney/shared"

	"github.com/gandrin/ASharedJourney/supervisor"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/gandrin/ASharedJourney/tiles"
	"golang.org/x/image/colornames"

	"log"
)

const frameRate = 60

func updatePlayer(win *pixelgl.Window, sprite *pixel.Sprite, playerDirectionChannel chan *supervisor.PlayerDirections) {
	go func(playerDirection chan *supervisor.PlayerDirections) {
		playerOldPosition := win.Bounds().Center()
		for true {
			newPlayerDirection := <-playerDirection
			fmt.Println(newPlayerDirection)
			playerNewPosition := pixel.V(
				playerOldPosition.X+float64(newPlayerDirection.Player1.X*16),
				playerOldPosition.Y+float64(newPlayerDirection.Player1.Y*16),
			)
			sprite.Draw(win, pixel.IM.Moved(playerNewPosition))
			playerOldPosition = playerNewPosition
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

	spritesheet, tilesFrames := tiles.GenerateMap(win)

	log.Print("Hello")
	sprite := pixel.NewSprite(spritesheet, tilesFrames[203])

	fps := time.Tick(time.Second / frameRate)

	shared.Win = win
	playerDirectionChannel := supervisor.Start(supervisor.OnePlayer)
	updatePlayer(win, sprite, playerDirectionChannel)

	for !win.Closed() {
		supervisor.Sup.Play()
		win.Update()
		<-fps
	}
}

func main() {
	log.Printf("jello")
	pixelgl.Run(run)
}
