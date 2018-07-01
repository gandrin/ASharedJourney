package menu

import (
	"fmt"
	"image"
	"log"
	"os"
	"path"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/gandrin/ASharedJourney/music"
	"github.com/gandrin/ASharedJourney/shared"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const menuText string = "Press ENTER to PLAY"
const menuTextPosX float64 = 200
const menuTextPosY float64 = 100
const menuPicName string = "menu.png"

//draw menu to screen while player while player hasn't pressed enter
func Menu() {

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(menuTextPosX, menuTextPosY), basicAtlas)
	basicTxt.Color = colornames.White
	fmt.Fprintln(basicTxt, menuText)

	//get picture
	menupicture := path.Join(".", "assets", menuPicName)
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

	//menu loop
	for !shared.Win.JustPressed(pixelgl.KeyEnter) && !shared.Win.Closed() {
		time.Sleep(50 * time.Millisecond)
		shared.Win.Update()
	}
	music.Music.PlayEffect(music.SOUND_EFFECT_START_GAME)

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
