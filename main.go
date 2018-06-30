package main

import (
	"time"

	"github.com/gandrin/ASharedJourney/supervisor"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/gandrin/ASharedJourney/tiles"
	"golang.org/x/image/colornames"

	"github.com/gandrin/ASharedJourney/mechanics"
	"github.com/gandrin/ASharedJourney/menu"
	"github.com/gandrin/ASharedJourney/shared"
)

const (
	AMAZEING_LEVEL       string = "amazeing"
	FOREST_LEVEL         string = "forest"
	MY_LITTLE_PONY_LEVEL string = "myLittlePony"
	THE_LITTLE_PIG_LEVEL string = "theLittlePig"
)

const frameRate = 60

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "A Shared Journey",
		Bounds: pixel.R(0, 0, 800, 800),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	shared.Win = win

	menu.Menu()

	world := tiles.GenerateMap(AMAZEING_LEVEL)

	fps := time.Tick(time.Second / frameRate)

	playerDirectionChannel := supervisor.Start(supervisor.OnePlayer)

	newWorldChannel := mechanics.Start(playerDirectionChannel, world)

	for !win.Closed() {
		win.Clear(colornames.Black)
		supervisor.Sup.Play()
		mechanics.Mecha.Play()
		upToDateWorld := <-newWorldChannel
		tiles.DrawMap(upToDateWorld.BackgroundTiles)
		tiles.DrawMap(upToDateWorld.Obstacles)
		tiles.DrawMap(upToDateWorld.WinStars)
		tiles.DrawMap(upToDateWorld.Water)
		tiles.DrawMap(upToDateWorld.Movables)
		tiles.DrawMap(upToDateWorld.Players)
		tiles.DrawMap(upToDateWorld.Holes)
		win.Update()
		<-fps
	}
}

func main() {

	pixelgl.Run(run)
}
