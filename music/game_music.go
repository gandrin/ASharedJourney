package music

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

const musicMTfileName string = "/music/MainThemeMiroir.mp3"

type musicStreamers struct {
	//list of loaded musics ( streamer )
	mainTheamStreamer beep.Streamer
	backgroundMusic   *beep.Buffer
	gameEffects       map[soundEffect]*beep.Buffer
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
	go m.playMainTheme()
	m.loadEffects()
}

func (m *musicStreamers) loadMainTheam() beep.Format {
	var format beep.Format
	m.mainTheamStreamer, format = getStream(musicMTfileName)
	m.backgroundMusic = beep.NewBuffer(format)
	m.backgroundMusic.Append(m.mainTheamStreamer)
	m.streamControl.Paused = false
	return format
}

func (m *musicStreamers) playMainTheme() {
	var streamer = m.backgroundMusic.Streamer(0, m.backgroundMusic.Len())
	loopedaudio := beep.Loop(5, streamer)
	speaker.Play(beep.Seq(loopedaudio))
}

func getfilename(fileName string) string {
	rootDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal("error loading called")
	}
	//log.Print("file ", rootDirectory+fileName)
	return rootDirectory + fileName
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
