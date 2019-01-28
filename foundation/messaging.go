package foundation

import (
	"fmt"
	"sync"
)

//Observer はメッセージの監視を行うインターフェイスです。
type Observer interface {
	Observe(mq *MsgQueue, msgID MsgIdentifier, cond MsgSelector, h MsgHandler)
}

//TheObserver は埋込で用いられる標準的なObserverです
type TheObserver struct{}

//Observe はハンドラメソッドをキューに登録します
func (to TheObserver) Observe(mq *MsgQueue, msgID MsgIdentifier, ms MsgSelector, h MsgHandler) {
	mq.RegistHandler(msgID, ms, h)
}

// -------------------------------------------------------------------------------

//Notifier はメッセージの送信を行うインターフェイスです。
type Notifier interface {
	Notify(msg Msg)
}

//TheNotifier は埋込で用いられる標準的なNotifierです
type TheNotifier struct {
	Mq *MsgQueue
}

//Notify はMsgQuereにメッセージを送信します
func (tn TheNotifier) Notify(msg Msg) {
	tn.Mq.SndMsg(msg)
}

// -------------------------------------------------------------------------------

//MsgQueue はメッセージの仲介を行います
type MsgQueue struct {
	qmtx     sync.Mutex                             // キュー更新のロック制御
	queue    []Msg                                  // メッセージ配列
	handlers map[MsgIdentifier][]msgSelectorHandler // 監視メッセージテーブル
	Running  bool                                   // 実行中フラグ
}

//NewMsgQueue はメッセージCueを生成します
func NewMsgQueue() MsgQueue {
	return MsgQueue{
		queue:    make([]Msg, 0),
		handlers: make(map[MsgIdentifier][]msgSelectorHandler),
		Running:  true,
	}
}

//elementID はRegisterHanler用MsgSelector関数としてElementIDでのチェックを行う関数ポインタを返却します
func elementID(eid ElementID) MsgSelector {
	return func(m *Msg) bool {
		return m.SenderID == eid
	}
}

//RegistHandler はメッセージ送信時の通知先に関する情報を登録します
func (mq *MsgQueue) RegistHandler(ID MsgIdentifier, ms MsgSelector, mh MsgHandler) {
	msh := msgSelectorHandler{ms, mh}
	mq.handlers[ID] = append(mq.handlers[ID], msh)
}

//SndMsg はメッセージをキューに登録します。
func (mq *MsgQueue) SndMsg(m Msg) {
	mq.qmtx.Lock()
	mq.queue = append(mq.queue, m)
	mq.qmtx.Unlock()
}

//Loop は、queueの内容を確認し、MessageIDとmsgSelectorに適合するハンドラメソッドを呼び出します。
func (mq *MsgQueue) Loop() {
	for mq.Running {
		if len(mq.queue) > 0 {
			mq.qmtx.Lock()
			m := mq.queue[0]
			mq.queue = mq.queue[1:]
			mq.qmtx.Unlock()

			hs, ok := mq.handlers[m.ID]
			if !ok {
				panic(fmt.Sprintf("MessageQueue received unregisterd MessageID:%s", m.ID))
			}
			for _, h := range hs {
				go func(m *Msg, h *msgSelectorHandler) {
					if h.Selector == nil || h.Selector(m) || h.Handler != nil {
						h.Handler(m)
					}
				}(&m, &h)
			}
		}
	}
}

// -------------------------------------------------------------------------------

// msgSelectorHandler は条件関数とハンドラー関数の組み合わせです
type msgSelectorHandler struct {
	Selector MsgSelector
	Handler  MsgHandler
}

// -------------------------------------------------------------------------------

//Msg はエレメント間の通信データを表します。
type Msg struct {
	ID       MsgIdentifier
	SenderID ElementID
	Sender   Notifier
	Params   MsgParams
}

//MsgParams メッセージのパラメータを保存します（キャストして取り扱うこと）
type MsgParams interface{}

// -------------------------------------------------------------------------------

//MsgIdentifier はメッセージを一意に表す識別子です
type MsgIdentifier string

//MsgSelector メッセージを絞り込むための条件関数へのポインタです
type MsgSelector func(m *Msg) bool

//MsgHandler は、メッセージ受信次に実行するハンドラー関数のポインタです
type MsgHandler func(m *Msg)
