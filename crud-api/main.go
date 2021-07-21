package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// H is a shortcut for map[string]interface{}
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
