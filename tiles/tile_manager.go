package tiles

import (
	"errors"
	"image"
	"os"

	"github.com/gandrin/ASharedJourney/shared"

	_ "image/png"

	"log"

	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lafriks/go-tiled"
)

// Level names
const (
	amazeingLevel     string = "amazeing"
	forestLevel       string = "forest"
	myLittlePonyLevel string = "myLittlePony"
	theLittlePigLevel string = "theLittlePig"
)

// CurrentLevel played
var CurrentLevel = -1

// Levels list
var Levels = [...]string{amazeingLevel, forestLevel, myLittlePonyLevel, theLittlePigLevel}

const tilesPath = "/tiles/map.png" // path to your tileset
var TileSize int = 32
var mapWidth int
var mapHeight int

type World struct {
	BackgroundTiles []SpriteWithPosition
	Players         []SpriteWithPosition
	Movables        []SpriteWithPosition
	Obstacles       []SpriteWithPosition
	Water           []SpriteWithPosition
	Holes           []SpriteWithPosition
	WinStars        []SpriteWithPosition
}

//SpriteWithPosition holds the sprite and its position into the window
type SpriteWithPosition struct {
	Sprite     *pixel.Sprite
	Position   pixel.Vec
	InTheWater bool
	InTheHole  bool
	HasWon     bool
}

// loadPicture load the picture
func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func getTilesFrames(spritesheet pixel.Picture) []pixel.Rect {
	var tilesFrames []pixel.Rect
	for y := spritesheet.Bounds().Max.Y - float64(TileSize); y > spritesheet.Bounds().Min.Y-float64(TileSize); y -= float64(TileSize) {
		for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += float64(TileSize) {
			tilesFrames = append(tilesFrames, pixel.R(x, y, x+float64(TileSize), y+float64(TileSize)))
		}
	}

	return tilesFrames
}

func getOrigin(win *pixelgl.Window) pixel.Vec {
	centerPosition := win.Bounds().Center()
	originXPosition := centerPosition.X - float64(mapWidth)/2*float64(TileSize)
	originYPosition := centerPosition.Y + float64(mapHeight)/2*float64(TileSize) - float64(TileSize)

	return pixel.V(originXPosition, originYPosition)
}

func getSpritePosition(spriteIndex int, origin pixel.Vec) pixel.Vec {
	spriteXPosition := origin.X + float64((spriteIndex%mapWidth)*TileSize) + float64(TileSize)/2
	spriteYPosition := origin.Y + float64(TileSize)/2 - float64((spriteIndex/mapWidth)*TileSize)

	return pixel.V(spriteXPosition, spriteYPosition)
}

// extractAndPlaceSprites filters out empty tiles and positions them properly on the screen
func extractAndPlaceSprites(
	layerTiles []*tiled.LayerTile,
	spritesheet pixel.Picture,
	tilesFrames []pixel.Rect,
	originPosition pixel.Vec,
) (positionedSprites []SpriteWithPosition) {
	for index, layerTile := range layerTiles {
		if !layerTile.IsNil() {
			sprite := pixel.NewSprite(spritesheet, tilesFrames[layerTile.ID])
			spritePosition := getSpritePosition(index, originPosition)
			positionedSprites = append(positionedSprites, SpriteWithPosition{
				Sprite:   sprite,
				Position: spritePosition,
			})
		}
	}
	return positionedSprites
}

func findLayerIndex(layerName string, layers []*tiled.Layer) (layerIndex int, err error) {
	for index, layer := range layers {
		if layer.Name == layerName {
			return index, nil
		}
	}
	return -1, errors.New("Expected to find layer with name " + layerName)
}

// NextLevel goes to next level
func NextLevel() World {
	CurrentLevel = (CurrentLevel + 1) % len(Levels)
	return GenerateMap(Levels[CurrentLevel])
}

// RestartLevel reinitializes the current level
func RestartLevel() World {
	return GenerateMap(Levels[CurrentLevel])
}

// GenerateMap generates the map from a .tmx file
func GenerateMap(levelFileName string) World {
	//added support for relative file addressing
	rootDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal("error loading called")
	}
	filemap := rootDirectory + "/tiles/" + levelFileName + ".tmx"
	filetile := rootDirectory + tilesPath
	gameMap, err := tiled.LoadFromFile(filemap)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error parsing map")
		os.Exit(2)
	}
	mapWidth = gameMap.Width
	mapHeight = gameMap.Height

	spritesheet, err := loadPicture(filetile)
	if err != nil {
		panic(err)
	}

	tilesFrames := getTilesFrames(spritesheet)

	originPosition := getOrigin(shared.Win)

	backgroundLayerIndex, err := findLayerIndex("background", gameMap.Layers)
	if err != nil {
		panic(err)
	}
	playersLayerIndex, err := findLayerIndex("animals", gameMap.Layers)
	if err != nil {
		panic(err)
	}
	obstaclesLayerIndex, err := findLayerIndex("obstacles", gameMap.Layers)
	if err != nil {
		panic(err)
	}
	movablesLayerIndex, err := findLayerIndex("movables", gameMap.Layers)
	if err != nil {
		panic(err)
	}
	waterLayerIndex, err := findLayerIndex("water", gameMap.Layers)
	if err != nil {
		panic(err)
	}
	winStarsLayerIndex, err := findLayerIndex("win", gameMap.Layers)
	if err != nil {
		panic(err)
	}
	holesLayerIndex, err := findLayerIndex("holes", gameMap.Layers)
	if err != nil {
		panic(err)
	}

	backgroundSprite := extractAndPlaceSprites(gameMap.Layers[backgroundLayerIndex].Tiles, spritesheet, tilesFrames, originPosition)
	players := extractAndPlaceSprites(gameMap.Layers[playersLayerIndex].Tiles, spritesheet, tilesFrames, originPosition)
	if len(players) == 0 {
		panic(errors.New("no animal tile was placed"))
	}
	obstacles := extractAndPlaceSprites(gameMap.Layers[obstaclesLayerIndex].Tiles, spritesheet, tilesFrames, originPosition)
	movables := extractAndPlaceSprites(gameMap.Layers[movablesLayerIndex].Tiles, spritesheet, tilesFrames, originPosition)
	water := extractAndPlaceSprites(gameMap.Layers[waterLayerIndex].Tiles, spritesheet, tilesFrames, originPosition)
	holes := extractAndPlaceSprites(gameMap.Layers[holesLayerIndex].Tiles, spritesheet, tilesFrames, originPosition)
	winStars := extractAndPlaceSprites(gameMap.Layers[winStarsLayerIndex].Tiles, spritesheet, tilesFrames, originPosition)
	if len(winStars) == 0 {
		panic(errors.New("no win star tile was placed"))
	}

	world := World{
		BackgroundTiles: backgroundSprite,
		Players:         players,
		Movables:        movables,
		Obstacles:       obstacles,
		Water:           water,
		Holes:           holes,
		WinStars:        winStars,
	}
	return world
}

//DrawMap draws into window the given sprites
func DrawMap(positionedSprites []SpriteWithPosition) {
	for _, positionedSprite := range positionedSprites {
		positionedSprite.Sprite.Draw(shared.Win, pixel.IM.Moved(positionedSprite.Position))
	}
}
