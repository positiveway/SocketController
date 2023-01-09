package main

import (
	"SocketController/osSpec"
	"github.com/positiveway/gofuncs"
)

func control(event []byte) {
	//fmt.Println(event)

	firstByte := event[0]

	if firstByte <= 3 {
		x := int(event[1])
		y := int(event[2])

		switch firstByte {
		case 1:
			x *= -1
		case 2:
			y *= -1
		case 3:
			x *= -1
			y *= -1
		}
		osSpec.MoveMouse(x, y)
		return
	}

	commandType := rune(firstByte)
	command := string(event[1:])
	switch commandType {
	case 'p':
		//gofuncs.Print("press")
		osSpec.PressKeyOrMouse(command)
	case 'r':
		//gofuncs.Print("release")
		osSpec.ReleaseKeyOrMouse(command)
	case 'l':
		osSpec.TypeLetter(command)
	default:
		gofuncs.Panic("Unknown command %s", string(event))
	}
}
func main() {
	osSpec.InitInput()
	RunWebSocket()
}
