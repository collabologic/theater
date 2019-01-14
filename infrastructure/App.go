package infrastructure

import (
	"time"

	"github.com/collabologic/theater/foundation"
	"github.com/collabologic/theater/pilot"
)

//App はTheaterアプリケーションのインターフェイスです
type App interface {
	foundation.Observer
	foundation.Notifier
	Init()
	Run()
}

//TheApp は標準的なAppインターフェイスの実装です
type TheApp struct {
	foundation.TheObserver
	foundation.TheNotifier

	scenes         map[SceneID]*Scene
	currentSceneID SceneID
	//spaces: map[SpaceID]Space
	mq              foundation.MsgQueue
	componentPulsar foundation.Pulsar
	spacePulsar     foundation.Pulsar
	pilot           *pilot.Pilot
}

//NewTheApp は新規に新しいAppを作成します
func NewTheApp(mq *foundation.MsgQueue, cd time.Duration, sd time.Duration) TheApp {
	//TODO: SceneとSpaceの定義
	return TheApp{
		scenes: make(map[SceneID]*Scene, 0),
		//spaces: make([]Spaces),
		mq:              foundation.NewMsgQueue(),
		componentPulsar: foundation.NewPulsar(cd),
		spacePulsar:     foundation.NewPulsar(sd),
	}
}

//AppendScene はAppにSceneを追加します
func (ta *TheApp) AppendScene(sid SceneID, scene *Scene) {
	ta.scenes[sid] = scene
}

//Init はAppを初期化します。embeded先で定義し直してください
func (ta *TheApp) Init() {}

//Run はゲームを走らせます。embeded先で定義し直してください
func (ta *TheApp) Run() {}
