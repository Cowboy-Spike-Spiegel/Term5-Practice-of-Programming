package output

import (
	"fmt"
	"os"
)

func Print(info string) {
	if os.Args[1] != "script_test" {
		Log_insert(info + "\n")
		fmt.Println(info)
	}
}
