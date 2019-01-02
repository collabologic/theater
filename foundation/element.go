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

//Displayer は画面に表示される要素を表すインターフェイスです
type Displayer interface {
	Display()
}

//Behavior は時間経過で自分の状態を変化させる要素を表すインターフェイスです
type Behavior interface {
	Behave()
}

//Component は画面に表示され、メッセージの送受信ができるインターフェイスです
type Component interface {
	Element
	Displayer
}

//Actor はメッセージの送受信と時間経過で状態を変化させる要素を表すインターフェイスです
type Actor interface {
	Element
	Behavior
}
