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

	var treesFrames []pixel.Rect
	for y := spritesheet.Bounds().Max.Y - 16; y > spritesheet.Bounds().Min.Y; y -= 16 {
		for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 16 {
			treesFrames = append(treesFrames, pixel.R(x, y, x+16, y+16))
		}
	}

	tree := pixel.NewSprite(spritesheet, treesFrames[0])
	tree.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	if err != nil {
		fmt.Println("Error parsing map")
		os.Exit(2)
	}

	fmt.Print(gameMap.Tilesets[0])
}
