package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"strconv"
	"temp/global"
	"temp/lib/utils"
	"temp/model"
	"temp/model/request"
	"temp/model/response"
)

func PostAddChart(c *gin.Context)  {
	var chartInfo model.ChartInfo
	_ = c.ShouldBindJSON(&chartInfo)
	fmt.Print("数据类型是:",reflect.TypeOf(chartInfo))
	tepModel := model.SysChartInfo{
		Name: chartInfo.Name,
		ChartType: chartInfo.ChartType,
		FileName: chartInfo.FileName,
		FileUrl: chartInfo.FileUrl,
		IsShare: chartInfo.IsShare,
		Desc: chartInfo.Desc,
		CoverImg:chartInfo.CoverImg,
	}
	if !errors.Is(global.GVA_DB.Where("name = ?", chartInfo.Name).First(&tepModel).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage("已经存在", c)
	}else{
		err := global.GVA_DB.Create(&tepModel).Error
		if err != nil {
			response.FailWithMessage("插入失败！", c)
		}else {
			db := global.GVA_DB.Model(&model.SysChartInfo{})
			var infoList []model.SysChartInfo
			var total int64
			err = db.Count(&total).Error
			fmt.Print(total)
			err = db.Find(&infoList).Limit(100).Offset(1).Error
			if err != nil {
				response.FailWithMessage("数据库异常，获取列表失败！", c)
			}else {
				response.OkWithDetailed(model.ChartInfoList{
					Info: infoList,
					Total: total,
				}, "添加成功", c)
			}
		}
	}
}

func GetChartList(c *gin.Context)  {
	page_query := c.Query("page")
	pageSize_query := c.Request.URL.Query().Get("pageSize")
	var page , _= strconv.Atoi(page_query)
	var pageSize, _ = strconv.Atoi(pageSize_query)
	var pageInfo = request.PageInfo{
		page,
		pageSize,
	}
	err := utils.Verify(pageInfo, utils.PageInfoVerify);
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.GVA_DB.Model(&model.SysChartInfo{})
	var infoList []model.SysChartInfo
	var total int64
	err = db.Count(&total).Error
	err = db.Find(&infoList).Limit(limit).Offset(offset).Error
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("数据库异常，获取列表失败！", c)
	}else {
		response.OkWithDetailed(model.ChartInfoList{
			Info: infoList,
			Total: total,
		}, "获取列表成功", c)
	}
}

func GetChartDetail(c *gin.Context)  {
	chartID := c.Param("id")
	var charData model.SysChartInfo
	if !errors.Is(global.GVA_DB.Where("id = ?", chartID).First(&charData).Error, gorm.ErrRecordNotFound) {
	  	global.GVA_DB.Where("id = ?", chartID).First(&charData)
		response.OkWithDetailed(model.ChartInfo{
			ID:charData.ID,
			Name:charData.Name,
			ChartType:charData.ChartType,
			FileName:charData.FileName,
			FileUrl:charData.FileUrl,
			IsShare:charData.IsShare,
			Desc:charData.Desc,
			CoverImg:charData.CoverImg,
		}, "获取成功", c)
	}else{
		response.FailWithMessage("查无此文", c)
	}
}

func PutChartDetail(c *gin.Context) {
	var charData model.ChartInfo
	_ = c.ShouldBindJSON(&charData)

	if errVerify := utils.Verify(charData, utils.ChartInfoVerify); errVerify != nil {
		response.FailWithMessage(errVerify.Error(), c)
		return
	}
	tempData := model.SysChartInfo{
		Name:charData.Name,
		ChartType:charData.ChartType,
		FileName:charData.FileName,
		FileUrl:charData.FileUrl,
		IsShare:charData.IsShare,
		Desc:charData.Desc,
		CoverImg:charData.CoverImg,
	}
	fmt.Print(charData.ChartType)
	err := global.GVA_DB.Where("id = ?", charData.ID).First(&model.SysChartInfo{}).Updates(&tempData).Error
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}
