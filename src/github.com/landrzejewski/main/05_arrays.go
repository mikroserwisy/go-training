package main

import "fmt"

func main() {
	var strings [5]string
	strings[0] = "Apple"
	strings[1] = "Orange"
	strings[2] = "Banana"
	strings[3] = "Grape"
	strings[4] = "Plum"

	for index, fruit := range strings {
		fmt.Println(index, fruit)
	}

	numbers := [4]int{10, 20, 30, 40}

	for index := 0; index < len(numbers); index++ {
		fmt.Println(index, numbers[index])
	}

	// Range is using the `five` array directly.
	five := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", five[1])

	for index := range five {
		five[1] = "Jack"
		if index == 1 {
			fmt.Printf("Aft[%s]\n", five[1])
		}
	}

	// Range makes a copy of the `five` array. The v variable is based on the copy.
	five = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", five[1])

	for index, value := range five {
		five[1] = "Jack"
		if index == 1 {
			fmt.Printf("v[%s]\n", value)
		}
	}

	// Range makes a copy of the `five` array's address.
	// The v variable is based on the five array directly.
	five = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", five[1])

	for index, value := range &five {
		five[1] = "Jack"
		if index == 1 {
			fmt.Printf("v[%s]\n", value)
		}
	}
}
