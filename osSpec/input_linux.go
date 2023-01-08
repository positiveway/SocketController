//go:build linux

package osSpec

import (
	"github.com/bendahl/uinput"
	"github.com/positiveway/gofuncs"
)

var mouse uinput.Mouse
var keyboard uinput.Keyboard

func CloseInputResources() {
	if keyboard != nil {
		keyboard.Close()
	}
	if mouse != nil {
		mouse.Close()
	}
}

func InitInput() {
	var err error

	// initialize keyboard and check for possible errors
	keyboard, err = uinput.CreateKeyboard("/dev/uinput", []byte("testkeyboard"))
	gofuncs.CheckErr(err)

	// initialize mouse and check for possible errors
	mouse, err = uinput.CreateMouse("/dev/uinput", []byte("testmouse"))
	gofuncs.CheckErr(err)
}

func PressKeyOrMouse(key string) {
	switch key {
	case LeftMouse:
		mouse.LeftPress()
	case RightMouse:
		mouse.RightPress()
	case MiddleMouse:
		mouse.MiddlePress()
	default:
		code := gofuncs.GetOrPanic(lettersToCodes, key)
		keyboard.KeyDown(code)
	}
}

func ReleaseKeyOrMouse(key string) {
	switch key {
	case LeftMouse:
		mouse.LeftRelease()
	case RightMouse:
		mouse.RightRelease()
	case MiddleMouse:
		mouse.MiddleRelease()
	default:
		code := gofuncs.GetOrPanic(lettersToCodes, key)
		keyboard.KeyUp(code)
	}
}

func TypeLetter(letter string) {
	code := gofuncs.GetOrPanic(lettersToCodes, letter)
	keyboard.KeyPress(code)
}

func MoveMouse(x, y int) {
	//robotgo.MoveRelative(x, -y)
	mouse.Move(int32(x), int32(-y))
}

func ScrollHorizontal(direction int) {
	mouse.Wheel(true, int32(direction))
}

func ScrollVertical(direction int) {
	mouse.Wheel(false, int32(direction))
}
