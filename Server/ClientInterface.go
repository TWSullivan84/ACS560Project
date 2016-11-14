package main

import (
	"net"
	"fmt"
	"log"
)

func main() {
	ln, err := net.Listen("tcp", ":11337")
	if err != nil{
		log.Fatal(err)
	}
	defer ln.Close()
	
	fmt.Println("Server ready!")
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	fmt.Println("A client has connected!")
	conn.Close()
}


