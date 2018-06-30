package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/gandrin/ASharedJourney/tiles"
	"golang.org/x/image/colornames"
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

	tiles.GenerateMap(win)

	fps := time.Tick(time.Second / frameRate)

	for !win.Closed() {
		win.Update()
		<-fps
	}
}

func main() {
	pixelgl.Run(run)
}
