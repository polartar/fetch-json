package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/polartar/fetch-test/src/controllers"
)

func receiptsRouter(baseRouter *gin.RouterGroup) {

	receipts := baseRouter.Group("/receipts")
	receipts.GET("/:id/points", controllers.GetReceiptByID)
	receipts.POST("/process", controllers.CreateReceipt)
}
