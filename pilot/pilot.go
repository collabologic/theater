/*
Pilotパッケージは、入出力を直接取り扱う処理を担当します。

Pilot

本体です。後述の三つのStructureをgoroutineとして起動します。

Controller

マウス、キーボード、ジョイパッドなどからユーザの入力を受け付けます。

Renderer

スプライト画像を合成し、画面に表示します。

Orchestra

音声ファイルを再生します
*/
package pilot

import "github.com/collabologic/theater/data"

/*
Pilot構造体はPilotの本体です。

controller, renderer, orchestraを初期先、それぞれgoroutineとしてプロセスを走らせます。
*/
type Pilot struct {
	Controller
	Renderer
	Orchestra
}

/*
RunはPilotを稼働させます。

具体的には、controler, renderer, orchestraを生成し、それぞれのRunをgoroutinとして走らせます。
*/
func (pilot *Pilot) Run(eventCh chan<- data.Event, spriteCh <-chan data.Sprite, soundCh <-chan data.Sound) error {
	return nil
}
