package routers

import (
	config "aqaryint/aqaryint/src/configs"
	"aqaryint/aqaryint/src/controllers"

	"github.com/gin-gonic/gin"
)

func Init() {

	router := NewRouter()

	router.Run(config.Appconfig.GetString("server.port"))
}

func NewRouter() *gin.Engine {
	router := gin.New()
	//router.Handler()

	{
		router.POST("/limitedPostMethod", controllers.PostMethod)
		// resource.GET("/GetQueryStringData", controllers.GetQueryStringData) AddUser
		router.POST("/add-user", controllers.AddUser)
	}
	return router
}
