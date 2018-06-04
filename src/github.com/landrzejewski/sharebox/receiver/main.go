package main

import (
	"io"
	"net"
	"os"
	"github.com/landrzejewski/sharebox/common"
	tm "github.com/buger/goterm"
	"fmt"
)

const BUFFER_SIZE = 1024

func createConnection(address string) net.Conn {
	connection, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	return connection
}

func updateProgressBar(progress int) {
	bar := ""
	for i := 0; i < progress / 10; i++ {
		bar += "*"
	}
	bar = common.Fill(bar, 10)
	common.UpdateInfo(fmt.Sprintf("Download progress: %v %v%%", bar, progress))
}

func writeBytesToFile(connection net.Conn, fileSize int64, file *os.File) {
	var receivedBytes int64
	for {
		if (fileSize - receivedBytes) < BUFFER_SIZE {
			io.CopyN(file, connection, fileSize - receivedBytes)
			connection.Read(make([]byte, receivedBytes + BUFFER_SIZE - fileSize))
			break
		}
		io.CopyN(file, connection, BUFFER_SIZE)
		receivedBytes += BUFFER_SIZE
		updateProgressBar(int(100 * receivedBytes / fileSize))
	}
}

func main() {
	connection := createConnection(":9999")
	defer connection.Close()
	fileName, fileSize := common.ReadFileMetadata(connection)
	file := common.CreateFile(fileName)
	defer file.Close()
	tm.Clear()
	writeBytesToFile(connection, fileSize, file)
}
