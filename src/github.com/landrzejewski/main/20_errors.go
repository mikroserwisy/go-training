package main

import (
	"fmt"
	"os"
	"log"
	"math"
)

type MathError struct {
	message string
	code int
}

func (error *MathError) Error() string {
	return fmt.Sprintf("Code: %v, %v", error.code, error.message)
}

func sqrt(number float64) (float64, error)  {
	if number < 0 {
		//return 0, errors.New("Square root of negative number")
		return 0, &MathError{"Square root of negative number", 23}
	}
	return math.Sqrt(number), nil
}

func main() {
	fmt.Println(sqrt(-9.0))

	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		log.Println("Error:", err)
		//log.Fatalln("Error:", err)
		panic(err)
	}
}