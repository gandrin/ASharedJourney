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
const mapWidth = 32

// LoadPicture load the picture
func LoadPicture(path string) (pixel.Picture, error) {
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

// GenerateMap generates the map
func GenerateMap(win *pixelgl.Window) {
	// parse tmx file
	gameMap, err := tiled.LoadFromFile(mapPath)

	spritesheet, err := LoadPicture("tiles/tileset.png")
	if err != nil {
		panic(err)
	}

	var tilesFrames []pixel.Rect
	for y := spritesheet.Bounds().Max.Y - tileSize; y > spritesheet.Bounds().Min.Y; y -= tileSize {
		for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += tileSize {
			tilesFrames = append(tilesFrames, pixel.R(x, y, x+tileSize, y+tileSize))
		}
	}

	for index, layerTile := range gameMap.Layers[0].Tiles {
		sprite := pixel.NewSprite(spritesheet, tilesFrames[layerTile.ID])
		spritePosition := pixel.V(float64((index%mapWidth*tileSize)+tileSize/2), 0+tileSize/2)
		sprite.Draw(win, pixel.IM.Moved(spritePosition))
	}

	if err != nil {
		fmt.Println("Error parsing map")
		os.Exit(2)
	}

	fmt.Print(gameMap.Layers[0].Tiles[0])
}
