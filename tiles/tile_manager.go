package tiles

import (
	"errors"
	"image"
	"os"
	"path"

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
	splashScreen         string = "orgyIsland"
	veryEasyLevel        string = "veryEasy"
	mazeLevel            string = "maze"
	theLongCorridorLevel string = "theLongCorridor"
	amazeingLevel        string = "amazeing"
	forestLevel          string = "forest"
	myLittlePonyLevel    string = "myLittlePony"
	theLittlePigLevel    string = "theLittlePig"
	inTheHoleTileID      int    = 61
)

// CurrentLevel played
var CurrentLevel = -1

// Levels list
var Levels = [...]string{veryEasyLevel, mazeLevel, theLongCorridorLevel, amazeingLevel, theLittlePigLevel, myLittlePonyLevel, forestLevel}

// Uncomment this for testing :)
// var Levels = [...]string{biggerLevel}

const tilesPath = "/map.png" // path to your tileset

// TileSize tile in pixels of squares
var TileSize int
var mapWidth int
var mapHeight int

// World struct with global tiles and obj positions
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
	filemap := path.Join(".", "assets", levelFileName+".tmx")
	filetile := path.Join(".", "assets", tilesPath)
	gameMap, err := tiled.LoadFromFile(filemap)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error parsing map")
		os.Exit(2)
	}
	mapWidth = gameMap.Width
	mapHeight = gameMap.Height
	TileSize = gameMap.TileHeight

	spritesheet, err := loadPicture(filetile)
	if err != nil {
		panic(err)
	}

	tilesFrames := getTilesFrames(spritesheet)

	originPosition := getOrigin(shared.Win)

	var background []SpriteWithPosition
	var players []SpriteWithPosition
	var obstacles []SpriteWithPosition
	var movables []SpriteWithPosition
	var water []SpriteWithPosition
	var winStars []SpriteWithPosition
	var holes []SpriteWithPosition

	backgoundIndex, err := findLayerIndex("background", gameMap.Layers)
	if err == nil {
		background = extractAndPlaceSprites(gameMap.Layers[backgoundIndex].Tiles, spritesheet, tilesFrames, originPosition)
	}
	playersLayerIndex, err := findLayerIndex("animals", gameMap.Layers)
	if err == nil {
		players = extractAndPlaceSprites(gameMap.Layers[playersLayerIndex].Tiles, spritesheet, tilesFrames, originPosition)
	}
	obstaclesLayerIndex, err := findLayerIndex("obstacles", gameMap.Layers)
	if err == nil {
		obstacles = extractAndPlaceSprites(gameMap.Layers[obstaclesLayerIndex].Tiles, spritesheet, tilesFrames, originPosition)
	}
	movablesLayerIndex, err := findLayerIndex("movables", gameMap.Layers)
	if err == nil {
		movables = extractAndPlaceSprites(gameMap.Layers[movablesLayerIndex].Tiles, spritesheet, tilesFrames, originPosition)
	}
	waterLayerIndex, err := findLayerIndex("water", gameMap.Layers)
	if err == nil {
		water = extractAndPlaceSprites(gameMap.Layers[waterLayerIndex].Tiles, spritesheet, tilesFrames, originPosition)
	}
	winStarsLayerIndex, err := findLayerIndex("win", gameMap.Layers)
	if err == nil {
		winStars = extractAndPlaceSprites(gameMap.Layers[winStarsLayerIndex].Tiles, spritesheet, tilesFrames, originPosition)
	}
	holesLayerIndex, err := findLayerIndex("holes", gameMap.Layers)
	if err == nil {
		holes = extractAndPlaceSprites(gameMap.Layers[holesLayerIndex].Tiles, spritesheet, tilesFrames, originPosition)
	}

	// Playable checks
	if len(players) == 0 {
		panic(errors.New("no animal tile was placed"))
	}

	if len(winStars) == 0 {
		panic(errors.New("no win star tile was placed"))
	}

	// Zz special tile
	sprite := pixel.NewSprite(spritesheet, tilesFrames[inTheHoleTileID])
	spritePosition := pixel.V(-100, -100)
	holes = append(holes, SpriteWithPosition{
		Sprite:   sprite,
		Position: spritePosition,
	})

	world := World{
		BackgroundTiles: background,
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
