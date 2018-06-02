package main

import "fmt"

func main() {
	state := true
	if value := !state; state && 6 < 100 && !value {
		fmt.Printf(value)
	}
}