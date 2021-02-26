package routes

import (
  "github.com/gin-gonic/gin"
)

func Build(router *gin.Engine) *gin.Engine {
  router.GET("/", indexHandler)

  return router
}
