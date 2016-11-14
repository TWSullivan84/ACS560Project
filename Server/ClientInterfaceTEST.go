package main

import (
	"net"
	"log"
)

func main(){
	conn, err := net.Dial("tcp", "ec2-35-163-106-205.us-west-2.compute.amazonaws.com:11337")
	if err != nil{
		log.Fatal(err)
	}
	
	conn.Close()
}