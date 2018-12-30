package foundation

import "github.com/rs/xid"

//ElementID は全てのエレメントを一意に認識する文字列です
type ElementID xid.ID

//Element はゲームを構成する様々な要素を表すインターフェイスです
type Element interface {
	Observer
	Notifier
}

//TheElement はゲームを構成する様々な要素です。
type TheElement struct {
	ID   ElementID // Instansを一意に表すID
	Name string    // エレメント名
}
