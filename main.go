package main

import (
  "assignment/initializers"
  "assignment/middlewares"
  "assignment/controllers"

  "fmt"
  "os"

  "github.com/gin-gonic/gin"
)

func init() {

  initializers.LoadEnvs()
  initializers.ConnectDB()

}

func SetupRouter() *gin.Engine {

  gin.SetMode(gin.ReleaseMode)
  engine := gin.Default()
  engine.SetTrustedProxies([]string{fmt.Sprintf(":%s", os.Getenv("PROXY_IP"))})

  authorized := engine.Group("/", middlewares.AuthMiddleWare())
  engine.GET("/patient/search/:id", controllers.PatientSearchV1)
  engine.POST("/staff/create", controllers.StaffCreate)
  engine.POST("/staff/login", controllers.StaffLogin)
  authorized.POST("/patient/search", controllers.PatientSearchV2)

  return engine

}


func main() {

  defer initializers.DisconnectDB()

  engine := SetupRouter()
  engine.Run(fmt.Sprintf(":%s", os.Getenv("ENGINE_PORT")))

}