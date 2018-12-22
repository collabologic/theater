package data

import (
	"sync"

	"github.com/rs/xid"
	"github.com/veandco/go-sdl2/sdl"
)

// スプライトのデータ型です。
type Sprite struct {
	sync.Mutex
	LayerID    LayerIdentifier  // スプライトが書き出されるLayerのID
	Id         SpriteIdentifier // 個別のスプライトインスタンスの固有ID
	Updated    bool             // 更新フラグ。App側で管理するために使用
	DistRect   Rect             // 書き出し先の矩形
	Priority   int8             // 書き出しの優先度（大きいほど上にくる）
	SrcImageID ImageIdentifier  // スプライト画像のリソースID
	Rotate     Rotate           // 回転
	Flip       Flip             // 裏返し
}

func NewSprite(layerID LayerIdentifier) Sprite {
	guid := SpriteIdentifier(xid.New())
	sprite := Sprite{
		LayerID: layerID,
		Id:      guid,
	}
	return sprite
}

// スプライトID
type SpriteIdentifier xid.ID

// 矩形
type Rect struct {
	Left   int32
	Top    int32
	Width  int32
	Height int32
}

func (rect *Rect) ToSdlRect() *sdl.Rect {
	return &sdl.Rect{
		X: rect.Left,
		Y: rect.Top,
		H: rect.Height,
		W: rect.Width,
	}
}

// 描画レイヤーの識別子の列挙型（Viewで定義しておくのが望ましい）
type LayerIdentifier int8

// スプライト画像の識別子（Viewで定義しておくのが望ましい）
type ImageIdentifier int8

// 書き出しを行うメソッド
type DrawMethod *func() error

// イメージ定義（スプライトマップファイルとその中の矩形)
type SpriteImage struct {
	SpriteTable *sdl.Texture // スプライトテーブル（スプライト並べた画像）
	Rect        Rect         // スプライトテーブル上の矩形
}

// イメージの取得方法
type ImageType int8

// ImageType型の値
const (
	Resource  ImageType = iota // リソースリストから取得
	Generated                  // Viewが生成した画像を取得
)

// 回転のパラメータ
type Rotate struct {
	CenterX int32
	CenterY int32
	Angle   float64
}

// 裏返しの列挙型
type Flip int8

// Flip型の値
const (
	NoFlip     Flip = iota // 裏返しなし
	Horizontal             // 横方向
	Vertical               // 縦方向
)
