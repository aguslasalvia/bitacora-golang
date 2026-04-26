package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	// allow cors && specific methods
	r.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"POST", "GET"},
			AllowHeaders: []string{"Content-Type"},
		},
	))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// inport && init routes
	RecordRouter(r)

	return r
}
