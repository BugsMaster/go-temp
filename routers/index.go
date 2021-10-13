package routers

import "github.com/gin-gonic/gin"


func IndexRouter(e *gin.Engine)  {
	login(e)
	loadShop(e)
}