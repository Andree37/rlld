package server

import (
	"os"

	"github.com/andree37/rlld/controllers"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	// mode of the release
	gin.SetMode(os.Getenv("RELEASE_MODE"))

	router := gin.Default()

	err := router.SetTrustedProxies(nil)
	if err != nil {
		return nil
	}

	router.Use(CORSMiddleware())
	// init controllers
	url := new(controllers.URLController)

	// set routes and groups
	api := router.Group("api")
	api.POST("/tinify", url.Tinify)
	api.GET("/:short_id", url.GetURLFromID)

	return router
}
