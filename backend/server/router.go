package server

import (
	"github.com/andree37/rlld/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// add here the frontend ip
	router.SetTrustedProxies(nil)

	// init controllers
	ping := new(controllers.PingController)

	// set middleware for auth or other things

	// set routes and groups
	router.GET("/ping", ping.Ping)
	router.POST("/ping", ping.DoPing)

	return router
}
