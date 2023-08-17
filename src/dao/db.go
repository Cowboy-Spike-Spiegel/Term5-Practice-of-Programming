package dao

import (
	"database/sql"
	"fmt"
	"log"
)

// table struction
type USER struct {
	account  string  `db:"account"`
	password string  `db:"password"`
	name     string  `db:"name"`
	amount   float32 `db:"change"`
	balance  float32 `db:"change"`
}

// database init
func Init_db() {
	var err error

	// link
	db, err = sql.Open("mysql", "user_practice_of_programming:password@tcp(127.0.0.1:3306)/project_practice_of_programming")
	if err != nil {
		fmt.Println(err)
		fmt.Println("____________Dao- Create sql.DB fail.")
	}

	// detect link state
	err = db.Ping()
	if err != nil {
		fmt.Println("____________Dao- Database link fail.")
		log.Fatal(err)
	} else {
		fmt.Println("____________Dao- Database link successfully.")
	}
}

func Close_db() {
	db.Close()
}
