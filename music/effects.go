package music

import (
	"log"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

type soundEffect string

const (
	SOUND_EFFECT_START_GAME soundEffect = "/music/MenuEffect.wav"

	SOUND_EFFECT_WIN_GAME soundEffect = "/music/win2.mp3"

	SOUND_EFFECT_LOSE_GAME soundEffect = "/music/lose2.mp3"

	SOUND_EFFECT_WATER soundEffect = "/music/Acid_Bubble.mp3"
)

func (m *musicStreamers) PlayEffect(effectType soundEffect) {

	go func(m *musicStreamers, effectType soundEffect) {
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
				,
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
	}(m, effectType)

}

func (m *musicStreamers) loadEffects() {
	m.gameEffects = make(map[soundEffect]*beep.Buffer, 0)
	//make new buffers and add to buffer
	stream1,format1 := getStream(string(SOUND_EFFECT_START_GAME))
	m.gameEffects[SOUND_EFFECT_START_GAME]= beep.NewBuffer(format1)
	m.gameEffects[SOUND_EFFECT_START_GAME].Append(stream1)

	stream2,format2 := getStream(string(SOUND_EFFECT_WIN_GAME))
	m.gameEffects[SOUND_EFFECT_WIN_GAME]= beep.NewBuffer(format2)
	m.gameEffects[SOUND_EFFECT_WIN_GAME].Append(stream2)

	stream3,format3 := getStream(string(SOUND_EFFECT_WATER))
	m.gameEffects[SOUND_EFFECT_WATER]= beep.NewBuffer(format3)
	m.gameEffects[SOUND_EFFECT_WATER].Append(stream3)

}
