package main

import (
	"SocketController/osSpec"
	"fmt"
	"github.com/positiveway/gofuncs"
)

func toNum(bytes []byte) int {
	num := int(bytes[1])
	switch bytes[0] {
	case 0:
		return num
	case 1:
		return -num
	default:
		panic(fmt.Sprintf("unsupported value: %v", bytes[0]))
	}
}

func control(event []byte) {
	//fmt.Println(event)

	if event[0] == 0 || event[0] == 1 {
		if len(event) != 4 {
			gofuncs.Panic("Incorrect length: %v", len(event))
		}
		x, y := toNum(event[0:2]), toNum(event[2:4])
		osSpec.MoveMouse(x, y)
		return
	}

	commandType := rune(event[0])
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
