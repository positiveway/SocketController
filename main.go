package main

import (
	"SocketController/osSpec"
	"fmt"
	"github.com/positiveway/gofuncs"
)

func toNum(bytes []byte) int {
	num := int(bytes[1])
	switch rune(bytes[0]) {
	case '+':
		return num
	case '-':
		return -num
	default:
		panic(fmt.Sprintf("unsupported value: %s", string(bytes)))
	}
}

func control(event []byte) {
	//fmt.Println(event)

	commandType := rune(event[0])
	command := string(event[1:])

	switch commandType {
	case '+', '-':
		if len(event) != 4 {
			gofuncs.Panic("Incorrect length: %v", len(event))
		}
		x, y := toNum(event[0:2]), toNum(event[2:4])
		osSpec.MoveMouse(x, y)
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
