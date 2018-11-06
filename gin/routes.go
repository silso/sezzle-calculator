package gin

import (
	//"github.com/sezzle-calculator/config"
	//"time"

	"github.com/gin-gonic/gin"
	//cors "github.com/itsjamie/gin-cors"

	"github.com/gin-gonic/contrib/static"
	"html"
	"fmt"
)

// InitRoutes : Creates all of the routes for our application and returns a router
func InitRoutes() *gin.Engine {

	router := gin.New()

	//router.Use(gin.Logger())
	//router.Use(gin.Recovery())

	router.Use(static.Serve("/", static.LocalFile("./src", true)))

	router.GET("/", func(c *gin.Context) {
		//c.Header("Content-Type", "application/json")
		//c.JSON(200, hist)
	})
	router.POST("/calculator-post", func(c *gin.Context) {
		mathToClients := c.PostForm("mathToServer")
		c.JSON(
			200,
			gin.H{"mathToClients": html.EscapeString(mathToClients)},
		)
	})

/*
	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:        "*", //cfg.Origins,
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		// ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: true, //Should be true for production. - is more secure because we validate headers as opposed to ember.
	}))

	if config.Debug {
		debugGroup := router.Group("/debug")
		setDebugRoutes(debugGroup)
	}


	v1 := router.Group("/v1")
	{
		setUserRoutes(v1)
	}
*/

	return router
}

func setUserRoutes(g *gin.RouterGroup) {
	g.Use(static.Serve("/", static.LocalFile("./src", true)))
	g.GET("/", func(c *gin.Context) {
		//c.Header("Content-Type", "application/json")
		//c.JSON(200, hist)
	})
	g.POST("/calculator-post", func(c *gin.Context) {
		mathToClients := c.PostForm("mathToServer")
		fmt.Println(mathToClients)
		c.JSON(
			200,
			gin.H{"mathToClients": html.EscapeString(mathToClients)},
		)
	})
}

func setDebugRoutes(g *gin.RouterGroup) {
	g.GET("/test")
}
