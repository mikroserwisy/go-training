package bank

type Account struct{
	id int
	number string
	balance int
}

func (account *Account) deposit(funds int) {
	account.balance += funds
}


func (account *Account) withdraw(funds int) {
	account.balance -= funds
}
