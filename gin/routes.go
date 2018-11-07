package gin

import (
	//"github.com/sezzle-calculator/config"
	//"time"

	"github.com/gin-gonic/gin"
	//cors "github.com/itsjamie/gin-cors"

	"github.com/gin-gonic/contrib/static"
	"io"
	"fmt"
	"github.com/dustin/go-broadcast"
)

const MATH_LOG_SIZE = 10
var mathLog [MATH_LOG_SIZE]string
var mathLogIdx int
var mathLogUpdated bool

var br broadcast.Broadcaster

// InitRoutes : Creates all of the routes for our application and returns a router
func InitRoutes() *gin.Engine {
	mathLogIdx = 0
	mathLogUpdated = false
	br = broadcast.NewBroadcaster(10)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(static.Serve("/", static.LocalFile("./src", true)))

	router.GET("/", handleMainGet)
	router.POST("/calculator-post", handlePost)
	router.POST("/first-post", handleFirstPost)
	router.GET("/stream", handleStream)

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

func handleMainGet(c *gin.Context) {

}

func handlePost(c *gin.Context) {
	mathToClients := c.PostForm("mathToServer")
	if mathLogIdx < MATH_LOG_SIZE {
		mathLog[mathLogIdx] = mathToClients
		mathLogIdx++
	} else {
		//enqueue at last index, dequeue index 0
		mathLogAppended := append(mathLog[1:MATH_LOG_SIZE], mathToClients)
		copy(mathLog[:], mathLogAppended)
	}
	mathLogUpdated = true
	fmt.Println(mathLog)
	br.Submit(mathToClients)
}

func handleFirstPost(c *gin.Context) {
	c.JSON(
		200,
		gin.H{"mathLog": mathLog},
	)
}

func handleStream(c *gin.Context) {
	b := br
	listener := make(chan interface{})
	b.Register(listener)

	defer func() {
		b.Unregister(listener)
		close(listener)
	}()

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-listener; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
	})
}

func setUserRoutes(g *gin.RouterGroup) {
}

func setDebugRoutes(g *gin.RouterGroup) {
	g.GET("/test")
}
