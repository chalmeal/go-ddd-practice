package controller

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		appControllers(api)
	}

	master := api.Group("/master")
	{
		masterControllers(master)
	}

	return r
}
