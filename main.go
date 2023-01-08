package main

import (
	"SocketController/osSpec"
	"github.com/positiveway/gofuncs"
	"strings"
)

func control(event string) {
	//fmt.Println(event)

	commandType := event[0]
	command := event[1:]
	switch commandType {
	case 'm':
		nums := strings.Split(command, ",")
		x, y := gofuncs.StrToInt(nums[0]), gofuncs.StrToInt(nums[1])
		osSpec.MoveMouse(x, y)
	case 'p':
		//gofuncs.Print("press")
		osSpec.PressKeyOrMouse(command)
	case 'r':
		//gofuncs.Print("release")
		osSpec.ReleaseKeyOrMouse(command)
	case 'l':
		osSpec.TypeLetter(command)
	}
}
func main() {
	osSpec.InitInput()
	RunWebSocket()
}
