package main

import "fmt"

func main() {
	var data []string
	lastCap := cap(data)
	for record := 1; record <= 100000; record++ {
		data = append(data, fmt.Sprintf("Rec: %d", record))
		if lastCap != cap(data) {
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100
			lastCap = cap(data)
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n", &data[0], record, cap(data), capChg)
		}
	}
}