package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	myPhrase := "mmm, licorice"
	rdr := strings.NewReader(myPhrase)

	bs, err := ioutil.ReadAll(rdr)
	if err != nil {
		log.Fatalln("my program broke again")
	}

	str := string(bs)
	fmt.Println(str)
}
