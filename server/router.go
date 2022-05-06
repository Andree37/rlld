package server

import (
	"fmt"

	"github.com/andree37/rlld/config"
	"github.com/andree37/rlld/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	c := config.Getconfig()
	frontend := fmt.Sprintf("%s://%s:%d", c.GetString("frontend.protocol"), c.GetString("frontend.ip"), c.GetInt("frontend.port"))

	router := gin.Default()
	router.SetTrustedProxies(nil)
	// add here the frontend ip
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{frontend}

	router.Use(cors.New(corsConfig))
	// init controllers
	url := new(controllers.URLController)

	// set middleware for auth or other things

	// set routes and groups
	router.POST("/url/tinify", url.Tinify)
	router.GET("/:short_id", url.GetURLFromID)

	return router
}
