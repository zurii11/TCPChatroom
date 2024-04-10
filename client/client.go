package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", ":5555")
	fmt.Println("MOGESALMEBI PAPUNITO")

	if err != nil {
		fmt.Println("NAXUI BICHO: ", err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ENTER MSG: ")
	text, _ := reader.ReadString('\n')

	bytesSent, err := connection.Write([]byte(text))

	if err != nil {
		fmt.Println("NAXUI BICHO: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("BYTES SENT: ", bytesSent)

	buff := make([]byte, 1024)

	bytesRead, err := connection.Read(buff)

	fmt.Println("MSG: ", string(buff[:bytesRead]))
	connection.Close()
}
