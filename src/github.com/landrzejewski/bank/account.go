package bank

type Account struct{
	ID int `gorm:"AUTO_INCREMENT"`
	Number string `gorm:"size:26"`
	Balance int
}

func (account *Account) deposit(funds int) {
	account.Balance += funds
}

func (account *Account) withdraw(funds int) {
	account.Balance -= funds
}
