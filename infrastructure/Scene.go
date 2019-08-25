package infrastructure

import (
	"github.com/collabologic/theater/foundation"
)

//Scene は場面を定義するインターフェイスです
type Scene interface {
	foundation.Observer
	foundation.Notifier
	Init()
}

//SceneID はSceneを特定するID（文字列）です
type SceneID string

//TheScene は標準的なSceneの実装です
type TheScene struct {
	foundation.TheObserver
	foundation.TheNotifier

	mq       *foundation.MsgQueue                          // メッセージキューへのポインタ
	cp       *foundation.Pulsar                            // コンポーネントパルサーへのポインタ
	id       SceneID                                       // 自身のSceneID
	Contains map[foundation.ElementID]foundation.Component // Sceneに含まれるコンポーネント群
}

//Init はSceneを初期化します。embeded先で定義し直してください
func (scene *TheScene) Init() {}

//NewTheScene ばSceneオブジェクトを生成するコンストラクタです
func NewTheScene(mq *foundation.MsgQueue, cp *foundation.Pulsar, sid SceneID) Scene {
	th := TheScene{
		mq:       mq,
		cp:       cp,
		id:       sid,
		Contains: make(map[foundation.ElementID]foundation.Component),
	}
	return Scene(&th)
}
