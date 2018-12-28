package pilot

import (
	"errors"
	"sort"

	"github.com/collabologic/theater/data"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

/*
Rendererは、スプライトデータをチャンネルから受け取って画面に書き出します。
*/
type Renderer struct {
	// 論理画面サイズ
	LogicalSize
	// レイヤーごとのスプライト
	Layers map[data.LayerIdentifier]map[data.SpriteIdentifier]*data.Sprite
	// レイヤーの画像
	LayerTextures map[data.LayerIdentifier]*sdl.Texture
	// レイヤーの更新フラグ
	LayerUpdated map[data.LayerIdentifier]bool
	// スプライトイメージ（個別の画像）
	SpriteImages map[data.ImageIdentifier]*data.SpriteImage
	// ウィンドウ
	Window *sdl.Window
	// ウィンドウに対するSDLレンダラー
	SdlRenderer *sdl.Renderer
}

/*
NewRendererはRendererを初期化します。指定した Windowにsdl.Rendererを追加します。
}*/
func NewRenderer(win *sdl.Window, ls LogicalSize) (*Renderer, error) {
	var err error
	renderer := Renderer{}
	renderer.LogicalSize = ls
	renderer.Window = win
	renderer.SdlRenderer, err = sdl.CreateRenderer(
		win, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_TARGETTEXTURE)

	//renderer.SdlRenderer, err = win.GetRenderer()
	if err != nil {
		return &renderer, err
	}
	renderer.Layers = make(map[data.LayerIdentifier]map[data.SpriteIdentifier]*data.Sprite)
	renderer.LayerTextures = make(map[data.LayerIdentifier]*sdl.Texture)
	renderer.LayerUpdated = make(map[data.LayerIdentifier]bool)
	renderer.SpriteImages = make(map[data.ImageIdentifier]*data.SpriteImage)
	return &renderer, nil
}

/*
AddRayerはRendererに指定したIDで描画レイヤーを追加します。
*/
func (renderer *Renderer) AddLayer(identifier data.LayerIdentifier) error {
	var err error
	renderer.Layers[identifier] = make(map[data.SpriteIdentifier]*data.Sprite)
	renderer.LayerTextures[identifier], err = renderer.SdlRenderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		renderer.LogicalSize.Width,
		renderer.LogicalSize.Height,
	)
	if err != nil {
		return err
	}
	renderer.LayerUpdated[identifier] = false
	return nil
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
	for y := 0; y < len(identifiers)/horizontal; y++ {
		for x := 0; x < horizontal; x++ {
			r := data.Rect{
				int32(x) * height,
				int32(y) * width,
				width,
				height,
			}
			img := data.SpriteImage{tx, r}
			renderer.SpriteImages[identifiers[id]] = &img
			id += 1
		}
	}
	return nil
}

/*
addSpriteForLayerはレイヤーにスプライトを追加・または更新します
*/
func (renderer *Renderer) AddSpriteForLayer(sprite data.Sprite) {
	layerID := sprite.LayerID
	renderer.Layers[layerID][sprite.Id] = &sprite
	renderer.LayerUpdated[layerID] = true
}

/*
drawSpriteForLayerはレイヤーを書き出します
*/
func (renderer *Renderer) drawSpriteForLayer(layerID data.LayerIdentifier) error {
	var err error
	// 対象のレイヤーテクスチャをクリアーする
	if err = renderer.LayerTextures[layerID].Destroy(); err != nil {
		return err
	}
	w, h := renderer.Window.GetSize()
	if renderer.LayerTextures[layerID], err = renderer.SdlRenderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		w,
		h,
	); err != nil {
		return err
	}
	// 対象レイヤーを編集ターゲットにする
	texture := renderer.LayerTextures[layerID]
	if err := renderer.SdlRenderer.SetRenderTarget(texture); err != nil {
		return err
	}
	if err := renderer.SdlRenderer.Clear(); err != nil {
		return err
	}
	spriteArray := getSpriteArraySortedPriority(renderer.Layers[layerID])
	// 実際に書き出す
	for _, sprite := range spriteArray {
		if sprite == nil {
			continue
		}
		si, ok := renderer.SpriteImages[sprite.SrcImageID]
		if !ok {
			return errors.New("Unknown Sprite Image.")
		}
		point := sdl.Point{
			X: sprite.Rotate.CenterX,
			Y: sprite.Rotate.CenterY,
		}
		var flip sdl.RendererFlip
		switch sprite.Flip {
		case data.NoFlip:
			flip = sdl.FLIP_NONE
		case data.Horizontal:
			flip = sdl.FLIP_HORIZONTAL
		case data.Vertical:
			flip = sdl.FLIP_VERTICAL
		}
		renderer.SdlRenderer.CopyEx(
			si.SpriteTable,
			si.Rect.ToSdlRect(),
			sprite.DistRect.ToSdlRect(),
			sprite.Rotate.Angle,
			&point,
			flip,
		)
	}
	return nil
}

/*
DrawLayersはレイヤーを順番に書き出します。
*/
func (renderer *Renderer) DrawLayers() error {
	ids := renderer.getLayerIDs()
	for _, id := range ids {
		// 更新ずみの場合のみ、スプライト書き出し処理を行う
		if _, ok := renderer.LayerUpdated[id]; ok {
			if err := renderer.drawSpriteForLayer(id); err != nil {
				return err
			}
			renderer.LayerUpdated[id] = false
		} else {
			return errors.New("Access By Invalid Layer Id")
		}
		// 画面に書き出す
		if err := renderer.SdlRenderer.SetRenderTarget(nil); err != nil {
			return err
		}
		// 書き出し元（論理サイズ）
		src := sdl.Rect{X: 0, Y: 0, W: renderer.Width, H: renderer.Height}
		// 書き出し先（スクリーンサイズ）
		ww, wh := renderer.Window.GetSize()
		dst := sdl.Rect{X: 0, Y: 0, W: ww, H: wh}
		renderer.SdlRenderer.Copy(renderer.LayerTextures[id], &src, &dst)

	}
	renderer.SdlRenderer.Present()
	renderer.Window.UpdateSurface()
	return nil
}

/*
getLayerIDsは、指定レイヤーの全てのレイヤーIDを昇順にソートして返却します
*/
func (renderer Renderer) getLayerIDs() []data.LayerIdentifier {
	ids := make([]data.LayerIdentifier, len(renderer.LayerUpdated))
	for key, _ := range renderer.LayerUpdated {
		ids = append(ids, key)
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	return ids
}

// getSpriteArraySortedPriorityは、SpriteIDをキーにしたSpriteのmapから、priority順の配列に変換
func getSpriteArraySortedPriority(
	spriteMap map[data.SpriteIdentifier]*data.Sprite,
) []*data.Sprite {
	spriteArray := make([]*data.Sprite, len(spriteMap))
	for _, tx := range spriteMap {
		spriteArray = append(spriteArray, tx)
	}
	//fmt.Println(spriteArray)
	// LayerのSprite配列をPriority順にソート
	/*
		sort.Slice(spriteArray, func(i, j int) bool {
			return spriteArray[i].Priority < spriteArray[j].Priority
		})
	*/
	return spriteArray
}
