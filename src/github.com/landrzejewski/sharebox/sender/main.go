package main

import (
	"fmt"
	"net"
	"os"
	"github.com/landrzejewski/sharebox/common"
	"io"
)

const BUFFER_SIZE = 1024

func createServer(address string) net.Listener{
	server, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error listetning: ", err)
		os.Exit(1)
	}
	return server
}

func listenForConnections(server net.Listener) {
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		go sendFileToClient(connection)
	}
}

func sendFileMetadata(connection net.Conn, file *os.File)  {
	fileName, fileSize := common.GetFileMetadata(file)
	connection.Write([]byte(fileSize))
	connection.Write([]byte(fileName))
}

func sendFile(connection net.Conn, file *os.File) {
	sendBuffer := make([]byte, BUFFER_SIZE)
	var sendBytes int64
	for {
		_, err := file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		connection.Write(sendBuffer)
		sendBytes += BUFFER_SIZE
	}
}

func sendFileToClient(connection net.Conn) {
	defer connection.Close()
	file := common.OpenFile(os.Args[1])
	defer file.Close()
	fmt.Println("Sending file...")
	sendFileMetadata(connection, file)
	sendFile(connection, file)
	fmt.Println("File has been sent, closing connection!")
	return
}

func main() {
	server := createServer(":9999")
	defer server.Close()
	fmt.Println("Server started! Waiting for connections...")
	listenForConnections(server)
}



