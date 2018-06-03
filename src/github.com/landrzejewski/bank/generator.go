package bank

import (
	"fmt"
	"database/sql"
	"github.com/jinzhu/gorm"
	"strconv"
)

type AccountNumberGenerator interface {

	next() string

}

type IncrementalAccountNumberGenerator struct {

	counter int

}

func (generator *IncrementalAccountNumberGenerator) next() string {
	generator.counter++
	return fmt.Sprintf("%026d", generator.counter)
}

type IncrementalDbAccountNumberGenerator struct {

	Generator *IncrementalAccountNumberGenerator

	Db sql.DB

}

func (generator *IncrementalDbAccountNumberGenerator) next() string {
	return generator.Generator.next()
}

func (generator *IncrementalDbAccountNumberGenerator) Refresh()  {
	rows, _ := generator.Db.Query("select max(number) from account")
	if rows.Next() {
		rows.Scan(&generator.Generator.counter)
	}
}

type IncrementalGormAccountNumberGenerator struct {

	Generator *IncrementalAccountNumberGenerator

	Db *gorm.DB

}

func (generator *IncrementalGormAccountNumberGenerator) next() string {
	return generator.Generator.next()
}

func (generator *IncrementalGormAccountNumberGenerator) Refresh()  {
	account := Account{}
	generator.Db.Last(&account)
	generator.Generator.counter, _ = strconv.Atoi(account.Number)
}
