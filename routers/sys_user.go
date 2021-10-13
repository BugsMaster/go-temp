package routers

import (
	"github.com/gin-gonic/gin"
	"temp/api"
	"temp/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		UserRouter.POST("register", api.Register)                 // 用户注册账号
		UserRouter.POST("changePassword", api.ChangePassword)     // 用户修改密码
		UserRouter.POST("getUserList", api.GetUserList)           // 分页获取用户列表
		UserRouter.POST("setUserAuthority", api.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("deleteUser", api.DeleteUser)           // 删除用户
		UserRouter.PUT("setUserInfo", api.SetUserInfo)            // 设置用户信息
	}
}
