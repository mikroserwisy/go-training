package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		i++
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}