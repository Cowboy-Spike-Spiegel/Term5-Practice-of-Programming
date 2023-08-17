package router

import (
	"demo/src/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateServer() {
	// engine__________________________________________________________________
	// create router engine
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// init assets for router( mainly two pages, discribed below )
	router.LoadHTMLGlob("./templates/*")
	router.StaticFS("/static", http.Dir("./static"))

	// login page
	group_login := router.Group("/login")
	group_login.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "login.html",
		})
	})
	group_login.POST("/", controller.Call_login)

	// service page
	group_service := router.Group("/service")
	group_service.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "service.html", gin.H{
			"title": "service.html",
		})
	})
	group_service.POST("/", controller.Call_service)

	// port
	http.ListenAndServe(":8848", router)
}
