package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", ":5555")
	fmt.Println("MOGESALMEBI PAPUNITO")

	defer connection.Close()

	if err != nil {
		fmt.Println("NAXUI BICHO: ", err.Error())
		os.Exit(1)
	}

	go io.Copy(os.Stdout, connection)

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		connection.Write([]byte(text))
	}
}
