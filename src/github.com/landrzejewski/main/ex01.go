package main

func each(arr []int, callback func(int, int)){
	for index, value := range arr {
		callback(index, value)
	}
}

func printElement(index, value int) {
	println(index, value)
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

func main() {
	numbers := []int{1,2,3,4,5,6,7,8,9,10}
	each(numbers, printElement)

	each(filter(numbers, isEven), printElement)
}