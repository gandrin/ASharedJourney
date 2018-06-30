package main

import (
	"time"

	"github.com/gandrin/ASharedJourney/shared"

	"github.com/gandrin/ASharedJourney/supervisor"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/gandrin/ASharedJourney/tiles"
	"golang.org/x/image/colornames"

	"github.com/gandrin/ASharedJourney/mechanics"
)

const frameRate = 60

func updatePlayer(win *pixelgl.Window, sprite *pixel.Sprite, playerDirectionChannel chan *supervisor.PlayerDirections, playerPosition *pixel.Vec) {
	go func(playerDirection chan *supervisor.PlayerDirections) {
		for true {
			newPlayerDirection := <-playerDirection
			playerNewPosition := pixel.V(
				playerPosition.X+float64(newPlayerDirection.Player1.X*32),
				playerPosition.Y+float64(newPlayerDirection.Player1.Y*32),
			)
			playerPosition.X = playerNewPosition.X
			playerPosition.Y = playerNewPosition.Y
		}
	}(playerDirectionChannel)
}

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

	world := tiles.GenerateMap()

	fps := time.Tick(time.Second / frameRate)

	playerDirectionChannel := supervisor.Start(supervisor.OnePlayer)

	//game mechanics
	var p1 mechanics.PlayerManager
	p1.Pos.X = 10
	p1.Pos.Y = 10
	p1.PType = mechanics.FOX
	var p2 mechanics.PlayerManager
	p2.Pos.X = 20
	p2.Pos.Y = 20
	p2.PType = mechanics.BEE
	//init rules map
	ruleMap := make([][]mechanics.TileRules, 40)
	for i := 0; i < 40; i++ {
		ruleMap[i] = make([]mechanics.TileRules, 40)
		for j := 0; j < 40; j++ {
			ruleMap[i][j] = 0
		}

	}
	//init event map
	eventMap := make([][]*mechanics.EventType, 40)
	for i := 0; i < 40; i++ {
		eventMap[i] = make([]*mechanics.EventType, 40)
		for j := 0; j < 40; j++ {
			eventMap[i][j] = nil // no events set
		}

	}
	//init object map
	objMap := make([][]*mechanics.Object, 40)
	for i := 0; i < 40; i++ {
		objMap[i] = make([]*mechanics.Object, 40)
		for j := 0; j < 40; j++ {
			objMap[i][j] = nil // no events set
		}

	}
	_ = mechanics.Start(playerDirectionChannel, p1, p2, ruleMap, eventMap, objMap)

	playerNewPosition := world.Players[0].Position

	//if you want direction to work comment out this line but lose animations
	updatePlayer(win, world.Players[0].Sprite, playerDirectionChannel, &playerNewPosition)
	for !win.Closed() {
		win.Clear(colornames.White)
		supervisor.Sup.Play()
		//mechanics.Mecha.Play()
		tiles.DrawMap(world.BackgroundTiles)
		world.Players[0].Sprite.Draw(win, pixel.IM.Moved(playerNewPosition))
		win.Update()
		<-fps
	}
}

func main() {
	pixelgl.Run(run)
}
