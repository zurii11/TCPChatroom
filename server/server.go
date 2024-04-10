package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	connectionCount := 0
	connectionList := make([]net.Conn, 5)
	messages := make([]byte, 0, 1024)
	server, err := net.Listen("tcp", ":5555")
	fmt.Println("MOGESALMEBI PAPUNITO")

	if err != nil {
		fmt.Println("NAXUI BICHO: ", err.Error())
		os.Exit(1)
	}

	for {
		connection, err := server.Accept()

		if err != nil {
			fmt.Println("NAXUI BICHO: ", err.Error())
			os.Exit(1)
		}

		connectionList[connectionCount] = connection
		connectionCount++

		fmt.Println("CONNECTIONS: ", connectionList)

		go handleConnection(connection, messages)
	}
}

func handleConnection(connection net.Conn, messages []byte) {
	buff := make([]byte, 128)

	bytesRead, err := connection.Read(buff)

	if err != nil {
		fmt.Println("NAXUI BICHO: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("BYTES READ: ", bytesRead)
	fmt.Println("READ: ", string(buff))

	messages = append(messages, buff...)

	fmt.Println(messages)
	bytesSent, err := connection.Write(messages)

	if err != nil {
		fmt.Println("NAXUI BICHO: ", err.Error())
		os.Exit(1)
	}

	fmt.Println(bytesSent)
}
