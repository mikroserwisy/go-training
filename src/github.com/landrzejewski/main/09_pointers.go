package main

import "fmt"

func main() {
	age:=40
	fmt.Println(age)
	fmt.Println(&age)

	var temp = &age
	fmt.Println(temp)
	fmt.Println(*temp)
	*temp = 5
	fmt.Println(age)
}
