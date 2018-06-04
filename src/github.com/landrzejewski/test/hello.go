package main

import (
	"fmt"
	"math"
)

func hi(name string) string {
	return fmt.Sprintf("Hi %v!", name)
}

func absAdd(a, b float64) float64 {
	return math.Abs(a) + math.Abs(b)
}