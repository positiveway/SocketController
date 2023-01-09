package main

import (
	"SocketController/osSpec"
	"github.com/positiveway/gofuncs"
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
		x := toNum(event[0])
		y := toNum(event[1])

		osSpec.MoveMouse(x, y)
		return
	}

	commandType := string(event[0:2])
	command := string(event[2:])
	switch commandType {
	case "pr":
		//gofuncs.Print("press")
		osSpec.PressKeyOrMouse(command)
	case "re":
		//gofuncs.Print("release")
		osSpec.ReleaseKeyOrMouse(command)
	case "ty":
		osSpec.TypeLetter(command)
	default:
		gofuncs.Panic("Unknown command %s", string(event))
	}
}
func main() {
	osSpec.InitInput()
	RunWebSocket()
}
