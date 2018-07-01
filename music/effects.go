package music

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"log"
)

type soundEffect string

const (
	SOUND_EFFECT_START_GAME soundEffect = "/music/MenuEffect.wav"

	SOUND_EFFECT_WIN_GAME soundEffect = "/music/win2.mp3"

	SOUND_EFFECT_LOSE_GAME soundEffect = "/music/lose2.mp3"
)

func (m *musicStreamers) PlayEffect(effectType soundEffect){

	go func(m *musicStreamers,effectType soundEffect) {
		//play the sound effect
		es, ok := m.gameEffects[effectType]
		if ok {
			speaker.Lock()
			m.streamControl.Paused = true
			speaker.Unlock()
			log.Print("Creating new stream entry")
			done := make(chan struct{})
			//effect exists -> play
			speaker.Play(beep.Seq(
				es,
				beep.Callback(func() {
					close(done)
				}),
			))
			log.Print("finished playing effect ")
			speaker.Lock()
			m.streamControl.Paused = false
			speaker.Unlock()
			<-done

			log.Print("stream of sound finished")
		}
	}(m,effectType )





}

func (m *musicStreamers) loadEffects(){
	m.gameEffects = make(map[soundEffect]beep.Streamer,0)
	m.gameEffects[SOUND_EFFECT_START_GAME], _ = getStream(string(SOUND_EFFECT_START_GAME))
	m.gameEffects[SOUND_EFFECT_LOSE_GAME], _ = getStream(string(SOUND_EFFECT_WIN_GAME))
	m.gameEffects[SOUND_EFFECT_LOSE_GAME], _ = getStream(string(SOUND_EFFECT_LOSE_GAME))
}