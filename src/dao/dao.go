package dao

import (
	"database/sql"
	"demo/src/output"
	"fmt"
)

// Database block______________________________________________________________
var db *sql.DB

// Login for user
func Login(account string, password string) string {
	// query all users
	rows, err := db.Query("select * from users where account='" + account + "' and password='" + password + "'")
	output.Print("____________Dao- Login check: " + account + " " + password)
	if err != nil {
		fmt.Println(err)
	}

	// judge login state
	if rows.Next() {
		// store a instance into user
		var user USER
		err = rows.Scan(&user.account, &user.password, &user.name, &user.amount, &user.balance)
		if err != nil {
			fmt.Println(err)
		}
		return user.name
	}
	return ""
}

// query for user
func Query(account string, query string) string {
	// query
	rows, err := db.Query("select " + query + " from users where account='" + account + "'")
	output.Print("____________Dao- Select " + query + " from users where account='" + account + "'")
	if err != nil {
		fmt.Println(err)
	}

	ans := ""
	// store ans
	if rows.Next() {
		// store a instance into user
		var tmp string
		err = rows.Scan(&(tmp))
		if err != nil {
			fmt.Println(err)
		}
		if ans == "" {
			ans += tmp
		} else {
			ans += ("&" + tmp)
		}
	}

	output.Print("____________Dao- Query " + query + ":" + ans)
	return ans
}

// update for user
func Update(account string, update string) string {
	// no use for this
	return ""
}

// Script block________________________________________________________________
var steps map[string]Step

// get step depend on state
func Get_step(state string) Step {
	output.Print("____________Dao- Get step for " + state)
	return steps[state]
}
