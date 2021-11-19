package initialize

import (
	"github.com/gin-gonic/gin"
	"temp/global"
	"temp/middleware"
	"temp/routers"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	//Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	//global.GVA_LOG.Info(global.GVA_CONFIG.Local.Path)
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		routers.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		routers.InitTestRouter(PublicGroup) // 自动初始化相关
		routers.InitMockDataRouter(PublicGroup) // mockdata相关
	}
	/*PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		routers.InitApiRouter(PrivateGroup)                   // 注册功能api路由
		//routers.InitJwtRouter(PrivateGroup)                   // jwt相关路由
		routers.InitUserRouter(PrivateGroup)                  // 注册用户路由
	}*/
	global.GVA_LOG.Info("router register success")
	return Router
}
