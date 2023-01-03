package main

import (
	"SocketController/osSpec"
	"fmt"
)

func control(event string) {
	fmt.Println(event)

	commandType := event[0]
	command := event[1:]
	switch commandType {
	case 'l':
		osSpec.TypeLetter(command)
	}
}
func main() {
	osSpec.InitInput()
	RunWebSocket()
}
