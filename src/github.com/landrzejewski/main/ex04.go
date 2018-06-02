package main

import "fmt"

func main() {
	x := make([]int, 7)
	for i := 0; i < 7; i++ {
		x[i] = i * 100
	}
	twohundred := &x[1]
	x = append(x, 800)
	x[1]++
	fmt.Println("Pointer:", *twohundred, "Element", x[1])
}
