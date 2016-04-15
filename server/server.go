package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		//Loop waiting for incomming call
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	//resrve buffer to store data
	buffLength := 1024
	buf := make([]byte, buffLength)
	//Read incomming Data
	//cont := true
	tcpString := ""
	//for cont {
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading in Buffer:", err.Error())
	}
	//cast the input to string
	tcpString += string(buf[:reqLen])
	fmt.Println("Received:", tcpString, reqLen)
	//cont = false
	//TODO think about good way to handle longer messages
	//}
	fmt.Println("End of loop", tcpString)
	handleDatabaseRequest(tcpString)
	fmt.Println("Eof")
	conn.Close()
}
