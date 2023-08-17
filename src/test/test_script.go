package test

import (
	"bufio"
	"demo/src/dao"
	"demo/src/output"
	"demo/src/service"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const script_path = "./scripts"
const prefixe_input = "./scripts/test/input/input_"
const prefixe_reference = "./scripts/test/reference/reference_"
const test_account = "2020211376"
const symbol_Start = "Start"
const symbol_Exit = "Exit"
const symbol_End = "End"

func Test_scripts() {
	// init
	service.Init()

	// init dao layer data(script and db)
	output.Print("Init dao layer.")
	dao.Init_db()

	// detect all scripts
	files, err := ioutil.ReadDir(script_path)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range files {
		// select all scripts
		if !item.IsDir() {
			// load this script into dao.script
			fmt.Println("\nTest:", item.Name())
			dao.Init_script(item.Name())

			// test logic & response
			if Test_script_logic() {
				fmt.Println("Logic correct!")
			} else {
				fmt.Println("Logic error!")
			}
			if Test_script_response(item.Name()) {
				fmt.Println("Response correct!")
			} else {
				fmt.Println("Response error!")
			}
		}
	}
}

func Test_script_logic() bool {
	// get current script steps
	steps := dao.Return_script_setps()

	// judge Start
	if _, ok := steps[symbol_Start]; !ok {
		fmt.Println("No Start step")
		return false
	}
	// judge Exit
	if _, ok := steps[symbol_Exit]; !ok {
		fmt.Println("No Exit step")
		return false
	}
	for k, value := range steps {
		// judge listen
		if value.S_listen < 0 {
			fmt.Println(k, "Listen wrong")
			return false
		}
		// judge branch
		for _, v := range value.S_branch {
			if _, ok := steps[v]; !ok {
				fmt.Println(k, "Branch wrong")
				return false
			}
		}
		// judge silence
		if _, ok := steps[value.S_silence]; value.S_listen > 0 && !ok {
			fmt.Println(k, "Silence wrong")
			return false
		}
		// judge default
		if _, ok := steps[value.S_default]; value.S_listen == 0 && !ok && value.S_default != symbol_End {
			fmt.Println(k, "Default wrong")
			return false
		}
	}
	return true
}

func Test_script_response(script_name string) bool {
	// read reference
	data, err := ioutil.ReadFile(prefixe_reference + strings.Split(script_name, ".")[0] + ".txt")
	if err != nil {
		fmt.Println("Open test reference file error!", err)
	}
	test := string(data)

	// read in lines & generate response into ans
	ans := ""
	input_file, err := os.OpenFile(prefixe_input+strings.Split(script_name, ".")[0]+".txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open input file error!", err)
	}
	defer input_file.Close()
	br := bufio.NewReader(input_file)
	for {
		// read into line
		line, err := br.ReadString('\n')
		line = strings.Replace(line, "\n", "", -1)
		if err != nil {
			break
		}
		// send this line to service layer and generate string-ans
		msg, _ := service.Service(test_account, line)
		ans += msg + "\n"
	}
	service.Clear_User(test_account)

	return ans == test
}
