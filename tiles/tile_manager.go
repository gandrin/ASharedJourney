package tiles

import (
	"fmt"
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lafriks/go-tiled"
)

const mapPath = "tiles/tilemap.tmx" // path to your map
const tileSize = 16
const mapWidth = 30

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

func getSpritePosition(spriteIndex int) pixel.Vec {
	spriteXPosition := (spriteIndex%mapWidth)*tileSize + tileSize/2
	spriteYPosition := (spriteIndex/mapWidth)*tileSize + tileSize/2

	return pixel.V(float64(spriteXPosition), float64(spriteYPosition))
}

// GenerateMap generates the map
func GenerateMap(win *pixelgl.Window) {
	// parse tmx file
	gameMap, err := tiled.LoadFromFile(mapPath)
	if err != nil {
		fmt.Println("Error parsing map")
		os.Exit(2)
	}

	spritesheet, err := loadPicture("tiles/tileset.png")
	if err != nil {
		panic(err)
	}

	tilesFrames := getTilesFrames(spritesheet)

	for index, layerTile := range gameMap.Layers[0].Tiles {
		sprite := pixel.NewSprite(spritesheet, tilesFrames[layerTile.ID])
		spritePosition := getSpritePosition(index)
		sprite.Draw(win, pixel.IM.Moved(spritePosition))
	}
}
