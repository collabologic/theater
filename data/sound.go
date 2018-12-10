package data

// 音声再生の指示データ
type Sound struct {
	PlayType PlayType
}

// 演奏方法の列挙型
type PlayType int

// PlayType型の値
const (
	BGM    int = iota // PlayType型のBGM
	EFFECT            // PlayType型の効果音
)
