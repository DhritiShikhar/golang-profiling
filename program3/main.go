package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	r := newServer()
	r.Run(":8081")
}

func newServer() *gin.Engine {
	r := gin.Default()
	r.GET("/api/status", func(c *gin.Context) { 
		c.Status(200) 
	})
	return r
}
