package main

import (
	"bufio"
	"fmt"
	"net"
	"net/rpc"
	"os"
	"strconv"
	"strings"
)

type Server struct{}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func server(ln net.Listener) {
	rpc.Register(new(Server))

	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}

func client(msg string) {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	number, err := strconv.Atoi(msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	var result int64
	err = c.Call("Server.Negate", int64(number), &result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Server.Negate(%d) = %d\n", number, result)

}

func ReadMsg() {
	for {
		in := bufio.NewReader(os.Stdin)

		msg, err := in.ReadString('\n')
		if err != nil || msg == "\n" {
			return
		} else {
			client(strings.ReplaceAll(msg, "\n", ""))
		}
	}
}

func main() {

	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	go server(ln)
	ReadMsg()
}
