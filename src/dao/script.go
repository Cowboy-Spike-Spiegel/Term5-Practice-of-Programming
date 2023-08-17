package dao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const script_path = "./scripts/"

// script struct
type Step struct {
	S_speak   string            `json:"speak"`
	S_listen  int               `json:"listen"`
	S_branch  map[string]string `json:"branch"`
	S_silence string            `json:"silence"`
	S_default string            `json:"default"`
}
type Script struct {
	Steps map[string]Step `json:"Steps"`
}

// functions for service
func Init_script(filename string) {
	// read script json
	bytes, err := ioutil.ReadFile(script_path + filename)
	if err != nil {
		fmt.Println(err)
	}
	var script Script
	err = json.Unmarshal(bytes, &script)
	if err != nil {
		fmt.Println("____________Dao- Script parse fail!", err)
		return
	}

	// store into steps( map: easy to use )
	steps = script.Steps
}

// (for test currently)
func Return_script_setps() map[string]Step {
	return steps
}
