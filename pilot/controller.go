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
	leftDrag  bool // 左ボタンの押下 true: ドラッグ中　false: 非ドラッグ中
	rightDrag int8 // 右ボタンの押下　RightDragOff RightDragStart RightDragOn
}

// マウス右ボタン状態
const (
	RightDragOff int8 = iota
	RightDragStart
	RightDragOn
)

func (controller *Controller) Init() error {
	controller.leftDrag = false
	controller.rightDrag = RightDragOff
	return nil
}

/*
Runは入力を受け取るイベントハンドリングループです。
*/
func (controller *Controller) ReceiveEvent() (bool, data.Event, error) {
	// イベントがなかったという意味のイベント
	NoEvent := data.Event{
		Code: data.NoEvent,
	}
	for sdlEvent := sdl.PollEvent(); sdlEvent != nil; sdlEvent = sdl.PollEvent() {
		switch t := sdlEvent.(type) {
		case *sdl.QuitEvent:
			return false, NoEvent, nil
		case *sdl.MouseMotionEvent:
			return true, controller.motionEvent(t), nil
		case *sdl.MouseButtonEvent:
			return true, controller.buttonEvent(t), nil
		case *sdl.MouseWheelEvent:
			return true, controller.wheelEvent(t), nil
		case *sdl.KeyboardEvent:
			return true, controller.keyboardEvent(t), nil
		}
	}
	return true, NoEvent, nil
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
	evt.Mouse = data.Mouse{
		X:     sdlEvent.X,
		Y:     sdlEvent.Y,
		MoveX: sdlEvent.XRel,
		MoveY: sdlEvent.YRel,
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
	evt.Mouse = data.Mouse{
		X: sdlEvent.X,
		Y: sdlEvent.Y,
	}
	return evt
}

// マウスホイールを動かした時のイベント処理
func (c *Controller) wheelEvent(sdlEvent *sdl.MouseWheelEvent) data.Event {
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
func (c *Controller) keyboardEvent(sdlEvent *sdl.KeyboardEvent) data.Event {
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
	evt.Keyboard.Keycode = data.Scancode(sdlEvent.Keysym.Scancode)
	return evt
}
