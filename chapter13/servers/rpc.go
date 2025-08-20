package main

import (
	"fmt"
	"net"
	"net/rpc"
)

/*
	The `net/rpc` (remote procedure call) and `net/rpc/jsonrpc` packages provide an easy way 
	to expose methods so they can be invoked over a network.

	This program is similar to the TCP example, except now we created an object 
	to hold all the methods we want to expose and we call the Negate method from the client. 
	
	See the documentation in `net/rpc` for more details.
*/

type Server struct {}

func (this *Server) Negate(i int64, reply *int64) error { 
	*reply = -i 
	return nil 
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return 
	} else {
		fmt.Println("Server runs on 127.0.0.1:9999")
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("Serve, Accept", err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return 
	} else {
		fmt.Println("Client connected to 127.0.0.1:9999")
	}
	
	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) =", result) 
	}
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}