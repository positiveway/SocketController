package main

import (
	"fmt"
	"github.com/bendahl/uinput"
	"github.com/positiveway/gofuncs"
	"net"
	"runtime"
	"runtime/debug"
)

func toNum(oneByte byte) int32 {
	num := int32(oneByte)
	if num > 128 {
		num -= 256
	}
	return num
}

func getSign(num int32) int32 {
	if num > 0 {
		return 1
	} else if num < 0 {
		return -1
	} else {
		return 0
	}
}

func main() {
	const LeftMouse = 90
	const RightMouse = 91
	const MiddleMouse = 92

	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("testkeyboard"))
	gofuncs.CheckErr(err)
	defer keyboard.Close()
	mouse, err := uinput.CreateMouse("/dev/uinput", []byte("testmouse"))
	gofuncs.CheckErr(err)
	defer mouse.Close()

	addr := net.UDPAddr{
		Port: 5005,
		IP:   net.ParseIP("0.0.0.0"),
	}

	server, err := net.ListenUDP("udp4", &addr)
	if err != nil {
		panic(fmt.Sprintf("Client is already running: %v", err))
	}
	defer server.Close()
	gofuncs.Print("Listening at %v", addr.String())

	msg := make([]byte, 2)

	debug.SetGCPercent(-1)
	runtime.GC()

	for {
		msgLen, err := server.Read(msg)
		if err != nil {
			gofuncs.Print("Client disconnected. Read err: %v", err)
			continue
		}
		//fmt.Printf("%v %v\n", int(msg[0]), int(msg[1]))

		if msgLen == 2 {
			if msg[0] == 128 {
				y := toNum(msg[1])
				y = getSign(y)
				mouse.Wheel(false, y)
			} else if msg[1] == 128 {
				x := toNum(msg[0])
				x = getSign(x)
				mouse.Wheel(true, x)
			} else {
				x := toNum(msg[0])
				y := toNum(msg[1])
				//fmt.Printf("%v %v\n", x, y)
				mouse.Move(x, -y)
			}
		} else if msgLen == 1 {
			if msg[0] > 128 {
				msg[0] -= 128
				switch msg[0] {
				case LeftMouse:
					mouse.LeftPress()
				case RightMouse:
					mouse.RightPress()
				case MiddleMouse:
					mouse.MiddlePress()
				default:
					keyboard.KeyDown(int(msg[0]))
				}
			} else {
				switch msg[0] {
				case LeftMouse:
					mouse.LeftRelease()
				case RightMouse:
					mouse.RightRelease()
				case MiddleMouse:
					mouse.MiddleRelease()
				default:
					keyboard.KeyUp(int(msg[0]))
				}
				runtime.GC()
			}
		}
	}
}
