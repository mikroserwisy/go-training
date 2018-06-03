package main

import (
	"github.com/landrzejewski/bank"
	"sync"
	"time"
	"runtime"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var group sync.WaitGroup

func deposit(accountNumber string, accountService bank.AccountService) {
	i := 0
	for i < 100 {
		accountService.DepositFunds(accountNumber, 1000)
		time.Sleep(500)
		i++
	}
	group.Done()
}

func withdraw(accountNumber string, accountService bank.AccountService) {
	i := 0
	for i < 100 {
		accountService.WithdrawFunds(accountNumber, 1000)
		time.Sleep(200)
		i++
	}
	group.Done()
}

func init()  {
	println("Init...")
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	//repository := bank.MapAccountRepository{Accounts: make(map[string]*bank.Account)}
	db, _ := sql.Open("mysql", "root:admin@/go")
	defer db.Close()
	repository := bank.DbAccountsRepository{Db: *db}
	//generator := bank.IncrementalAccountNumberGenerator{}
	generator := bank.IncrementalDbAccountNumberGenerator{Generator:&bank.IncrementalAccountNumberGenerator{}, Db: *db}
	generator.Refresh()
	accountService := bank.AccountServiceDefault{Repository: &repository, Generator: &generator}
	loggingProxyAccountService := bank.AccountServiceLoggingProxy{Service:&accountService}
	//atomicAccountService := bank.AtomicAccountService{Service: &loggingProxyAccountService, Mutex: sync.RWMutex{}}

	accountNumber := loggingProxyAccountService.CreateAccount()

	loggingProxyAccountService.DepositFunds(accountNumber, 1090)

	//group.Add(2)
	//
	//go deposit(accountNumber, &atomicAccountService)
	//go withdraw(accountNumber, &atomicAccountService)
	//
	//group.Wait()

	loggingProxyAccountService.PrintReport()
}
