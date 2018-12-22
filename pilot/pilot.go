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

import (
	"sync"

	"github.com/collabologic/theater/data"
	"github.com/veandco/go-sdl2/sdl"
)

/*
Pilot構造体はPilotの本体です。

controller, renderer, orchestraを初期先、それぞれgoroutineとしてプロセスを走らせます。
*/
type Pilot struct {
	*Controller
	*Renderer
	*Orchestra
	running    bool // 処理中か否か
	mtxRunning sync.Mutex
}

/*
Pilotを初期化します
*/
func NewPilot(win *sdl.Window, numEffectChannel int) (*Pilot, error) {
	var err error
	p := Pilot{}
	p.Controller = NewController()
	if p.Renderer, err = NewRenderer(win); err != nil {
		return nil, err
	}
	if p.Orchestra, err = NewOrchestra(numEffectChannel); err != nil {
		return nil, err
	}
	p.Renderer.SdlRenderer.SetDrawColor(0, 0, 0, 255)
	p.Renderer.SdlRenderer.Clear()
	p.Renderer.SdlRenderer.Present()
	return &p, nil
}

/*
RunはPilotを稼働させます。

具体的には、controler, renderer, orchestraを生成し、それぞれの送受信ループをgoroutinとして走らせます。
*/
func (pilot *Pilot) Run(eventCh chan<- data.Event, spriteCh <-chan data.Sprite, soundCh <-chan data.Conduct) error {

	// スプライトの受信ループ
	go func(ch <-chan data.Sprite) {
		for sprt := range ch {
			pilot.Renderer.AddSpriteForLayer(sprt)
		}
	}(spriteCh)

	// 画面描画ループ
	/*go func() {
		sdl.Do(func() {
			for {
				if !pilot.running {
					break
				}
				if err := pilot.Renderer.DrawLayers(); err != nil {
					panic(err)
				}
			}
		})
	}()*/

	// 音声の受信ループ
	go func(ch <-chan data.Conduct) {
		for cndct := range ch {
			pilot.Orchestra.Play(cndct)
		}
	}(soundCh)
	// 入力イベントの送信ループ
	go func(evtch chan<- data.Event) {
		defer close(evtch)

		pilot.mtxRunning.Lock()
		pilot.running = true
		pilot.mtxRunning.Unlock()
		for pilot.running {
			sdl.Do(func() {
				if !pilot.running {
					return
				}

				running, res, err := pilot.Controller.ReceiveEvent()
				if err != nil {
					panic(err)
				}
				if err := pilot.Renderer.DrawLayers(); err != nil {
					panic(err)
				}
				if !running {
					pilot.mtxRunning.Lock()
					pilot.running = false
					pilot.mtxRunning.Unlock()
				} else if res.Code != data.NoEvent {
					evtch <- res
				}

			})
		}

	}(eventCh)

	return nil
}
