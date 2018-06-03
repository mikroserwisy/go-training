package main

import (
	"github.com/landrzejewski/bank"
	"sync"
	"time"
	"runtime"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

func initDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:admin@/go")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&bank.Account{})
	return db
}

func main() {
	//db, _ := sql.Open("mysql", "root:admin@/go")
	db := initDb()
	db.LogMode(true)
	defer db.Close()

	//repository := bank.MapAccountRepository{Accounts: make(map[string]*bank.Account)}
	//repository := bank.DbAccountsRepository{Db: *db}
	repository := bank.GormAccountsRepository{Db: db}

	//generator := bank.IncrementalAccountNumberGenerator{}
	//generator := bank.DbAccountNumberGenerator{Generator:&bank.IncrementalAccountNumberGenerator{}, Db: *db}
	generator := bank.GormAccountNumberGenerator{Generator:&bank.IncrementalAccountNumberGenerator{}, Db: db}
	generator.Refresh()

	accountService := bank.DefaultAccountService{Repository: &repository, Generator: &generator}
	loggingAccountService := bank.LoggingAccountServiceLogging{Service:&accountService}
	//atomicAccountService := bank.AtomicAccountService{Service: &loggingProxyAccountService, Mutex: sync.RWMutex{}}

	accountNumber := loggingAccountService.CreateAccount()
	loggingAccountService.DepositFunds(accountNumber, 1000)
	loggingAccountService.PrintReport()

	//group.Add(2)
	//go deposit(accountNumber, &atomicAccountService)
	//go withdraw(accountNumber, &atomicAccountService)
	//group.Wait()
}