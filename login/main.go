package main

import (
	"login/controllers"
	"login/initializers"
	"login/middlewares"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main(){
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/signin", controllers.Signin)
	r.POST("/signincookie", controllers.SigninCookie)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	r.Run()
}