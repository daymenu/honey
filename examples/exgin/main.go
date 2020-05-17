package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	app.GET("/hello/*name", func(c *gin.Context) {
		c.JSON(200, "hello"+c.Params.ByName("name"))
	})

	app.GET("/hello/*age", func(c *gin.Context) {
		c.JSON(200, "hello"+c.Params.ByName("age"))
	})
	app.Run(":8090")
}
