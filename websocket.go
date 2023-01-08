package main

import (
	"fmt"
	"net"
)

func RunWebSocket() {
	addr := net.UDPAddr{
		Port: 5005,
		IP:   net.ParseIP("0.0.0.0"),
	}
	server, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(fmt.Sprintf("Client is already running: %v", err))
	}
	fmt.Printf("Listening at %v", addr.String())

	p := make([]byte, 128)

	for {
		nn, _, err := server.ReadFromUDP(p)
		if err != nil {
			fmt.Printf("Read err  %v", err)
			continue
		}

		content := string(p[:nn])
		control(content)
	}
}
