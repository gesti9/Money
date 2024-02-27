package server

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	botToken = "6796961656:AAGimXMVJzd0a1JwkFvSEqR28mbMQr2aL1k"
)

func DataBase(s string, chatId int64) {
	c := New(botToken)
	database, err := sql.Open("sqlite3", "users.db")
	Check(err)
	defer database.Close()
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS appeals (id INTEGER PRIMARY KEY, text TEXT, account TEXT)")
	Check(err)
	statement.Exec()

	rows, err := database.Query("SELECT account, text FROM appeals")
	Check(err)

	var id int
	var url1 string
	chek := false

	for rows.Next() {
		rows.Scan(&id, &url1)
		if s == url1 {
			chek = true
		}
	}

	rows.Close()

	statement, err = database.Prepare("INSERT INTO appeals (text) VALUES (?)")
	Check(err)

	coin := 0

	if !chek {
		coin++
		fmt.Println("Новое объявление ", coin)
		c.SendMessage(s, int64(chatId))
		statement.Exec(s)
	}
}

func Check(err error) {

	if err != nil {
		fmt.Println(err)
	}
}
