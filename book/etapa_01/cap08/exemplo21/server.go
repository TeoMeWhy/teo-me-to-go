package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server(ln net.Listener) {

	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// trata a conexão
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	//recebe a mensagem
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Recebida:", msg)
	}
	c.Close()
}

func main() {

	// ouve uma porta
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	go server(ln)

	var input string
	fmt.Scanln(&input)
}
