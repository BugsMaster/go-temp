package routers

import (
	"github.com/gin-gonic/gin"
	"temp/api"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("")
	{
		BaseRouter.GET("captcha", api.Captcha)
		BaseRouter.POST("register", api.Register)
		BaseRouter.POST("login", api.Login)
		BaseRouter.PUT("changePassword", api.ChangePassword)
	}
	return BaseRouter
}
