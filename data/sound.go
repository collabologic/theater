package data

import "github.com/veandco/go-sdl2/mix"

// 音声の再生指示
type Conduct struct {
	ID       SoundIdentifier
	Repeat   int
	FadeTime int
	Volume   int
}

// 音声データ
type Sound struct {
	ID       SoundIdentifier
	PlayType PlayType
	Music    *mix.Music
	Effect   *mix.Chunk
}

// サウンドIDの列挙型（値はApp側で設定する）
type SoundIdentifier int32

// 演奏方法の列挙型
type PlayType int

// PlayType型の値
const (
	BGM    PlayType = iota // PlayType型のBGM
	EFFECT                 // PlayType型の効果音
)

//
