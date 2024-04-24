package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

func main() {
	connectionCount := 0
	connectionMap := &sync.Map{}
	server, err := net.Listen("tcp", ":5555")
	fmt.Println("MOGESALMEBI PAPUNITO")
	var logfile *os.File
	filename := "logfile.txt"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("Log file does not exist, creating it")
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating file: ", err.Error())
		}
		logfile = file
		defer file.Close()
	} else {
		fmt.Println("Log file exists, opening it")
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file: ", err.Error())
		}
		logfile = file
		defer file.Close()
	}

	if err != nil {
		fmt.Println("NAXUI BICHO: ", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	for {
		connection, err := server.Accept()

		if err != nil {
			fmt.Println("NAXUI BICHO acc: ", err.Error())
			os.Exit(1)
		}

		connectionMap.Store(strconv.Itoa(connectionCount), connection)
		connectionCount++

		fmt.Println("CONNECTIONS: ", connectionMap)

		go handleConnection(connection, connectionCount-1, connectionMap, logfile)
	}
}

func handleConnection(connection net.Conn, id int, connectionMap *sync.Map, logfile *os.File) {
	defer func() {
		connection.Close()
		connectionMap.Delete(strconv.Itoa(id))
	}()

	for {
		buff, err := bufio.NewReader(connection).ReadString('\n')

		if err != nil {
			fmt.Println("NAXUI BICHOrr: ", err.Error())
			return
		}

		_, err = logfile.WriteString(buff)
		if err != nil {
			fmt.Println("Error writing to file: ", err.Error())
		}

		connectionMap.Range(func(key, value interface{}) bool {
			if conn, ok := value.(net.Conn); ok {
				if conn != connection {
					if _, err := conn.Write([]byte(buff)); err != nil {
						fmt.Println("NAXUI BICHO: ", err.Error())
					}
				}
			}

			return true
		})
	}
}
