package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	routes := r.Group("/")
	receiptsRouter(routes)

	return r
}
