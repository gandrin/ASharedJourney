package music

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

const musicMTfileName string = "/MainThemeMiroir.mp3"

type musicStreamers struct {
	//list of loaded musics ( streamer )
	mainTheamStreamer beep.Streamer
	gameEffects       map[soundEffect]beep.Streamer
	streamControl     beep.Ctrl
}

var Music musicStreamers

//called when package is called into scope the first time
func (m *musicStreamers) Start() {
	var format beep.Format

	format = m.loadMainTheam()
	err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}
	m.streamControl.Streamer = m.mainTheamStreamer
	m.loadEffects()
}

func (m *musicStreamers) loadMainTheam() beep.Format {
	var format beep.Format
	m.mainTheamStreamer, format = getStream(musicMTfileName)
	return format
}

func (m *musicStreamers) Play() {
	m.streamControl.Paused = false
	go m.playMainTheme()
}

func (m *musicStreamers) playMainTheme() {

	log.Print("Starting music")
	speaker.Play(m.streamControl.Streamer)

	log.Print("Music finished")

}

func getfilename(fileName string) string {
	_, rootName, _, _ := runtime.Caller(1)
	f, err := os.OpenFile("/Users/Pierpo/ASharedJourneyLogs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("rootName " + rootName)
	if !ok {
		log.Fatal("error loading called")
	}
	path := path.Join(path.Dir(rootName), fileName)
	return path
}

func getStream(filename string) (beep.StreamCloser, beep.Format) {

	absfilepath := getfilename(filename)
	var newStreamer beep.StreamCloser
	var format beep.Format
	var err error
	file, err := os.Open(absfilepath)
	if err != nil {
		log.Fatal("Music file  ", err)
	}
	ext := filepath.Ext(filename)
	if ext == ".mp3" {
		newStreamer, format, err = mp3.Decode(file)
	} else if ext == ".wav" {
		newStreamer, format, err = wav.Decode(file)
	}

	if err != nil {
		log.Fatal("Decorer error on file ", absfilepath, " ", err)
	}
	return newStreamer, format
}
