package main

import ( 
	// "github.com/gin-bagonic/gin"
	"net/http"
	"fmt"
	"k8s.io/klog/v2"
)

func Heatlh(appName string) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "Hello from " + appName,
		})
	}
}

func GetEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", health.Health("Main app"))
	return r
}

func main() {
	fmt.Println("Run main application router")
	_ = GetEngine().Run(":8080")
}