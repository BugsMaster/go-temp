package api

import (
	"github.com/gin-gonic/gin"
	"temp/model"
	"temp/model/response"
)

func Weeklist(c *gin.Context) {
	//var l request.Login
	//_ = c.ShouldBindJSON(&l)
	//if err := utils.Verify(l, utils.LoginVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	var weekData = []string{"Jack", "Mark", "Nick"}
	var weekVal = []int{1, 2, 3}
	info:= map[string]int{
		"周一":1,
		"周二":2,
		"周三":3,
		"周四":4,
		"周五":5,
		"周六":6,
		"周日":7,
	}
	response.OkWithDetailed(model.WeeklistData{
		Info: info,
		Date:weekData,
		Value:weekVal,
		Slim: []model.Weekbase{
			{Date:"周一",Value: 1},
			{Date:"周二",Value: 2},
			{Date:"周三",Value: 3},
			{Date:"周四",Value: 4},
			{Date:"周五",Value: 5},
			{Date:"周六",Value: 6},
			{Date:"周日",Value: 7},
		},
	},"获取到了大批数据",c)
}
