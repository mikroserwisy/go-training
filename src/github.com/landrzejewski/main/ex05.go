package main

import "fmt"

func main() {
	a := []string{"a1", "a2", "a3", "a4", "a5"}
	b := []string{"b1", "b2", "b3", "b4", "b5"}

	for index, value := range a {
		a = b
		fmt.Println(index, value)
	}

	for index, value := range a {
		a = a[:3]
		fmt.Println(index, value )
	}
}
