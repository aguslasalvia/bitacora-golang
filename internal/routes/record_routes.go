package routes

import (
	"bitacora/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RecordRouter(router *gin.Engine) {
	recordController := controllers.NewRecordController()
	route := router.Group("/records")
	{
		route.GET("/open", recordController.GetRecordByID)
		route.GET("/machine", recordController.GetRecordByMachine)
		route.POST("/add", recordController.AddRecord)
	}
}
