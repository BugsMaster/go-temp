package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"temp/api"
)


func myMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		fmt.Print("看我这个中间件")
		//response.OkWithMessage("ok,开始爬虫", c)
	}
}

func InitTestRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	//TestRouter := Router.Group("user").Use(middleware.OperationRecord())
	TestRouter := Router.Group("test").Use(myMiddleware())
	//TestRouter := Router.Group("test")
	{
		TestRouter.GET("weeklist", api.Weeklist)
		TestRouter.GET("userlist", api.GetUserList)
		TestRouter.GET("spider", api.Spider)
		TestRouter.GET("quotedata/history", api.GETQuoteHistoryData)
		TestRouter.GET("quotedata/today", api.GETQuoteTodayData)
	}
	return TestRouter
}