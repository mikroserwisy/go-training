package bank

import "fmt"

type AccountService interface {

	CreateAccount() string

	DepositFunds(number string, funds int) error

	WithdrawFunds(number string, funds int) error

	PrintReport()

}

type AccountServiceDefault struct {

	Repository AccountRepository
	Generator AccountNumberGenerator

}

func (accountService *AccountServiceDefault) CreateAccount() string {
	accountNumber := accountService.Generator.next()
	accountService.Repository.save(&Account{number:accountNumber})
	return accountNumber
}

func (accountService *AccountServiceDefault) DepositFunds(number string, funds int) error {
	return accountService.process(number, func(account *Account) {account.deposit(funds)})
}

func (accountService *AccountServiceDefault) WithdrawFunds(number string, funds int) error {
	return accountService.process(number, func(account *Account) {account.withdraw(funds)})
}

func (accountService *AccountServiceDefault) process(number string, callback func(account *Account)) error {
	account, err := accountService.Repository.getByNumber(number)
	if err != nil {
		return err
	}
	callback(account)
	accountService.Repository.update(account)
	return nil
}

func (accountService *AccountServiceDefault) PrintReport() {
	accounts, _ := accountService.Repository.getAll()
	for _, account := range accounts {
		fmt.Printf("%v: %v\n", account.number, account.balance)
	}
}

type AccountServiceLoggingProxy struct {

	Service AccountService

}

func (service *AccountServiceLoggingProxy) CreateAccount() string {
	return service.Service.CreateAccount()
}

func (service *AccountServiceLoggingProxy) DepositFunds(number string, funds int) error {
	result := service.Service.DepositFunds(number, funds)
	fmt.Printf("Transaction: %v <- %v\n", number, funds)
	return result
}

func (service *AccountServiceLoggingProxy) WithdrawFunds(number string, funds int) error {
	result := service.Service.WithdrawFunds(number, funds)
	fmt.Printf("Transaction: %v -> %v\n", number, funds)
	return result
}

func (service *AccountServiceLoggingProxy) PrintReport() {
	service.Service.PrintReport()
}