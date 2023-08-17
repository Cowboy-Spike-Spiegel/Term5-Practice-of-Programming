package service

import (
	"demo/src/dao"
	"demo/src/output"
	"strings"
)

// consts
const NULL = 0 // const timer 0
const symbol_devide = "&"
const symbol_name = "$name"
const symbol_amount = "$amount"
const symbol_balance = "$balance"
const symbol_overtime = "mtc_Overtime"
const state_start = "Start"
const state_exit = "Exit"

// init variable states(map)
func Init() {
	states = make(map[string]string)
}

// clear user(for test currently)
func Clear_User(account string) {
	delete(states, account)
}

// variable for store users' states, and change
var states map[string]string // when get users' post, generate response(msg) depend on this state
func Refresh_state(account string, new_state string) string {
	delete(states, account)
	// Exit, delete record
	if new_state == state_exit {
		return state_exit
	}
	states[account] = new_state
	output.Print("________Service- User states: " + states[account])
	return ""
}

// service for user
func Service(account string, input string) (string, int) {
	// select user state record________________________________________________
	_, ok := states[account]

	// service block___________________________________________________________
	var curStep dao.Step
	msg := ""
	timer := NULL

	// error: no state for this account____________________
	if !ok {
		output.Print("________Service- This account has no state, start serve.")
		states[account] = state_start
		output.Print("________Service- User states: " + states[account])
	}

	// get current step____________________________________
	curStep = dao.Get_step(states[account])
	output.Print("________Service- Current: " + states[account])

	// overtime____________________________________________
	if strings.Contains(input, symbol_overtime) {
		// convert state previously depend on silence
		if Refresh_state(account, curStep.S_silence) == state_exit {
			msg = Msg_append(account, msg, dao.Get_step(state_exit).S_speak)
			return msg, timer
		}
	} else
	// receive_____________________________________________
	{
		// convert state previously depend on branch&input
		flag_find := false
		for k, v := range curStep.S_branch {
			if strings.Contains(input, k) {
				flag_find = true
				if Refresh_state(account, v) == state_exit {
					msg = Msg_append(account, msg, dao.Get_step(state_exit).S_speak)
					return msg, timer
				}
			}
		}
		// don't find in branch, should jump to default
		if !flag_find {
			if Refresh_state(account, curStep.S_default) == state_exit {
				msg = Msg_append(account, msg, dao.Get_step(state_exit).S_speak)
				return msg, timer
			}
		}
	}

	// work until listen exists, stop______________________
	for true {
		// get step and inner attributes(speak and listen)
		curStep = dao.Get_step(states[account])
		msg = Msg_append(account, msg, curStep.S_speak)
		timer = curStep.S_listen
		// if listen is NULL, convert current states
		if timer != NULL {
			break
		}
		// listen is NULL, convert current states
		if Refresh_state(account, curStep.S_default) == state_exit {
			msg = Msg_append(account, msg, dao.Get_step(state_exit).S_speak)
			return msg, timer
		}
	}

	return msg, timer
}

// msg append
func Msg_append(account string, msg string, add string) string {
	// convert add
	output.Print("________Service- Add: " + add)
	if strings.Contains(add, symbol_name) {
		output.Print("________Service- Find msg_add has $name, replace it.")
		add = strings.Replace(add, "$name", dao.Query(account, "name"), -1)
	}
	if strings.Contains(add, symbol_amount) {
		output.Print("________Service- Find msg_add has $name, replace it.")
		add = strings.Replace(add, "$amount", dao.Query(account, "amount"), 1)
	}
	if strings.Contains(add, symbol_balance) {
		output.Print("________Service- Find msg_add has $name, replace it.")
		add = strings.Replace(add, "$balance", dao.Query(account, "balance"), 1)
	}

	if msg == "" {
		msg += add
	} else {
		msg += symbol_devide + add
	}
	return msg
}
