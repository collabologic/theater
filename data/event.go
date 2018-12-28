package data

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

//Eventは、マウス、キーボードなどからの入力情報です。
type Event struct {
	Device   Device    // 入力機器
	Code     EventCode // 動作の種類
	Keyboard           // キーボード
	Mouse              // マウス
	Joypad             // ジョイパッド
}

func (event *Event) String() string {
	switch event.Device {
	case DeviceKeyboard:
		return fmt.Sprintf("Device: KEYBOARD Code:%d", event.Code, event.Keyboard)
	case DeviceMouse:
		return fmt.Sprintf("Device: MOUSE Code:%d", event.Code, event.Mouse)
	}
	return "Device: Unkown"

}

// デバイスの種別を取り扱う列挙型です
type Device int32

// Device型の値
const (
	DeviceUnkown   Device = iota // 不明なデバイス
	DeviceKeyboard               // キーボード
	DeviceMouse                  // マウス
	DeviceJoypad                 // ジョイパッド
)

// 動作の種類の列挙型です
type EventCode int32

// EventType型の値
const (
	NoEvent            EventCode = iota // 何もなかった場合
	Unknown                             // 不明のイベント
	MouseLeftDown                       // 左ボタンを押した
	MouseLeftUp                         // 左ボタン離した
	MouseRightDown                      // 右ボタン押した
	MouseRightUp                        // 右ボタン離した
	MouseLeftDragging                   // 左ボタンを押したまま移動した
	MouseRightDragging                  // 右ボタンを押したまま移動した
	MouseLeftDrop                       // 左ボタンを押したまま移動して離した
	MouseRightDrop                      // 右ボタンを押したたまま移動して離した
	MouseMove                           // ボタンを押さずに移動した
	MouseWheelUp                        // ホイールを上に動かした
	MouseWheelDown                      // ホイールを下に動かした
	KeyPressOff                         // 離した時
	KeyPressOn                          // 押した時
	KeyPressRepeat                      // キーを押し続けている時
)

// キーボードからの入力情報です。
type Keyboard struct {
	Keycode Scancode // SDLスキャンコード（物理キーコード）
	Repeat  uint8    // キーが押しっぱなしなら1（Event.code>0の場合のみ）
}

// スキャンコード（キーのID)
type Scancode sdl.Scancode

