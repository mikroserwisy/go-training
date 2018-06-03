package bank

import (
	"errors"
	"database/sql"
	"github.com/jinzhu/gorm"
)

type AccountRepository interface {

	getByNumber(number string) (*Account, error)

	save(account *Account) error

	update(account *Account) error

	getAll() ([]Account, error)

}

type MapAccountRepository struct {

	Accounts map[string]*Account

}

func (repository *MapAccountRepository) getByNumber(number string) (*Account, error)  {
	account, found := repository.Accounts[number]
	if !found {
		return nil, errors.New("Account not found")
	}
	return account, nil
}

func (repository *MapAccountRepository) save(account *Account) error {
	repository.Accounts[account.Number] = account
	return nil
}

func (repository *MapAccountRepository) update(account *Account) error {
	return repository.save(account)
}

func (repository *MapAccountRepository) getAll() ([]Account, error) {
	accounts := make([]Account, 0, len(repository.Accounts))
	for _, account := range repository.Accounts {
		accounts = append(accounts, *account)
	}
	return accounts, nil
}

type DbAccountsRepository struct {

	Db sql.DB

	Transaction *sql.Tx

}

func (repository *DbAccountsRepository) getByNumber(number string) (*Account, error) {
	rows, err := repository.Db.Query("select * from account where number = ?", number)
	account := Account{}
	if rows.Next() {
		err = rows.Scan(&account.ID, &account.Number, &account.Balance)
	} else {
		return nil, errors.New("Not found")
	}
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (repository *DbAccountsRepository) save(account *Account) error {
	stmt, _ := repository.Db.Prepare("insert into account values(null,?,?)")
	_, err := stmt.Exec(account.Number, account.Balance)
	return err
}

func (repository *DbAccountsRepository) update(account *Account) error {
	stmt, _ := repository.Db.Prepare("update account set balance = ? where id = ?")
	_, err := stmt.Exec(account.Balance, account.ID)
	return err
}

func (repository *DbAccountsRepository) getAll() ([]Account, error)  {
	rows, err := repository.Db.Query("select * from account")
	accounts := make([]Account, 0)
	for rows.Next() {
		account := Account{}
		err = rows.Scan(&account.ID, &account.Number, &account.Balance)
		accounts = append(accounts, account)
	}
	return accounts, err
}

type GormAccountsRepository struct {

	Db *gorm.DB

}

func (repository *GormAccountsRepository) getByNumber(number string) (*Account, error) {
	account := Account{}
	repository.Db.First(&account, "number = ?", number)
	return &account, nil
}

func (repository *GormAccountsRepository) save(account *Account) error {
	repository.Db.Create(&account)
	return nil
}

func (repository *GormAccountsRepository) update(account *Account) error {
	repository.Db.Model(&account).Update("balance", account.Balance)
	return nil
}

func (repository *GormAccountsRepository) getAll() ([]Account, error)  {
	accounts := make([]Account, 0)
	repository.Db.Find(&accounts)
	return accounts, nil
}