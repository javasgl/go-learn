package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/set", func(c *gin.Context) {
		for i, cookie := range c.Request.Cookies() {
			fmt.Println(i, "----", cookie.Name, cookie.Domain, cookie.Path)
		}
		secure := false
		httpOnly := true
		c.SetCookie("cookie-name-1", "value-1", 36000, "/", "a.localhost", secure, httpOnly)
		c.SetCookie("cookie-name-1", "value-2", 36000, "/", "localhost", secure, httpOnly)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		for i, cookie := range c.Request.Cookies() {
			fmt.Println(i, "----", cookie.Name, cookie.Domain, cookie.Path)
		}
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
