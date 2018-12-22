package pilot

import (
	"errors"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/collabologic/theater/data"
	mix "github.com/veandco/go-sdl2/mix"
)

/*
Orchestraは、チャンネルから支持された音声ファイルを再生します
ループする音楽と単なる効果音は別に扱う必要がありmす。
*/
type Orchestra struct {
	Sounds   map[data.SoundIdentifier]data.Sound
	Channels []bool
}

/* Orchestraを初期化します */
func NewOrchestra(effectChanel int) (*Orchestra, error) {
	orchestra := Orchestra{}
	orchestra.Sounds = make(map[data.SoundIdentifier]data.Sound)
	if err := mix.Init(mix.INIT_OGG); err != nil {
		return &Orchestra{}, err
	}
	err := mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE)
	if err != nil {
		return &Orchestra{}, err
	}
	ac := mix.AllocateChannels(effectChanel)
	orchestra.Channels = make([]bool, ac)
	mix.ChannelFinished(func(id int) {
		orchestra.Channels[id] = false
	})
	return &orchestra, nil
}

/*
BGMファイルを追加します。
*/
func (orchestra *Orchestra) AddBGM(id data.SoundIdentifier, filename string) error {
	var m *mix.Music
	var err error
	if m, err = mix.LoadMUS(filename); err != nil {
		return err
	}
	s := data.Sound{
		ID:       id,
		PlayType: data.BGM,
		Music:    m,
	}
	orchestra.Sounds[id] = s
	return nil
}

/*
効果音ファイルを追加します。
*/
func (orchestra *Orchestra) AddEffect(id data.SoundIdentifier, filename string) error {
	var e *mix.Chunk
	var err error
	if e, err = mix.LoadWAV(filename); err != nil {
		return err
	}
	s := data.Sound{
		ID:       id,
		PlayType: data.EFFECT,
		Effect:   e,
	}
	orchestra.Sounds[id] = s
	return nil
}

/*
音声の再生を指示します。
*/
func (orchestra Orchestra) Play(conduct data.Conduct) error {
	var (
		s  data.Sound
		ok bool
	)
	if s, ok = orchestra.Sounds[conduct.ID]; !ok {
		return errors.New(fmt.Sprintf("No Sound file.:%d", conduct.ID))
	}
	switch s.PlayType {
	case data.BGM:
		orchestra.bgm(conduct.ID, conduct.Repeat, conduct.FadeTime)
	case data.EFFECT:
		orchestra.effect(conduct.ID, conduct.Volume, conduct.Repeat, conduct.FadeTime)
	}
	return nil
}

/*
BGMとして設定された音声を再生します
*/
func (orchestra Orchestra) bgm(id data.SoundIdentifier, repeat int, fadetime int) error {
	s, ok := orchestra.Sounds[id]
	if !ok {
		return errors.New(fmt.Sprintf("No music file:%s", id))
	}
	music := s.Music
	if fadetime == 0 {
		if err := music.Play(repeat); err != nil {
			return err
		}
	} else {
		if err := music.FadeIn(repeat, fadetime); err != nil {
			return err
		}
	}
	return nil
}

/*
効果音として設定された音声を再生します
*/
func (orchestra Orchestra) effect(id data.SoundIdentifier, volume, repeat, fadetime int) error {
	s, ok := orchestra.Sounds[id]
	effect := s.Effect
	if !ok {
		return errors.New(fmt.Sprintf("No music:%s", id))
	}
	effect.Volume((sdl.MIX_MAXVOLUME / 10) * volume)
	c, err := orchestra.getFreeChan()
	if err != nil {
		return err
	}
	if fadetime == 0 {
		if _, err := effect.Play(c, repeat); err != nil {
			return err
		}
	} else {
		if _, err := effect.FadeIn(c, repeat, fadetime); err != nil {
			return err
		}
	}
	return nil
}

/*
再生可能なチャンネルを取得して再生中の状態にします
*/
func (orchestra *Orchestra) getFreeChan() (int, error) {
	var fc int
	var ok bool = false
	for c, v := range orchestra.Channels {
		if !v {
			fc = c
			orchestra.Channels[c] = true
			ok = true
			break
		}
	}
	if !ok {
		return 0, errors.New("No playable Channnel")
	}
	return fc, nil
}
