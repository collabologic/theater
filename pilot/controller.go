package pilot

import (
	"github.com/collabologic/theater/data"
	"github.com/veandco/go-sdl2/sdl"
)

/*
Controllerは、マウス、キーボード、ジョイパッドなどの機器から入力を受け取り、
送信チャンネルにEvent構造体を出力します。

Controllerの呼び出し前には、sdl.init()が呼ばれている必要があります
*/
type Controller struct {
	leftDrag    bool        // 左ボタンの押下 true: ドラッグ中　false: 非ドラッグ中
	rightDrag   int8        // 右ボタンの押下　RightDragOff RightDragStart RightDragOn
	LogicalSize             // 論理画面サイズ
	window      *sdl.Window // ウィンドwオブジェクト
}

// マウス右ボタン状態
const (
	RightDragOff int8 = iota
	RightDragStart
	RightDragOn
)

/*
Controllerを生成する
*/
func NewController(win *sdl.Window, ls LogicalSize) *Controller {

	controller := Controller{}
	controller.window = win
	controller.leftDrag = false
	controller.rightDrag = RightDragOff
	controller.LogicalSize = ls
	return &controller
}

/*
Runは入力を受け取るイベントハンドリングループです。
*/
func (controller *Controller) ReceiveEvent() (bool, data.Event, error) {
	var event data.Event

	// イベントがなかったという意味のイベント
	NoEvent := data.Event{
		Code: data.NoEvent,
	}

	sdlEvent := sdl.PollEvent()
	switch t := sdlEvent.(type) {
	case *sdl.QuitEvent:
		return false, NoEvent, nil
	case *sdl.MouseMotionEvent:
		event = controller.motionEvent(t)
	case *sdl.MouseButtonEvent:
		event = controller.buttonEvent(t)
	case *sdl.MouseWheelEvent:
		event = controller.wheelEvent(t)
	case *sdl.KeyboardEvent:
		event = controller.keyboardEvent(t)
	}

	return true, event, nil
}

// マウスカーソルが動いた時のイベント処理
func (controller *Controller) motionEvent(sdlEvent *sdl.MouseMotionEvent) data.Event {
	evt := data.Event{}
	evt.Device = data.DeviceMouse
	switch sdlEvent.State {
	case sdl.BUTTON_LEFT:
		controller.leftDrag = true
		evt.Code = data.MouseLeftDragging
	default:
		if controller.rightDrag == RightDragStart {
			controller.rightDrag = RightDragOn
			evt.Code = data.MouseRightDragging
		} else if controller.rightDrag == RightDragOn {
			evt.Code = data.MouseRightDragging
		} else {
			evt.Code = data.MouseMove
		}
	}
	x, y := controller.screen2logical(controller.window, sdlEvent.X, sdlEvent.Y)
	mx, my := controller.screen2logicalMove(controller.window, sdlEvent.XRel, sdlEvent.YRel)
	evt.Mouse = data.Mouse{
		X:     x,
		Y:     y,
		MoveX: mx,
		MoveY: my,
	}
	return evt
}

// マウスボタンの状態を変えた時のイベント処理
func (controller *Controller) buttonEvent(sdlEvent *sdl.MouseButtonEvent) data.Event {
	evt := data.Event{}
	evt.Device = data.DeviceMouse
	switch sdlEvent.Button {
	case sdl.BUTTON_LEFT:
		switch sdlEvent.State {
		case sdl.PRESSED:
			evt.Code = data.MouseLeftDown
		case sdl.RELEASED:
			if controller.leftDrag {
				controller.leftDrag = false
				evt.Code = data.MouseLeftDrop
			} else {
				evt.Code = data.MouseLeftUp
			}
		}
	case sdl.BUTTON_RIGHT:
		switch sdlEvent.State {
		case sdl.PRESSED:
			controller.rightDrag = RightDragStart
			evt.Code = data.MouseRightDown
		case sdl.RELEASED:
			if controller.rightDrag == RightDragOn {
				controller.rightDrag = RightDragOff
				evt.Code = data.MouseRightDrop
			} else {
				evt.Code = data.MouseRightUp
			}
		}
	default:
		evt.Code = data.Unknown
	}
	x, y := controller.screen2logical(controller.window, sdlEvent.X, sdlEvent.Y)
	evt.Mouse = data.Mouse{
		X: x,
		Y: y,
	}
	return evt
}

// マウスホイールを動かした時のイベント処理
func (controller *Controller) wheelEvent(sdlEvent *sdl.MouseWheelEvent) data.Event {
	evt := data.Event{}
	evt.Device = data.DeviceMouse
	if sdlEvent.Y > 0 {
		evt.Code = data.MouseWheelDown
	} else if sdlEvent.Y < 0 {
		evt.Code = data.MouseWheelUp
	} else {
		evt.Code = data.Unknown
	}
	return evt
}

// キーボードを動かした時のイベント処理
func (controller *Controller) keyboardEvent(sdlEvent *sdl.KeyboardEvent) data.Event {
	evt := data.Event{}
	evt.Device = data.DeviceKeyboard
	switch sdlEvent.Type {
	case sdl.KEYDOWN: // 押している、または押しっぱなし
		if sdlEvent.Repeat > 0 { // 押しっぱなしだった場合
			evt.Code = data.KeyPressRepeat
			evt.Keyboard.Repeat = sdlEvent.Repeat
		} else { // 押し始めた場合
			evt.Code = data.KeyPressOn
		}
	case sdl.KEYUP: // 離した場合
		evt.Code = data.KeyPressOff
	}
	evt.Keyboard.Keycode = data.Scancode(sdlEvent.Keysym.Sym)
	return evt
}
