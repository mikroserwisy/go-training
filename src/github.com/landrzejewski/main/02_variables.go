package main

import "fmt"

// package scope
var fullName = "Jan Kowalski"
var firstName, lastName string = "Marek", "Nowak"
var age int

func main() {
	age = 40

	// function scope
	var score = 90
	text := "Go languge"  // var text = "Go language"
	name, version := "Go", 1.8

	fmt.Println(fullName)
	fmt.Println(firstName, lastName, age)
	fmt.Printf("%v, %v, %v, %v", score, text, name, version)
}
