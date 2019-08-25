package infrastructure

import (
	"github.com/collabologic/theater/data"
	"github.com/collabologic/theater/foundation"
)

//入力系メッセージID
const (
	MsgIDMouseLeftDown      foundation.MsgIdentifier = "TheInputMouseLeftDown"
	MsgIDMouseLeftUp        foundation.MsgIdentifier = "TheInputMouseLeftUp"
	MsgIDMouseRightDown     foundation.MsgIdentifier = "TheInputMouseRightDown"
	MsgIDMouseRightUp       foundation.MsgIdentifier = "TheInputMouseRightUp"
	MsgIDMouseLeftDragging  foundation.MsgIdentifier = "TheInputMouseLeftDragging"
	MsgIDMouseRightDragging foundation.MsgIdentifier = "TheInputMouseRightDragging"
	MsgIDMouseLeftDrop      foundation.MsgIdentifier = "TheInputMouseLeftDrop"
	MsgIDMouseRightDrop     foundation.MsgIdentifier = "TheInputMouseRightDrop"
	MsgIDMouseMove          foundation.MsgIdentifier = "TheInputMouseMove"
	MsgIDMouseWheelUp       foundation.MsgIdentifier = "TheInputMouseWheelUp"
	MsgIDMouseWheelDown     foundation.MsgIdentifier = "TheInputMouseWheelDown"
	MsgIDKeyPressOff        foundation.MsgIdentifier = "TheInputKeyPressOff"
	MsgIDKeyPressOn         foundation.MsgIdentifier = "TheInputKeyPressOn"
	MsgIDKeyPressRepeat     foundation.MsgIdentifier = "TheInputKeyPressRepeat"
)

//EventCode2MsgID はイベントCodeと対応する入力系メッセージIDを返却します
func EventCode2MsgID(evcode data.EventCode) foundation.MsgIdentifier {
	table := map[data.EventCode]foundation.MsgIdentifier{
		data.MouseLeftDown:      MsgIDMouseLeftDown,
		data.MouseLeftUp:        MsgIDMouseLeftUp,
		data.MouseRightDown:     MsgIDMouseRightDown,
		data.MouseRightUp:       MsgIDMouseRightUp,
		data.MouseLeftDragging:  MsgIDMouseLeftDragging,
		data.MouseRightDragging: MsgIDMouseRightDragging,
		data.MouseLeftDrop:      MsgIDMouseLeftDrop,
		data.MouseRightDrop:     MsgIDMouseRightDrop,
		data.MouseMove:          MsgIDMouseMove,
		data.MouseWheelUp:       MsgIDMouseWheelUp,
		data.MouseWheelDown:     MsgIDMouseWheelDown,
		data.KeyPressOff:        MsgIDKeyPressOff,
		data.KeyPressOn:         MsgIDKeyPressOn,
		data.KeyPressRepeat:     MsgIDKeyPressRepeat,
	}
	return table[evcode]
}
