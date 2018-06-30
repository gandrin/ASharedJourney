package tiles

import (
	"fmt"
	"image"
	"os"

	"github.com/gandrin/ASharedJourney/shared"

	_ "image/png"

	"log"
	"path"
	"runtime"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lafriks/go-tiled"
)

const mapPath = "tiles/tilemap.tmx"   // path to your map
const tilesPath = "tiles/tileset.png" // path to your tileset
const tileSize = 16
const mapWidth = 30
const mapHeight = 30

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
	for y := spritesheet.Bounds().Max.Y - tileSize; y > spritesheet.Bounds().Min.Y; y -= tileSize {
		for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += tileSize {
			tilesFrames = append(tilesFrames, pixel.R(x, y, x+tileSize, y+tileSize))
		}
	}

	return tilesFrames
}

//SpriteWithPosition holds the sprite and its position into the window
type SpriteWithPosition struct {
	Sprite   *pixel.Sprite
	Position pixel.Vec
}

func getOrigin(win *pixelgl.Window) pixel.Vec {
	centerPosition := win.Bounds().Center()
	originXPosition := centerPosition.X - mapWidth/2*tileSize
	originYPosition := centerPosition.Y + mapHeight/2*tileSize - tileSize

	return pixel.V(originXPosition, originYPosition)
}

func getSpritePosition(spriteIndex int, origin pixel.Vec) pixel.Vec {
	spriteXPosition := origin.X + float64((spriteIndex%mapWidth)*tileSize) + tileSize/2
	spriteYPosition := origin.Y + tileSize/2 - float64((spriteIndex/mapWidth)*tileSize)

	return pixel.V(spriteXPosition, spriteYPosition)
}

// GenerateMap generates the map from a .tmx file
func GenerateMap() (pixel.Picture, []pixel.Rect, [mapWidth * mapHeight]SpriteWithPosition) {
	//get path to file from current programme root
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("error loading called")
	}
	filemap := path.Join(path.Dir(filename), mapPath)
	filetile := path.Join(path.Dir(filename), tilesPath)

	gameMap, err := tiled.LoadFromFile(filemap)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error parsing map")
		os.Exit(2)
	}

	spritesheet, err := loadPicture(filetile)
	if err != nil {
		panic(err)
	}

	tilesFrames := getTilesFrames(spritesheet)

	originPosition := getOrigin(shared.Win)

	var positionedSprites [mapWidth * mapHeight]SpriteWithPosition
	for index, layerTile := range gameMap.Layers[0].Tiles {
		sprite := pixel.NewSprite(spritesheet, tilesFrames[layerTile.ID])
		spritePosition := getSpritePosition(index, originPosition)
		positionedSprites[index] = SpriteWithPosition{
			Sprite:   sprite,
			Position: spritePosition,
		}
	}

	return spritesheet, tilesFrames, positionedSprites
}

//DrawMap draws into window the given sprites
func DrawMap(positionedSprites [mapWidth * mapHeight]SpriteWithPosition) {
	for _, positionedSprite := range positionedSprites {
		positionedSprite.Sprite.Draw(shared.Win, pixel.IM.Moved(positionedSprite.Position))
	}
}
