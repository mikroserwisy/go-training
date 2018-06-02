package main

import "fmt"

func main() {
	// Slice with a length of 5 elements and a capacity of 8.
	slice := make([]string, 5, 1000)
	slice[0] = "Apple"
	slice[1] = "Orange"
	slice[2] = "Banana"
	slice[3] = "Grape"
	slice[4] = "Plum"

	inspectSlice(slice)

	// Take a slice of slice1. We want just indexes 2 and 3.
	// Parameters are [starting_index : (starting_index + length)]
	slice2 := slice[2:4]
	inspectSlice(slice2)

	fmt.Println("*************************")

	// Change the value of the index 0 of slice2.
	slice2[0] = "CHANGED"

	// Display the change across all existing slices.
	inspectSlice(slice)
	inspectSlice(slice2)
}

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for index, value := range slice {
		fmt.Printf("[%d] %p %s\n", index, &slice[index], value)
	}
}

