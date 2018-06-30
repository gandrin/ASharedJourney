package main

import (
	"fmt"
	"time"

	"github.com/gandrin/ASharedJourney/shared"

	"github.com/gandrin/ASharedJourney/supervisor"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"github.com/ASharedJourney/tiles"
)

const frameRate = 60

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 500, 500),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.White)

	spritesheet, tilesFrames := tiles.GenerateMap(win)

	sprite := pixel.NewSprite(spritesheet, tilesFrames[203])
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	fps := time.Tick(time.Second / frameRate)

	shared.Win = win
	playerDirectionChannel := supervisor.Start(supervisor.OnePlayer)
	go func(playerDirection chan *supervisor.PlayerDirections) {
		for true {
			newPlayerDirection := <-playerDirection
			fmt.Println(newPlayerDirection.Player1)
		}
	}(playerDirectionChannel)

	for !win.Closed() {
		supervisor.Sup.Play()
		win.Update()
		<-fps
	}
}

func main() {
	pixelgl.Run(run)
}
