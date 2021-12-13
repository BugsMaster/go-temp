package routers

import (
	"github.com/gin-gonic/gin"
	"temp/api"
)

func InitUIRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("ui")
	{
		BaseRouter.POST("add/chart", api.PostAddChart)
		BaseRouter.GET("chartlist", api.GetChartList)
		BaseRouter.GET("chart/detail/:id", api.GetChartDetail)
		BaseRouter.PUT("chart/detail", api.PutChartDetail)
	}
	return BaseRouter
}
