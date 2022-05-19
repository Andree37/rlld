package server

import (
	"fmt"
	"github.com/andree37/rlld/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"os"
)

func SetupRouter() *gin.Engine {

	frontend := fmt.Sprintf("%s://%s:%s", os.Getenv("FRONTEND_PROTOCOL"), os.Getenv("FRONTEND_IP"), os.Getenv("FRONTEND_PORT"))

	router := gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		return nil
	}
	// add here the frontend ip
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{frontend}

	router.Use(cors.New(corsConfig))
	// init controllers
	url := new(controllers.URLController)

	// url for serving static frontend files
	router.Use(static.Serve("/", static.LocalFile("./build", true)))

	router.NoRoute(func(c *gin.Context) {
		c.File("./build/index.html")
	})

	// set routes and groups
	api := router.Group("api")
	api.POST("/tinify", url.Tinify)
	api.GET("/:short_id", url.GetURLFromID)

	return router
}
