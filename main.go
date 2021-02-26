package main

import (
  "net/http"

  "github.com/gin-gonic/gin"

  "github.com/JacobGDG/gin-basic/app/routes"
)

var router *gin.Engine

func main() {
  router = gin.New()

  router.LoadHTMLGlob("public/templates/*")

  routes.Build(router)

  http.ListenAndServe(":8080", router)
}

