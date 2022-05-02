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
	url := new(controllers.URLController)

	// set middleware for auth or other things

	// set routes and groups
	router.POST("/url/tinify", url.Tinify)
	router.GET("/:short_id", url.GetURLFromID)

	return router
}
