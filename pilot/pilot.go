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
	"fmt"
	"math"
	"os"
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
	LogicalSize      // 初期ウィンドウまたは論理座標
	running     bool // 処理中か否か
	mtxRunning  sync.Mutex
}

/*
SDLのメインスレッドを有効にします
*/
func MainThread(f func()) {
	sdl.Main(f)
}

/*
Pilotを初期化します
*/
func NewPilot(title string, numEffectChannel int, width, height int32, flags uint32) (*Pilot, error) {
	var err error
	var window *sdl.Window
	p := Pilot{}
	p.LogicalSize = LogicalSize{width, height}
	sdl.Do(func() {
		window, err = sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			height, width, flags)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
			return
		}
	})
	p.Controller = NewController(window, p.LogicalSize)
	if p.Renderer, err = NewRenderer(window, p.LogicalSize); err != nil {
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

	// 音声の受信ループ
	go func(ch <-chan data.Conduct) {
		for cndct := range ch {
			pilot.Orchestra.Play(cndct)
		}
	}(soundCh)
	// メインループ
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
		sdl.Do(func() {
			pilot.Renderer.Window.Destroy()
		})
	}(eventCh)

	return nil
}

/*
論理座標サイズを表す構造体
*/
type LogicalSize struct {
	Width  int32
	Height int32
}

// 論理移動量を与えてスクリーン移動量を計算する
func (ls LogicalSize) screen2logicalMove(win *sdl.Window, mx, my int32) (x, y float64) {
	sw, sh := win.GetSize()
	x = (float64(mx) / float64(sw)) * float64(ls.Width)
	y = (float64(my) / float64(sh)) * float64(ls.Height)
	return x, y
}

// 論理座標を与えてスクリーン座標を計算する
func (ls LogicalSize) screen2logical(win *sdl.Window, lx, ly int32) (x, y int32) {
	sw, sh := win.GetSize()
	x = int32(math.Floor((float64(lx) / float64(sw)) * float64(ls.Width)))
	y = int32(math.Floor((float64(ly) / float64(sh)) * float64(ls.Height)))
	return x, y
}
