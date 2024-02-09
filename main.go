package main

import (
	"fmt"
	"os"
	"net"
	"io"
)

func main() {
	fmt.Println("Hello World")

	// set connection to listen to incoming requests
	// creates a listener
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error creating listener",err)
		return
	}

	// start receiving requests
	// creates connection , this way multiple connections can be accepted 
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error creating connection", err)
		return
	}

	defer conn.Close() //close  connection when finished

	for {
		buf := make([]byte, 1024)

		//read message from client
		_, err := conn.Read(buf)
		if err !=nil {
			// end of file read
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from client:",err.Error())
			os.Exit(1)
		}

		// ignoring request and snding PONG
		conn.Write([]byte("+OK\r\n"))
	}
}

