package main

import (
	"SocketController/osSpec"
)

func toNum(oneByte byte) int {
	num := int(oneByte)
	if num > 128 {
		num -= 256
	}
	return num
}

func control(event []byte) {
	//fmt.Println(event)

	if len(event) == 2 {
		if event[0] == 128 {
			y := toNum(event[1])
			osSpec.ScrollVertical(y)
		} else if event[1] == 128 {
			x := toNum(event[0])
			osSpec.ScrollHorizontal(x)
		} else {
			x := toNum(event[0])
			y := toNum(event[1])
			//fmt.Printf("%v %v\n", x, y)
			osSpec.MoveMouse(x, y)
		}
	} else if len(event) == 1 {
		if event[0] > 128 {
			event[0] -= 128
			osSpec.PressKeyOrMouse(int(event[0]))
		} else {
			osSpec.ReleaseKeyOrMouse(int(event[0]))
		}
	}
}
func main() {
	osSpec.InitInput()
	RunWebSocket()
}
