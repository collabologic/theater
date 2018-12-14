package pilot

import (
	"github.com/collabologic/theater/data"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

/*
Rendererは、スプライトデータをチャンネルから受け取って画面に書き出します。
*/
type Renderer struct {
	// ウィンドウへの参照
	window *sdl.Window
	// レイヤーごとのスプライト
	Laysers map[data.LayerIdentifier]map[data.SpriteIdentifier]*data.Sprite
	// スプライトイメージ（個別の画像）
	SpriteImages map[data.ImageIdentifier]*data.SpriteImage
	// ウィンドウ
	Window sdl.Window
	// ウィンドウに対するSDLレンダラー
	SdlRenderer *sdl.Renderer
}

/*
InitはRendererを初期化します。指定した Windowにsdl.Rendererを追加します。
}*/
func (renderer *Renderer) Init(win *sdl.Window) error {
	var err error
	renderer.window = win
	renderer.SdlRenderer, err = sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return err
	}
	renderer.Laysers = make(map[data.LayerIdentifier]map[data.SpriteIdentifier]*data.Sprite)
	return nil
}

/*
AddRayerはRendererに指定したIDで描画レイヤーを追加します。
*/
func (renderer *Renderer) AddLayer(identifier data.LayerIdentifier) error {
	var err error
	renderer.Laysers[identifier] = make(map[data.SpriteIdentifier]*data.Sprite)
	return err
}

/*
AddSpriteImagesは指定したファイルをwidth, heightの大きさで裁断してSpriteととしてRendererに追加します。
指定する画像は、width,heightの大きさで横にhorizontal分だけ並んでいる想定です。（つまりhorizontalの個数で折り返します）
identitiesに指定した識別子の件数分だけ読み込みます。
*/
func (renderer *Renderer) AddSpriteImages(
	filename string,
	width int32,
	height int32,
	horizontal int,
	identifiers []data.ImageIdentifier,
) error {
	image, err := img.Load(filename)
	if err != nil {
		return err
	}
	tx, err := renderer.SdlRenderer.CreateTextureFromSurface(image)
	if err != nil {
		image.Free()
		return err
	}

	id := 0 // 作業中のidentifier
	for y := 0; y < horizontal/len(identifiers); y++ {
		for x := 0; x < horizontal; x++ {
			r := sdl.Rect{
				int32(x) * height,
				int32(y) * width,
				width,
				height,
			}
			renderer.SpriteImages[identifiers[id]] = &(data.SpriteImage{tx, r})
			id += 1
		}
	}
	return nil
}

/*
レイヤーにスプライトを追加・または更新する
*/

/*
レイヤーを順番に書き出す
*/

/*
Runは、Spriteデータを受け取って画面に画像を出力します。
*/
func (renderer *Renderer) Run(sprite <-chan data.Sprite) {

}
