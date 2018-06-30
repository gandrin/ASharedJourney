package menu

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"github.com/faiface/pixel"
	"fmt"
	"golang.org/x/image/colornames"
	"github.com/gandrin/ASharedJourney/shared"
	"os"
	"image"
	"runtime"
	"log"
	"path"
	"time"
)

const menuText string = "Press ENTER to PLAY"
const menuTextPosX float64 = 200
const menuTextPosY float64 = 100
const menuPicName 	string = "tiles/menu.png"

//draw menu to screen while player while player hasn't pressed enter
func Menu(){

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(menuTextPosX, menuTextPosY), basicAtlas)
	basicTxt.Color = colornames.White
	fmt.Fprintln(basicTxt, menuText)

	//get picture
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("error loading called")
	}
	menupicture:= path.Join(path.Dir(filename), menuPicName)
	pic , err := loadPicture(menupicture)
	if err != nil{
		log.Fatal(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())

	mat := pixel.IM
	mat = mat.Moved(shared.Win.Bounds().Center())
	imageMatrix := mat.ScaledXY(shared.Win.Bounds().Center(), pixel.V(0.7, 0.7))

	//clear background
	shared.Win.Clear(colornames.Black)
	sprite.Draw(shared.Win, imageMatrix)

	//text
	basicTxt.Draw(shared.Win, pixel.IM.Scaled(basicTxt.Orig, 3))

	//menu loop
	var i int = 0
	for !shared.Win.JustPressed(pixelgl.KeyEnter) && !shared.Win.Closed() {

		if i == 0 {
			basicTxt.Color = colornames.Black
			basicTxt.Draw(shared.Win, pixel.IM.Scaled(basicTxt.Orig, 3))
			}else if i == 10{
			basicTxt.Color = colornames.White
			basicTxt.Draw(shared.Win, pixel.IM.Scaled(basicTxt.Orig, 3))
		}else if i == 20 {
			i = 0
		}
		i++
		time.Sleep(50 * time.Millisecond)
		shared.Win.Update()
	}

}

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
