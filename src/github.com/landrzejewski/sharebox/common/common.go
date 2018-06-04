package common

import (
	"os"
	"net"
	"strconv"
	"strings"
	tm "github.com/buger/goterm"
)

const FILL_SYMBOL  = "-"

func Fill(text string, toLength int) string {
	for {
		length := len(text)
		if length < toLength {
			text = text + FILL_SYMBOL
			continue
		}
		break
	}
	return text
}

func CreateFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

func OpenFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

func ReadFileMetadata(connection net.Conn) (string, int64) {
	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)
	connection.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
	connection.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName), FILL_SYMBOL)
	return fileName, fileSize
}

func GetFileMetadata(file *os.File) (string, string) {
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fileSize := Fill(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := Fill(fileInfo.Name(), 64)
	return fileName, fileSize
}

func UpdateInfo(text string)  {
	tm.MoveCursor(0, 0)
	tm.Print(tm.Background(tm.Color(tm.Bold(text), tm.MAGENTA), tm.WHITE))
	tm.Flush()
}