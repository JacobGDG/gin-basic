package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  // create a router object using the "default" one provided by Gin
  r := gin.Default()

  // process html pages at the start so they don't have to be loaded again.
  r.LoadHTMLGlob("templates/*")

  r.GET("/", func(c *gin.Context) {
    c.HTML(
      http.StatusOK,
      "index.html",
      gin.H{
        "title": "Basic Gin!",
      },
    )
  })
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  r.Run()
}
