package bank

import (
	"fmt"
	"sync"
)

type AccountService interface {

	CreateAccount() string

	DepositFunds(number string, funds int) error

	WithdrawFunds(number string, funds int) error

	PrintReport()

}

type DefaultAccountService struct {

	Repository AccountRepository

	Generator AccountNumberGenerator

}

func (accountService *DefaultAccountService) CreateAccount() string {
	accountNumber := accountService.Generator.next()
	accountService.Repository.save(&Account{Number:accountNumber})
	return accountNumber
}

func (accountService *DefaultAccountService) DepositFunds(number string, funds int) error {
	return accountService.process(number, func(account *Account) {account.deposit(funds)})
}

func (accountService *DefaultAccountService) WithdrawFunds(number string, funds int) error {
	return accountService.process(number, func(account *Account) {account.withdraw(funds)})
}

func (accountService *DefaultAccountService) process(number string, callback func(account *Account)) error {
	account, err := accountService.Repository.getByNumber(number)
	if err != nil {
		return err
	}
	callback(account)
	accountService.Repository.update(account)
	return nil
}

func (accountService *DefaultAccountService) PrintReport() {
	accounts, _ := accountService.Repository.getAll()
	for _, account := range accounts {
		fmt.Printf("%v: %v\n", account.Number, account.Balance)
	}
}

type LoggingAccountServiceLogging struct {

	Service AccountService

}

func (service *LoggingAccountServiceLogging) CreateAccount() string {
	return service.Service.CreateAccount()
}

func (service *LoggingAccountServiceLogging) DepositFunds(number string, funds int) error {
	result := service.Service.DepositFunds(number, funds)
	fmt.Printf("%v <- %v\n", number, funds)
	return result
}

func (service *LoggingAccountServiceLogging) WithdrawFunds(number string, funds int) error {
	result := service.Service.WithdrawFunds(number, funds)
	fmt.Printf("%v -> %v\n", number, funds)
	return result
}

func (service *LoggingAccountServiceLogging) PrintReport() {
	service.Service.PrintReport()
}

type AtomicAccountService struct {

	Service AccountService

	Mutex sync.RWMutex

}

func (service *AtomicAccountService) CreateAccount() string {
	service.Mutex.Lock()
	result := service.Service.CreateAccount()
	service.Mutex.Unlock()
	return result
}

func (service *AtomicAccountService) DepositFunds(number string, funds int) error {
	service.Mutex.Lock()
	result := service.Service.DepositFunds(number, funds)
	service.Mutex.Unlock()
	return result
}

func (service *AtomicAccountService) WithdrawFunds(number string, funds int) error {
	service.Mutex.Lock()
	result := service.Service.WithdrawFunds(number, funds)
	service.Mutex.Unlock()
	return result
}

func (service *AtomicAccountService) PrintReport() {
	service.Mutex.RLock()
	service.Service.PrintReport()
	service.Mutex.RUnlock()
}