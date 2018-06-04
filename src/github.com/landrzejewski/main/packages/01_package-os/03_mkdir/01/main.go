package main

import (
	"log"
	"os"
)

func main() {

	err := os.Mkdir("/somefolderthatdoesntexist", 0x777)
	if err != nil {
		log.Fatalln("my program broke on mkdir: ", err.Error())
	}

	f, err := os.Create("/somefolderthatdoesntexist/hello.txt")
	if err != nil {
		log.Fatalln("my program broke", err.Error())
	}
	defer f.Close()

	str := "Put some phrase here."
	bs := []byte(str)

	_, err = f.Write(bs)
	if err != nil {
		log.Fatalln("error writing to file: ", err.Error())
	}
}