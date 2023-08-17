package main

import (
	"demo/src/dao"
	"demo/src/output"
	"demo/src/router"
	"demo/src/service"
	"demo/src/test"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// main -----------------------------------------------------------------------
func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "script_test" {
			// Test mode: local test all scripts
			fmt.Println("#####SCRIPT_TEST MODE")
			test.Test_scripts()
		} else {
			// init____________________________________________________________________
			output.Log_init()
			// init serveice layer data(states)
			output.Print("Init service layer.")
			service.Init()
			// init dao layer data(script and db)
			output.Print("Init dao layer.")
			dao.Init_db()
			dao.Init_script(os.Args[1])

			// Release mode: create router layer
			fmt.Println("#####RELEASE MODE")
			router.CreateServer()
		}
	}
}
