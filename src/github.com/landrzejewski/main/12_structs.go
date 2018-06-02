package main

import "fmt"

type money float64

type account struct {
	number  string
	balance int
}

func (a *account) deposit(funds int) {
	a.balance += funds
}

func main() {
	var payment money = 1500.0
	payment = 2000.0
	fmt.Printf("%T\n", payment)

	myAccount := account{"000001", 100}
	otherAccount := account{
		number: "000002",
		balance: 100}

	fmt.Println(myAccount.number, myAccount.balance)
	fmt.Println(otherAccount.number, otherAccount.balance)

	myAccount.deposit(100)
	fmt.Println(myAccount.number, myAccount.balance)

	e := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	fmt.Printf("%+v\n", e)
	fmt.Println("Flag", e.flag)
	fmt.Println("Counter", e.counter)
	fmt.Println("Pi", e.pi)
}
