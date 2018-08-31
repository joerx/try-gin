package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// User bla
type User struct {
	Name   string `json:"name" binding:"required"`
	Handle string `json:"handle" binding:"required"`
}

func main() {
	r := gin.New()

	r.Use(func(c *gin.Context) {
		c.Next() // handle request
		log.Printf("%s %s - %d", c.Request.Method, c.Request.URL, c.Writer.Status())
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("hello %s", name),
		})
	})

	r.GET("/names", func(c *gin.Context) {
		names := c.QueryArray("name")
		c.JSON(200, gin.H{
			"names": names,
		})
	})

	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(400, gin.H{
				"error": fmt.Errorf("%s", err),
			})
			return
		}
		c.JSON(200, gin.H{
			"status": "OK",
			"user":   user,
		})
	})

	r.Static("/static", "./static")

	r.Run("localhost:9000")
}
