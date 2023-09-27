package api

import (
	"swap/api/middlewares"
	"swap/api/routes"
	"swap/api/validate"
	"swap/db"
	"github.com/gin-gonic/gin"
)

func Start() {

	//init mysql
	db.InitMysql()

	//gin bind go-playground-validator
	validate.BindingValidator()

	// gin start
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	app.Use(middlewares.Cors()) // 「 Cross domain Middleware 」
	routes.InitRoute(app)
	_ = app.Run(":8000")

}

/*
 If you change the version, you need to modify the following files'
 config/init.go
*/
