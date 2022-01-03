package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Home Pages",
	})
}
func PostHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Post Home Pages",
	})
}

func main() {

	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/", Home)
	r.POST("/postHome", PostHome)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080
}
