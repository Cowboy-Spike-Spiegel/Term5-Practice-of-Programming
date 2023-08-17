package controller

import (
	"demo/src/dao"
	"demo/src/output"
	"demo/src/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"net/http"
)

// Login page: send to dao and generate response
func Call_login(c *gin.Context) {
	// get login answer from dao layer
	account := c.PostForm("account")
	password := c.PostForm("password")
	output.Print("Controller- recv: account=" + account + ", password=" + password)

	// pre detect for account
	if len(account) != 10 {
		// account len = 10
		c.JSON(http.StatusUnauthorized, gin.H{
			"Identify": "Login fail",
			"account":  account,
			"info":     "Account.length should be equal to 10!",
			"jumpTo":   "/login",
		})
	} else {
		// analyze answer
		name := dao.Login(account, password)
		if name == "" {
			// identify error, return to login
			output.Print("Controller- send: Login fail!")
			c.JSON(http.StatusUnauthorized, gin.H{
				"Identify": "Login fail",
				"account":  account,
				"info":     "Account or password is wrong!",
				"jumpTo":   "/login",
			})
		} else {
			// identify success, send data and jump to new page
			output.Print("Controller- send: Login success!")
			c.JSON(http.StatusOK, gin.H{
				"Identify": "Login success",
				"account":  account,
				"info":     name,
				"jumpTo":   "/service",
			})
		}
	}
}

// Service page: generate msg for return
func Call_service(c *gin.Context) {
	text := c.PostForm("text")
	account := c.PostForm("account")
	output.Print("\nController- recv: " + account + " " + text)

	// get name for current account(if this account exist)
	name := dao.Query(account, "name")
	if name == "" {
		// illegal account
		c.JSON(http.StatusOK, gin.H{"msg": "This account has no record.", "timer": 0})
		return
	} else {
		// get msg & timer from script
		msg, timer := service.Service(account, text)
		output.Print("Controller- send: " + msg)
		c.JSON(http.StatusOK, gin.H{"msg": msg, "timer": timer})
	}
}
