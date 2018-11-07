package gin

import (
	//"github.com/sezzle-calculator/config"
	//"time"

	"github.com/gin-gonic/gin"
	//cors "github.com/itsjamie/gin-cors"

	//chris r: added these packages
	"github.com/gin-gonic/contrib/static"
	"io"
	"github.com/dustin/go-broadcast"
)

//count of math history that is kept track of
const MATH_LOG_SIZE = 10
var mathLog [MATH_LOG_SIZE]string
//index of last added math log
var mathLogIdx int

//broadcaster used to send mathLog to clients
var br broadcast.Broadcaster

// InitRoutes : Creates all of the routes for our application and returns a router
func InitRoutes() *gin.Engine {
	//init both of these
	mathLogIdx = 0
	br = broadcast.NewBroadcaster(10)

	router := gin.New()

	/*
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	*/

	//servers the React ./src/index.html as a static page
	//this doesn't feel like the right approach but it works for what I do here
	router.Use(static.Serve("/", static.LocalFile("./src", true)))

	//main GET doesn't do anything
	router.GET("/", handleMainGet)
	//used when calculator presses enter and shares its entry with the server
	router.POST("/calculator-post", handlePost)
	//requested at the start to populate the client's math log
	router.POST("/first-post", handleFirstPost)
	//used to send out 
	router.GET("/stream", handleStream)


	//chris r: could not figure out how to use this middleware
	/*-------------------------------------------------------------------/*
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
	/*-------------------------------------------------------------------*/


	//chris r: wasn't sure what to use the router groups for
	/*-------------------------------------------------------------------/*
	if config.Debug {
		debugGroup := router.Group("/debug")
		setDebugRoutes(debugGroup)
	}

	v1 := router.Group("/v1")
	{
		setUserRoutes(v1)
	}
	/*-------------------------------------------------------------------*/

	return router
}

func handleMainGet(c *gin.Context) {
	//handle maing GET (nothing needed here i don't think)
}

//response to post request when calculations are done
func handlePost(c *gin.Context) {
	mathToClients := c.PostForm("mathToServer")
	//if mathLog isn't full, just append the new math and increment mathLogIdx
	if mathLogIdx < MATH_LOG_SIZE {
		mathLog[mathLogIdx] = mathToClients
		mathLogIdx++
	} else {
		//enqueue at last index, dequeue index 0
		mathLogAppended := append(mathLog[1:MATH_LOG_SIZE], mathToClients)
		copy(mathLog[:], mathLogAppended)
	}
	br.Submit(mathToClients)
}

//response to post request upon initializing the Calculator component in App.js
func handleFirstPost(c *gin.Context) {
	c.JSON(
		200,
		gin.H{"mathLog": mathLog},
	)
}

//based on example here: https://github.com/gin-gonic/gin/tree/master/examples/realtime-advanced
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

/*
func setUserRoutes(g *gin.RouterGroup) {
}

func setDebugRoutes(g *gin.RouterGroup) {
	g.GET("/test")
}
*/
