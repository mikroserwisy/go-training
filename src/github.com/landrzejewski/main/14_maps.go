package main

import "fmt"

type user struct {
	name    string
	surname string
}

func main() {
	users := map[string]user{
		"s": user{"Rob", "Roy"},
	}

	users["Roy"] = user{"Rob", "Roy"}
	users["Ford"] = user{"Henry", "Ford"}
	users["Jackson"] = user{"Michael", "Jackson"}
	users["Jacsskson"] = user{"Michael", "Jackson"}

	for key, value := range users {
		fmt.Println(key, value)
	}

	fmt.Println()

	for key := range users {
		fmt.Println(key)
	}

	delete(users, "Roy")

	fmt.Println("=================")
	u, found := users["Roy"]
	fmt.Println("Roy", found, u)
}