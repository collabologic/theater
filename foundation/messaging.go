package foundation

import "sync"

//Observer はメッセージの監視を行うインターフェイスです。
type Observer interface {
	Observe(mq *MsgQueue, msg Message, cond MsgSelector, h MsgHandler)
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
	Notify(mq *MsgQueue, msg Message)
}

//TheNotifier は埋込で用いられる標準的なNotifierです
type TheNotifier struct{}

//Notify はMsgQuereにメッセージを送信します
func (tn TheNotifier) Notify(mq *MsgQueue, msg Message) {
	mq.SndMsg(msg)
}

// -------------------------------------------------------------------------------

//MsgQueue はメッセージの仲介を行います
type MsgQueue struct {
	qmtx     sync.Mutex                           // キュー更新のロック制御
	queue    []Message                            // メッセージ配列
	handlers map[MsgIdentifier]msgSelectorHandler // 監視メッセージテーブル
	Running  bool                                 // 実行中フラグ
}

//RegistHandler はメッセージ送信時の通知先に関する情報を登録します
func (mq *MsgQueue) RegistHandler(ID MsgIdentifier, ms MsgSelector, mh MsgHandler) {
	msh := msgSelectorHandler{ms, mh}
	mq.handlers[ID] = msh
}

//SndMsgは メッセージをキューに登録します。
func (mq *MsgQueue) SndMsg(m Message) {
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

			h := mq.handlers[m.ID]

			if h.Selector(&m) {
				go func(m *Message) {
					h.Handler(m)
				}(&m)
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

//Message はエレメント間の通信データを表します。
type Message struct {
	ID       MsgIdentifier
	SenderID ElementID
	Sender   *Notifier
	Params   MsgParams
}

//MsgParams メッセージのパラメータを保存する
type MsgParams map[string]interface{}

//getInt は指定のパラメータをint64で返却することを試みます
func (mps MsgParams) getInt(key string) int64 {
	return mps[key].(int64)
}

//getFloat は指定のパラメータをfloat64で返却することを試みます
func (mps MsgParams) getFloat(key string) float64 {
	return mps[key].(float64)
}

//getString は指定のパラメータをstringで返却することを試みます
func (mps MsgParams) getString(key string) string {
	return mps[key].(string)
}

// -------------------------------------------------------------------------------

//MsgIdentifier はメッセージを一意に表す識別子です
type MsgIdentifier string

//MsgSelector メッセージを絞り込むための条件関数へのポインタです
type MsgSelector func(m *Message) bool

//MsgHandler は、メッセージ受信次に実行するハンドラー関数のポインタです
type MsgHandler func(m *Message)
