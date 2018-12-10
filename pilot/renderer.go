package pilot

import "github.com/collabologic/theater/data"

/*
Rendererは、スプライトデータをチャンネルから受け取って画面に書き出します。
*/
type Renderer struct {
	// SDLのレイヤー情報
	// 画像ファイル
}

/*
Runは、Spriteデータを受け取って画面に画像を出力します。
*/
func (renderer *Renderer) Run(sprite <-chan data.Sprite) {

}
