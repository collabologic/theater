package pilot

import "github.com/collabologic/theater/data"

/*
Orchestraは、チャンネルから支持された音声ファイルを再生します
ループする音楽と単なる効果音は別に扱う必要がありmす。
*/
type Orchestra struct {
	// サウンドのファイルを保持します
}

/*
Runは音声再生をチャンネルから待ち受けるメインループです。
Sound構造体をチャンネルから受け取り、SDL_mixerを用いて再生します。
*/
func (orchestra Orchestra) Run(sound <-chan data.Sound) {

}
