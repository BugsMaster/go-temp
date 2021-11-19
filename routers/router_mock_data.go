package routers

import (
	"github.com/gin-gonic/gin"
	"temp/api"
)

func InitMockDataRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("")
	{
		BaseRouter.GET("color/chart", api.GETChartColor)
		BaseRouter.POST("color/chart", api.POSTChartColor)
		BaseRouter.POST("color/common", api.CommonColor)
	}
	return BaseRouter
}
