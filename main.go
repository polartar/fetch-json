package main

import (
	"github.com/polartar/fetch-test/src/middlewares"
	"github.com/polartar/fetch-test/src/routes"
	docs "github.com/polartar/fetch-test/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	docs.SwaggerInfo.BasePath = "/"
	
	router := routes.SetupRoutes()
	middlewares.RegisterMiddlewares(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":5000")

}
