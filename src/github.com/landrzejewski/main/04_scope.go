package main

import "fmt"

func sayHello(name string) string {
	return "Hello " + name
}

func main()  {
	x:=36
	fmt.Println(x)
	{
		// bad practice
		x:=2
		fmt.Println(x)
		y:=22
		fmt.Println(y)
	}
	// fmt.Println(y)


	// bad practice
	fmt.Printf("%T\n", sayHello)
	sayHello:=sayHello("Jan")
	fmt.Println(sayHello)
}
