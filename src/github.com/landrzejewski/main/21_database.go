package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"

	"log"
	"fmt"
)

func onError(error error)  {
	if error != nil {
		log.Fatal(error)
	}
}



func main() {
	var (
		userId int
		firstName string
		lastName string

	)

	db, _ := sql.Open("mysql", "root:admin@tcp(localhost:3306)/go")
	defer db.Close()

	db.Ping()
	tx, err := db.Begin()

	onError(err)

	stmt, _ := tx.Prepare("insert into users values(null,?,?)")

	response, _ := tx.Stmt(stmt).Exec("Jan", "Kowalski")

	id, err := response.LastInsertId()
	onError(err)

	fmt.Println("Id: ", id)

	err = tx.Rollback()
	onError(err)

	rows, _ := db.Query("select * from users where id = ?", 1)

	for rows.Next() {
		err := rows.Scan(&userId, &firstName, &lastName)
		onError(err)
		log.Println(userId, firstName, lastName)
	}


}