package main

import "fmt"

func each(arr []int, callback func(int, int))  {
	for index, value := range arr {
		callback(index, value)
	}
}

func filter(arr []int, predicate func(int) bool) []int {
	result := make([]int, 0)
	each(arr, func(index, value int) {
		if predicate(value) {
			result = append(result, value)
		}
	})
	return result
}

func isEven(element int) bool {
	return element % 2 == 0
}

func avg(numbers ...float64) float64  {
	sum := 0.0
	for _, value := range numbers {
		sum += value
	}
	return sum / float64(len(numbers))
}

func print(index, value int) {
	fmt.Printf("%v: %v\n", index, value)
}

func main() {
	numbers := []int{1,3,5,7,8} // slice
	fmt.Printf("%T", numbers)

	each(numbers, print)
	var result = filter(numbers, isEven)
	fmt.Println(numbers)
	fmt.Println(result)
	fmt.Println(avg(1,2.2))
	fmt.Println(avg([]float64{1.2,44.4}...))
}