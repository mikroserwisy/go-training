package main

import "fmt"

func simpleSayHello(name string) string {
	return "Hello " + name
}

func swap(text, otherText string) (string, string) {
	return otherText, text
}

func Name() (name string) {
	name = "Go"
	return
}

func execute(fn func())  {
	fn()
}

func create() func() {
	return func() {
		fmt.Println("Go...")
	}
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x - 1)
}

func main() {
	fmt.Println(simpleSayHello("Jan"))
	var otherText, text = swap("text", "otherText")
	fmt.Println(otherText, text)

	greet := func() {
		fmt.Println("Hello")
	}

	greet()
	execute(greet)

	goFn := create()
	goFn()

	fmt.Println(factorial(3))

	var result = func() string {
		fmt.Println("Done")
		return "result"
	}()

	fmt.Println(result)
}