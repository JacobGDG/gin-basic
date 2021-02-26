package routes

import (
  "net/http"

  "github.com/gin-gonic/gin"
)


func indexHandler(c *gin.Context) {
  c.HTML(
    http.StatusOK,
    "index.html",
    gin.H{
      "title": "Basic Gin 2!",
    },
  )
}
