package bank

import "errors"

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
	repository.Accounts[account.number] = account
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