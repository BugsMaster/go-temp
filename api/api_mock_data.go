package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"reflect"
	"temp/global"
	"temp/model"
	"temp/model/response"
)

func GETChartColor(c *gin.Context) {
	var colorListData []model.SysChartcolor
	err :=global.GVA_DB.Select("color_list").Find(&colorListData).Error
	if err!=nil {
		err.Error()
	}
	colorList := []string{}

	for _,v := range colorListData {
		colorList=append(colorList,v.ColorList)
	}
	response.OkWithDetailed(model.ChartColorList{ColorList: colorList}, "", c)
}
// BatchSave 批量插入数据
/*func batchSave(db *gorm.DB, emps []*model.SysChartcolor) error {
	var buffer bytes.Buffer
	sql := "insert into `sys_chartcolors` (`id`,`age`,`addr`) values"
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	for i, e := range emps {
		if i == len(emps)-1 {
			buffer.WriteString(fmt.Sprintf("('%s','%d',%s);", e.Name, e.Age, e.Addr))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s','%d',%s),", e.Name, e.Age, e.Addr))
		}
	}
	return db.Exec(buffer.String()).Error
}*/

func POSTChartColor(c *gin.Context)  {
	var colorList model.ChartColorList
	_ = c.ShouldBindJSON(&colorList)
	fmt.Println("数据类型是:",reflect.TypeOf(colorList.ColorList))
	var tepModel model.SysChartcolor
	exist:= false
	err := errors.New("")
	for i := 0; i <= len(colorList.ColorList)-1; i++ {
		color := colorList.ColorList[i]
		tepModel = model.SysChartcolor{
			UUID:uuid.NewV4(),
			ColorList:color,
		}
		if !errors.Is(global.GVA_DB.Where("color_list = ?", color).First(&tepModel).Error, gorm.ErrRecordNotFound) {
			exist = true
		}else{
			err = global.GVA_DB.Create(&tepModel).Error
		}
	}

	if exist {
		response.OkWithDetailed(model.ChartColorList{ColorList: colorList.ColorList}, "已经是最新款", c)
	}else {
		if err != nil {
			response.FailWithDetailed(model.ChartColorList{ColorList: colorList.ColorList},"插入失败！", c)
		} else {
			response.OkWithDetailed(model.ChartColorList{ColorList: colorList.ColorList}, "更新成功", c)
		}
	}
}
func POSTCommonColor(c *gin.Context)  {
	//data, _ := global.GVA_REDIS.Get("aaa").Result()
	//response.OkWithMessage("更新成功" + data, c)
	err := global.GVA_REDIS.Set("aaa","第一次存，好紧张",0).Err()
	if err != nil {
		response.FailWithMessage("更新失败", c)
	} else {
		data, _ := global.GVA_REDIS.Get("aaa").Result()
		response.OkWithMessage("更新成功" + data, c)
	}
}