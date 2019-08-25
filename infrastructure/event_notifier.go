package infrastructure

import (
	"github.com/collabologic/theater/data"
	"github.com/collabologic/theater/foundation"
)

//EventNotifier はpilotから受け取った入力情報をメッセージに変換して送信します
type EventNotifier struct {
	foundation.TheElement
	foundation.TheNotifier
	Mq    *foundation.MsgQueue
	evtch <-chan data.Event
}

//NewEventNotifier はEventNotifierを生成します
func NewEventNotifier(mq *foundation.MsgQueue, evtch <-chan data.Event) *EventNotifier {
	en := EventNotifier{
		Mq:    mq,
		evtch: evtch,
	}
	return &en
}

//Loop は、イベント受信チャンネルを待ち受け、イベントを受信するたびにメッセージを送信します
func (en *EventNotifier) Loop() {
	go func(evtch <-chan data.Event) {
		for evt := range evtch {
			msg := foundation.Msg{}
			msg.ID = EventCode2MsgID(evt.Code)
			msg.Sender = en
			msg.SenderID = en.ID
			if evt.Code >= data.MouseLeftDown && evt.Code <= data.MouseWheelDown {
				msg.Params = evt.Mouse
			}
			if evt.Code >= data.KeyPressOff && evt.Code <= data.KeyPressRepeat {
				msg.Params = evt.Keyboard
			}
			en.Mq.SndMsg(msg)
		}
	}(en.evtch)
}
