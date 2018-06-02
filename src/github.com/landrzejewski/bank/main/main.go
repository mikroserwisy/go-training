package main

import "github.com/landrzejewski/bank"

func main() {
	repository := bank.MapAccountRepository{Accounts: make(map[string]*bank.Account)}
	generator := bank.InMemoryAccountNumberGenerator{}
	accountService := bank.AccountServiceDefault{Repository: &repository, Generator: &generator}
	loggingProxy := bank.AccountServiceLoggingProxy{Service:&accountService}

	accountNumber := loggingProxy.CreateAccount()
	loggingProxy.DepositFunds(accountNumber, 200)
	loggingProxy.PrintReport()
}
