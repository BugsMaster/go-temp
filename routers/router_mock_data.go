package routers

import (
	"github.com/gin-gonic/gin"
	"temp/api"
)

func InitMockDataRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("mock")
	{
		BaseRouter.GET("color/chart", api.GETChartColor)
		BaseRouter.POST("color/chart", api.POSTChartColor)
		BaseRouter.POST("color/common", api.POSTCommonColor)
	}
	return BaseRouter
}
