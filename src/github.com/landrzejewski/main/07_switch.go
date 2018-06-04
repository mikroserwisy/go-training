package main

import "fmt"

type user struct {
	firstName string
	lastName string
}

func detectType(x interface{})  {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case user:
		fmt.Println(x.(user).firstName)
	default:
		fmt.Println("unknown")
	}
}

func main() {
	name := "Jan"
	switch name {
	case "Jan":
		fmt.Println("Jan")
		//fallthrough
	case "Maria":
		fmt.Println("Maria")
	default:
		fmt.Println("Not found")
	}

	switch {
	case len(name) > 3:
		fmt.Println("Valid")
	default:
		fmt.Println("Not found")
	}

	detectType(5)
	var currentUser = user{"Jan", "Kowalski"}
	detectType(currentUser)
}
