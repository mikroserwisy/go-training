package main

import "fmt"

func countDivisors(start, end, div int) int {
	var result int
	if start < 1 {
		start = 1
	}
	for number:=start; number <= end; number++ {
		if number % div == 0 {
			result++
		}
	}
	return result
}

func main()  {
	fmt.Println(countDivisors(0, 100, 3))
}