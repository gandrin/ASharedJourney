package menu

import (
	"fmt"
	"image"
	"log"
	"os"
	"path"
	"runtime"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/gandrin/ASharedJourney/shared"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"github.com/faiface/pixel/pixelgl"
	"time"
	"github.com/gandrin/ASharedJourney/music"
)

const menuTextPosX float64 = 200
const menuTextPosY float64 = 100
const MenuPicName string = "tiles/menu.png"

//draw menu to screen while player while player hasn't pressed enter
func Menu(pictureName string, menuText string, blocking bool, exitSoundEffect music.SoundEffect) {

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(menuTextPosX, menuTextPosY), basicAtlas)
	basicTxt.Color = colornames.White
	fmt.Fprintln(basicTxt, menuText)

	//get picture
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("error loading called")
	}
	menupicture := path.Join(path.Dir(filename), pictureName)
	pic, err := loadPicture(menupicture)
	if err != nil {
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
	shared.Win.Update()

	//menu loop
	if blocking{
		for !shared.Win.JustPressed(pixelgl.KeyEnter) && !shared.Win.Closed() {
			time.Sleep(50 * time.Millisecond)
			shared.Win.Update()
		}
		music.Music.PlayEffect(exitSoundEffect)
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
