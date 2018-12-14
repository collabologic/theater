package data

import (
	"github.com/veandco/go-sdl2/sdl"
)

// スプライトのデータ型です。
type Sprite struct {
	Id         SpriteIdentifier // 個別のスプライトインスタンスの固有ID
	DistRect   Rect             // 書き出し先の矩形
	ImageType  ImageType        // リソースから読み込みからイメージ自体をもらうか
	SrcImageId ImageIdentifier  // スプライト画像のリソースID
	Image      SpriteImage      // View側で生成したイメージ
	SrcRect    Rect             // 読み出し元矩形
	Rotate     Rotate           // 回転
	Flip       Flip             // 裏返し
}

// スプライトID
type SpriteIdentifier int32

// 矩形
type Rect struct {
	Left   int32
	Top    int32
	Widt   int32
	Height int32
}

// 描画レイヤーの識別子の列挙型（Viewで定義しておくのが望ましい）
type LayerIdentifier int8

// スプライト画像の識別子（Viewで定義しておくのが望ましい）
type ImageIdentifier int8

// イメージ定義（スプライトマップファイルとその中の矩形)
type SpriteImage struct {
	SpriteTable *sdl.Texture // スプライトテーブル（スプライト並べた画像）
	Rect        sdl.Rect     // スプライトテーブル上の矩形
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
