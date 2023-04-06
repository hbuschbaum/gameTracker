package main

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Page struct {
	ID int
	Title string
	Body []byte
}

func main() {
	fmt.Println("Hello World!")

	db, err := sql.Open("mysql", "gametracker:password@/gametracker")
	if err != nil {
		panic(err)
	}

	res, err := db.Exec("INSERT INTO games (id, name) VALUES (1, 'hello')")
	if err != nil {
		panic(err)
	}
	row, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println(row)
}