// Scancode型の値
const (
	K_UNKNOWN    sdl.Scancode = sdl.K_UNKNOWN    // "" (no name, empty string)
	K_RETURN                  = sdl.K_RETURN     // "Return" (the Enter key (main keyboard))
	K_ESCAPE                  = sdl.K_ESCAPE     // "Escape" (the Esc key)
	K_BACKSPACE               = sdl.K_BACKSPACE  // "Backspace"
	K_TAB                     = sdl.K_TAB        // "Tab" (the Tab key)
	K_SPACE                   = sdl.K_SPACE      // "Space" (the Space Bar key(s))
	K_EXCLAIM                 = sdl.K_EXCLAIM    // "!"
	K_QUOTEDBL                = sdl.K_QUOTEDBL   // """
	K_HASH                    = sdl.K_HASH       // "#"
	K_PERCENT                 = sdl.K_PERCENT    // "%"
	K_DOLLAR                  = sdl.K_DOLLAR     // "$"
	K_AMPERSAND               = sdl.K_AMPERSAND  // "&"
	K_QUOTE                   = sdl.K_QUOTE      // "'"
	K_LEFTPAREN               = sdl.K_LEFTPAREN  // "("
	K_RIGHTPAREN              = sdl.K_RIGHTPAREN // ")"
	K_ASTERISK                = sdl.K_ASTERISK   // "*"
	K_PLUS                    = sdl.K_PLUS       // "+"
	K_COMMA                   = sdl.K_COMMA      // ","
	K_MINUS                   = sdl.K_MINUS      // "-"
	K_PERIOD                  = sdl.K_PERIOD     // "."
	K_SLASH                   = sdl.K_SLASH      // "/"
	K_0                       = sdl.K_0          // "0"
	K_1                       = sdl.K_1          // "1"
	K_2                       = sdl.K_2          // "2"
	K_3                       = sdl.K_3          // "3"
	K_4                       = sdl.K_4          // "4"
	K_5                       = sdl.K_5          // "5"
	K_6                       = sdl.K_6          // "6"
	K_7                       = sdl.K_7          // "7"
	K_8                       = sdl.K_8          // "8"
	K_9                       = sdl.K_9          // "9"
	K_COLON                   = sdl.K_COLON      // ":"
	K_SEMICOLON               = sdl.K_SEMICOLON  // ";"
	K_LESS                    = sdl.K_LESS       // "<"
	K_EQUALS                  = sdl.K_EQUALS     // "="
	K_GREATER                 = sdl.K_GREATER    // ">"
	K_QUESTION                = sdl.K_QUESTION   // "?"
	K_AT                      = sdl.K_AT         // "@"
	/*
	   Skip uppercase letters
	*/
	K_LEFTBRACKET  = sdl.K_LEFTBRACKET  // "["
	K_BACKSLASH    = sdl.K_BACKSLASH    // "\"
	K_RIGHTBRACKET = sdl.K_RIGHTBRACKET // "]"
	K_CARET        = sdl.K_CARET        // "^"
	K_UNDERSCORE   = sdl.K_UNDERSCORE   // "_"
	K_BACKQUOTE    = sdl.K_BACKQUOTE    // "`"
	K_a            = sdl.K_a            // "A"
	K_b            = sdl.K_b            // "B"
	K_c            = sdl.K_c            // "C"
	K_d            = sdl.K_d            // "D"
	K_e            = sdl.K_e            // "E"
	K_f            = sdl.K_f            // "F"
	K_g            = sdl.K_g            // "G"
	K_h            = sdl.K_h            // "H"
	K_i            = sdl.K_i            // "I"
	K_j            = sdl.K_j            // "J"
	K_k            = sdl.K_k            // "K"
	K_l            = sdl.K_l            // "L"
	K_m            = sdl.K_m            // "M"
	K_n            = sdl.K_n            // "N"
	K_o            = sdl.K_o            // "O"
	K_p            = sdl.K_p            // "P"
	K_q            = sdl.K_q            // "Q"
	K_r            = sdl.K_r            // "R"
	K_s            = sdl.K_s            // "S"
	K_t            = sdl.K_t            // "T"
	K_u            = sdl.K_u            // "U"
	K_v            = sdl.K_v            // "V"
	K_w            = sdl.K_w            // "W"
	K_x            = sdl.K_x            // "X"
	K_y            = sdl.K_y            // "Y"
	K_z            = sdl.K_z            // "Z"

	K_CAPSLOCK = sdl.K_CAPSLOCK // "CapsLock"

	K_F1  = sdl.K_F1  // "F1"
	K_F2  = sdl.K_F2  // "F2"
	K_F3  = sdl.K_F3  // "F3"
	K_F4  = sdl.K_F4  // "F4"
	K_F5  = sdl.K_F5  // "F5"
	K_F6  = sdl.K_F6  // "F6"
	K_F7  = sdl.K_F7  // "F7"
	K_F8  = sdl.K_F8  // "F8"
	K_F9  = sdl.K_F9  // "F9"
	K_F10 = sdl.K_F10 // "F10"
	K_F11 = sdl.K_F11 // "F11"
	K_F12 = sdl.K_F12 // "F12"

	K_PRINTSCREEN = sdl.K_PRINTSCREEN // "PrintScreen"
	K_SCROLLLOCK  = sdl.K_SCROLLLOCK  // "ScrollLock"
	K_PAUSE       = sdl.K_PAUSE       // "Pause" (the Pause / Break key)
	K_INSERT      = sdl.K_INSERT      // "Insert" (insert on PC, help on some Mac keyboards (but does send code 73, not 117))
	K_HOME        = sdl.K_HOME        // "Home"
	K_PAGEUP      = sdl.K_PAGEUP      // "PageUp"
	K_DELETE      = sdl.K_DELETE      // "Delete"
	K_END         = sdl.K_END         // "End"
	K_PAGEDOWN    = sdl.K_PAGEDOWN    // "PageDown"
	K_RIGHT       = sdl.K_RIGHT       // "Right" (the Right arrow key (navigation keypad))
	K_LEFT        = sdl.K_LEFT        // "Left" (the Left arrow key (navigation keypad))
	K_DOWN        = sdl.K_DOWN        // "Down" (the Down arrow key (navigation keypad))
	K_UP          = sdl.K_UP          // "Up" (the Up arrow key (navigation keypad))

	K_NUMLOCKCLEAR = sdl.K_NUMLOCKCLEAR // "Numlock" (the Num Lock key (PC) / the Clear key (Mac))
	K_KP_DIVIDE    = sdl.K_KP_DIVIDE    // "Keypad /" (the / key (numeric keypad))
	K_KP_MULTIPLY  = sdl.K_KP_MULTIPLY  // "Keypad *" (the * key (numeric keypad))
	K_KP_MINUS     = sdl.K_KP_MINUS     // "Keypad -" (the - key (numeric keypad))
	K_KP_PLUS      = sdl.K_KP_PLUS      // "Keypad +" (the + key (numeric keypad))
	K_KP_ENTER     = sdl.K_KP_ENTER     // "Keypad Enter" (the Enter key (numeric keypad))
	K_KP_1         = sdl.K_KP_1         // "Keypad 1" (the 1 key (numeric keypad))
	K_KP_2         = sdl.K_KP_2         // "Keypad 2" (the 2 key (numeric keypad))
	K_KP_3         = sdl.K_KP_3         // "Keypad 3" (the 3 key (numeric keypad))
	K_KP_4         = sdl.K_KP_4         // "Keypad 4" (the 4 key (numeric keypad))
	K_KP_5         = sdl.K_KP_5         // "Keypad 5" (the 5 key (numeric keypad))
	K_KP_6         = sdl.K_KP_6         // "Keypad 6" (the 6 key (numeric keypad))
	K_KP_7         = sdl.K_KP_7         // "Keypad 7" (the 7 key (numeric keypad))
	K_KP_8         = sdl.K_KP_8         // "Keypad 8" (the 8 key (numeric keypad))
	K_KP_9         = sdl.K_KP_9         // "Keypad 9" (the 9 key (numeric keypad))
	K_KP_0         = sdl.K_KP_0         // "Keypad 0" (the 0 key (numeric keypad))
	K_KP_PERIOD    = sdl.K_KP_PERIOD    // "Keypad ." (the . key (numeric keypad))

	K_APPLICATION    = sdl.K_APPLICATION    // "Application" (the Application / Compose / Context Menu (Windows) key)
	K_POWER          = sdl.K_POWER          // "Power" (The USB document says this is a status flag, not a physical key - but some Mac keyboards do have a power key.)
	K_KP_EQUALS      = sdl.K_KP_EQUALS      // "Keypad =" (the = key (numeric keypad))
	K_F13            = sdl.K_F13            // "F13"
	K_F14            = sdl.K_F14            // "F14"
	K_F15            = sdl.K_F15            // "F15"
	K_F16            = sdl.K_F16            // "F16"
	K_F17            = sdl.K_F17            // "F17"
	K_F18            = sdl.K_F18            // "F18"
	K_F19            = sdl.K_F19            // "F19"
	K_F20            = sdl.K_F20            // "F20"
	K_F21            = sdl.K_F21            // "F21"
	K_F22            = sdl.K_F22            // "F22"
	K_F23            = sdl.K_F23            // "F23"
	K_F24            = sdl.K_F24            // "F24"
	K_EXECUTE        = sdl.K_EXECUTE        // "Execute"
	K_HELP           = sdl.K_HELP           // "Help"
	K_MENU           = sdl.K_MENU           // "Menu"
	K_SELECT         = sdl.K_SELECT         // "Select"
	K_STOP           = sdl.K_STOP           // "Stop"
	K_AGAIN          = sdl.K_AGAIN          // "Again" (the Again key (Redo))
	K_UNDO           = sdl.K_UNDO           // "Undo"
	K_CUT            = sdl.K_CUT            // "Cut"
	K_COPY           = sdl.K_COPY           // "Copy"
	K_PASTE          = sdl.K_PASTE          // "Paste"
	K_FIND           = sdl.K_FIND           // "Find"
	K_MUTE           = sdl.K_MUTE           // "Mute"
	K_VOLUMEUP       = sdl.K_VOLUMEUP       // "VolumeUp"
	K_VOLUMEDOWN     = sdl.K_VOLUMEDOWN     // "VolumeDown"
	K_KP_COMMA       = sdl.K_KP_COMMA       // "Keypad ," (the Comma key (numeric keypad))
	K_KP_EQUALSAS400 = sdl.K_KP_EQUALSAS400 // "Keypad = (AS400)" (the Equals AS400 key (numeric keypad))

	K_ALTERASE   = sdl.K_ALTERASE   // "AltErase" (Erase-Eaze)
	K_SYSREQ     = sdl.K_SYSREQ     // "SysReq" (the SysReq key)
	K_CANCEL     = sdl.K_CANCEL     // "Cancel"
	K_CLEAR      = sdl.K_CLEAR      // "Clear"
	K_PRIOR      = sdl.K_PRIOR      // "Prior"
	K_RETURN2    = sdl.K_RETURN2    // "Return"
	K_SEPARATOR  = sdl.K_SEPARATOR  // "Separator"
	K_OUT        = sdl.K_OUT        // "Out"
	K_OPER       = sdl.K_OPER       // "Oper"
	K_CLEARAGAIN = sdl.K_CLEARAGAIN // "Clear / Again"
	K_CRSEL      = sdl.K_CRSEL      // "CrSel"
	K_EXSEL      = sdl.K_EXSEL      // "ExSel"

	K_KP_00              = sdl.K_KP_00              // "Keypad 00" (the 00 key (numeric keypad))
	K_KP_000             = sdl.K_KP_000             // "Keypad 000" (the 000 key (numeric keypad))
	K_THOUSANDSSEPARATOR = sdl.K_THOUSANDSSEPARATOR // "ThousandsSeparator" (the Thousands Separator key)
	K_DECIMALSEPARATOR   = sdl.K_DECIMALSEPARATOR   // "DecimalSeparator" (the Decimal Separator key)
	K_CURRENCYUNIT       = sdl.K_CURRENCYUNIT       // "CurrencyUnit" (the Currency Unit key)
	K_CURRENCYSUBUNIT    = sdl.K_CURRENCYSUBUNIT    // "CurrencySubUnit" (the Currency Subunit key)
	K_KP_LEFTPAREN       = sdl.K_KP_LEFTPAREN       // "Keypad (" (the Left Parenthesis key (numeric keypad))
	K_KP_RIGHTPAREN      = sdl.K_KP_RIGHTPAREN      // "Keypad )" (the Right Parenthesis key (numeric keypad))
	K_KP_LEFTBRACE       = sdl.K_KP_LEFTBRACE       // "Keypad {" (the Left Brace key (numeric keypad))
	K_KP_RIGHTBRACE      = sdl.K_KP_RIGHTBRACE      // "Keypad }" (the Right Brace key (numeric keypad))
	K_KP_TAB             = sdl.K_KP_TAB             // "Keypad Tab" (the Tab key (numeric keypad))
	K_KP_BACKSPACE       = sdl.K_KP_BACKSPACE       // "Keypad Backspace" (the Backspace key (numeric keypad))
	K_KP_A               = sdl.K_KP_A               // "Keypad A" (the A key (numeric keypad))
	K_KP_B               = sdl.K_KP_B               // "Keypad B" (the B key (numeric keypad))
	K_KP_C               = sdl.K_KP_C               // "Keypad C" (the C key (numeric keypad))
	K_KP_D               = sdl.K_KP_D               // "Keypad D" (the D key (numeric keypad))
	K_KP_E               = sdl.K_KP_E               // "Keypad E" (the E key (numeric keypad))
	K_KP_F               = sdl.K_KP_F               // "Keypad F" (the F key (numeric keypad))
	K_KP_XOR             = sdl.K_KP_XOR             // "Keypad XOR" (the XOR key (numeric keypad))
	K_KP_POWER           = sdl.K_KP_POWER           // "Keypad ^" (the Power key (numeric keypad))
	K_KP_PERCENT         = sdl.K_KP_PERCENT         // "Keypad %" (the Percent key (numeric keypad))
	K_KP_LESS            = sdl.K_KP_LESS            // "Keypad <" (the Less key (numeric keypad))
	K_KP_GREATER         = sdl.K_KP_GREATER         // "Keypad >" (the Greater key (numeric keypad))
	K_KP_AMPERSAND       = sdl.K_KP_AMPERSAND       // "Keypad &" (the & key (numeric keypad))
	K_KP_DBLAMPERSAND    = sdl.K_KP_DBLAMPERSAND    // "Keypad &&" (the && key (numeric keypad))
	K_KP_VERTICALBAR     = sdl.K_KP_VERTICALBAR     // "Keypad |" (the | key (numeric keypad))
	K_KP_DBLVERTICALBAR  = sdl.K_KP_DBLVERTICALBAR  // "Keypad ||" (the || key (numeric keypad))
	K_KP_COLON           = sdl.K_KP_COLON           // "Keypad :" (the : key (numeric keypad))
	K_KP_HASH            = sdl.K_KP_HASH            // "Keypad #" (the # key (numeric keypad))
	K_KP_SPACE           = sdl.K_KP_SPACE           // "Keypad Space" (the Space key (numeric keypad))
	K_KP_AT              = sdl.K_KP_AT              // "Keypad @" (the @ key (numeric keypad))
	K_KP_EXCLAM          = sdl.K_KP_EXCLAM          // "Keypad !" (the ! key (numeric keypad))
	K_KP_MEMSTORE        = sdl.K_KP_MEMSTORE        // "Keypad MemStore" (the Mem Store key (numeric keypad))
	K_KP_MEMRECALL       = sdl.K_KP_MEMRECALL       // "Keypad MemRecall" (the Mem Recall key (numeric keypad))
	K_KP_MEMCLEAR        = sdl.K_KP_MEMCLEAR        // "Keypad MemClear" (the Mem Clear key (numeric keypad))
	K_KP_MEMADD          = sdl.K_KP_MEMADD          // "Keypad MemAdd" (the Mem Add key (numeric keypad))
	K_KP_MEMSUBTRACT     = sdl.K_KP_MEMSUBTRACT     // "Keypad MemSubtract" (the Mem Subtract key (numeric keypad))
	K_KP_MEMMULTIPLY     = sdl.K_KP_MEMMULTIPLY     // "Keypad MemMultiply" (the Mem Multiply key (numeric keypad))
	K_KP_MEMDIVIDE       = sdl.K_KP_MEMDIVIDE       // "Keypad MemDivide" (the Mem Divide key (numeric keypad))
	K_KP_PLUSMINUS       = sdl.K_KP_PLUSMINUS       // "Keypad +/-" (the +/- key (numeric keypad))
	K_KP_CLEAR           = sdl.K_KP_CLEAR           // "Keypad Clear" (the Clear key (numeric keypad))
	K_KP_CLEARENTRY      = sdl.K_KP_CLEARENTRY      // "Keypad ClearEntry" (the Clear Entry key (numeric keypad))
	K_KP_BINARY          = sdl.K_KP_BINARY          // "Keypad Binary" (the Binary key (numeric keypad))
	K_KP_OCTAL           = sdl.K_KP_OCTAL           // "Keypad Octal" (the Octal key (numeric keypad))
	K_KP_DECIMAL         = sdl.K_KP_DECIMAL         // "Keypad Decimal" (the Decimal key (numeric keypad))
	K_KP_HEXADECIMAL     = sdl.K_KP_HEXADECIMAL     // "Keypad Hexadecimal" (the Hexadecimal key (numeric keypad))

	K_LCTRL  = sdl.K_LCTRL  // "Left Ctrl"
	K_LSHIFT = sdl.K_LSHIFT // "Left Shift"
	K_LALT   = sdl.K_LALT   // "Left Alt" (alt, option)
	K_LGUI   = sdl.K_LGUI   // "Left GUI" (windows, command (apple), meta)
	K_RCTRL  = sdl.K_RCTRL  // "Right Ctrl"
	K_RSHIFT = sdl.K_RSHIFT // "Right Shift"
	K_RALT   = sdl.K_RALT   // "Right Alt" (alt, option)
	K_RGUI   = sdl.K_RGUI   // "Right GUI" (windows, command (apple), meta)

	K_MODE = sdl.K_MODE // "ModeSwitch" (I'm not sure if this is really not covered by any of the above, but since there's a special KMOD_MODE for it I'm adding it here)

	K_AUDIONEXT    = sdl.K_AUDIONEXT    // "AudioNext" (the Next Track media key)
	K_AUDIOPREV    = sdl.K_AUDIOPREV    // "AudioPrev" (the Previous Track media key)
	K_AUDIOSTOP    = sdl.K_AUDIOSTOP    // "AudioStop" (the Stop media key)
	K_AUDIOPLAY    = sdl.K_AUDIOPLAY    // "AudioPlay" (the Play media key)
	K_AUDIOMUTE    = sdl.K_AUDIOMUTE    // "AudioMute" (the Mute volume key)
	K_MEDIASELECT  = sdl.K_MEDIASELECT  // "MediaSelect" (the Media Select key)
	K_WWW          = sdl.K_WWW          // "WWW" (the WWW/World Wide Web key)
	K_MAIL         = sdl.K_MAIL         // "Mail" (the Mail/eMail key)
	K_CALCULATOR   = sdl.K_CALCULATOR   // "Calculator" (the Calculator key)
	K_COMPUTER     = sdl.K_COMPUTER     // "Computer" (the My Computer key)
	K_AC_SEARCH    = sdl.K_AC_SEARCH    // "AC Search" (the Search key (application control keypad))
	K_AC_HOME      = sdl.K_AC_HOME      // "AC Home" (the Home key (application control keypad))
	K_AC_BACK      = sdl.K_AC_BACK      // "AC Back" (the Back key (application control keypad))
	K_AC_FORWARD   = sdl.K_AC_FORWARD   // "AC Forward" (the Forward key (application control keypad))
	K_AC_STOP      = sdl.K_AC_STOP      // "AC Stop" (the Stop key (application control keypad))
	K_AC_REFRESH   = sdl.K_AC_REFRESH   // "AC Refresh" (the Refresh key (application control keypad))
	K_AC_BOOKMARKS = sdl.K_AC_BOOKMARKS // "AC Bookmarks" (the Bookmarks key (application control keypad))

	K_BRIGHTNESSDOWN = sdl.K_BRIGHTNESSDOWN // "BrightnessDown" (the Brightness Down key)
	K_BRIGHTNESSUP   = sdl.K_BRIGHTNESSUP   // "BrightnessUp" (the Brightness Up key)
	K_DISPLAYSWITCH  = sdl.K_DISPLAYSWITCH  // "DisplaySwitch" (display mirroring/dual display switch, video mode switch)
	K_KBDILLUMTOGGLE = sdl.K_KBDILLUMTOGGLE // "KBDIllumToggle" (the Keyboard Illumination Toggle key)
	K_KBDILLUMDOWN   = sdl.K_KBDILLUMDOWN   // "KBDIllumDown" (the Keyboard Illumination Down key)
	K_KBDILLUMUP     = sdl.K_KBDILLUMUP     // "KBDIllumUp" (the Keyboard Illumination Up key)
	K_EJECT          = sdl.K_EJECT          // "Eject" (the Eject key)
	K_SLEEP          = sdl.K_SLEEP          // "Sleep" (the Sleep key)
)

// マウスからの入力情報です
type Mouse struct {
	X     int32   // 現在座標X
	Y     int32   // 現在座標Y
	MoveX float64 // X移動量
	MoveY float64 // Y移動量
}

// ジョイパッドからの入力情報です。
type Joypad struct {
	// TODO: 現在未実装
}
