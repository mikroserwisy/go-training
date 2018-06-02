package bank

import "fmt"

type AccountNumberGenerator interface {

	next() string

}

type InMemoryAccountNumberGenerator struct {

	counter int

}

func (generator *InMemoryAccountNumberGenerator) next() string {
	generator.counter++
	return fmt.Sprintf("%026d", generator.counter)
}