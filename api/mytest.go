package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"temp/global"
	"temp/lib/utils"
	"temp/model"
	"temp/model/request"
	"temp/model/response"
	"time"
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
func Spider(c *gin.Context) {
	myColly := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"))
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	myColly.AllowedDomains = []string{"baidu.com", "www.baidu.com"}
	// On every a element which has href attribute call callback
	myColly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		//info := model.LinksInfo{Name: e.Text, Value: link }
		//err := global.GVA_DB.Create(&info).Error
		//if err != nil {
		//	global.GVA_LOG.Error("插入失败!", zap.Any("err", err))
		//}
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		myColly.Visit(e.Request.AbsoluteURL(link))
	})
	//https://xueqiu.com/service/v5/stock/screener/screen?category=CN&size=10&order=desc&order_by=deal7d&only_count=0&page=1&_=1663209838933
	// Before making a request print "Visiting ..."
	myColly.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	// Start scraping on https://hackerspaces.org
	myColly.Visit("https://www.baidu.com/")
}

func GETQuoteHistoryData(c *gin.Context){
	neTicker:=time.NewTicker(time.Second*5)
	count := 0
	fmt.Println("当前时间为:", time.Now().Format("2006-01-02 15:04:05"), "count = ", count)
	defer neTicker.Stop() // 需要关闭
	nowTime:= <- neTicker.C
	var actress model.ETFHistoryList
	currenTimeUnix:=time.Now().Unix()
	var res *http.Response
	go func() {
		for {
			//从定时器中获取数据
			nowTime = <-neTicker.C
			count++
			fmt.Println("当前时间为:", time.Unix(nowTime.Unix(), 0).Format("2006-01-02 15:04:05"), "count = ", count)
			//res,_ :=http.Get("https://quotedata.cnfin.com/quote/v1/kline?localDate=1663233097386&get_type=offset&prod_code=510310.SS&candle_period=6&candle_mode=1&fields=open_px%2Chigh_px%2Clow_px%2Cclose_px%2Cbusiness_amount%2Cbusiness_balance%2Cturnover_ratio%2Cpreclose_px%2Cspecial_close_px%2Csettlement%2Camount&data_count=1")
			currenTimeUnixFinal := currenTimeUnix - (24*60*60)*int64(count)
			formatTime:= time.Unix(currenTimeUnixFinal, 0).Format("20060102")
			fmt.Println("当前任务编号为:", formatTime)
			dataUrl:="https://quotedata.cnfin.com/quote/v1/kline?localDate=1663656817701&get_type=offset&prod_code=510310.SS&candle_period=1&candle_mode=1&fields=open_px%2Chigh_px%2Clow_px%2Cclose_px%2Cbusiness_amount%2Cbusiness_balance%2Cturnover_ratio%2Cpreclose_px%2Cspecial_close_px%2Csettlement%2Camount&data_count=240&date=" + formatTime + "&min_time=1501"
			fmt.Println("请求地址:", dataUrl)
			res,_ = http.Get(dataUrl)
			defer res.Body.Close()
			// Example: res,_ :=http.Get("https://quotedata.cnfin.com/quote/v1/kline?localDate=1663656817701&get_type=offset&prod_code=510310.SS&candle_period=1&candle_mode=1&fields=open_px%2Chigh_px%2Clow_px%2Cclose_px%2Cbusiness_amount%2Cbusiness_balance%2Cturnover_ratio%2Cpreclose_px%2Cspecial_close_px%2Csettlement%2Camount&data_count=240&date=20220920&min_time=1501")
			body,_ := ioutil.ReadAll(res.Body)
			err := json.Unmarshal(body,&actress)
			if err != nil {
				fmt.Println(err)
			}
			infoList := make([]model.ETFHistoryData,0,240)
			for i := 0; i <= len(actress.Data.Candle.Five10310SS)-1; i++ {
				item := actress.Data.Candle.Five10310SS[i]
				info := model.ETFHistoryData{
					MinTime:         int64(item[0]),
					OpenPx:          item[1],
					HighPx:          item[2],
					LowPx:           item[3],
					ClosePx:         item[4],
					BusinessAmount:  int(item[5]),
					BusinessBalance: int(item[6]),
					PreclosePx:      item[7],
					SpecialClosePx:  item[8],
					Settlement:      item[9],
					Amount: int(item[10]),
				}
				infoList = append(infoList, info)
			}
			//if errInfo := global.GVA_DB.Create(&infoList).Error;errInfo != nil {
			//	global.GVA_LOG.Error("插入失败!", zap.Any("err", err))
			//}
			if count == 10 {
				neTicker.Stop()
				runtime.Goexit()
			}
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second*1)
			fmt.Println("数据准备中——————————",time.Now().Format("2006-01-02 15:04:05"))
			if count == 10 {
				break
				runtime.Goexit()
			}
		}
	}()
	var total int64
	db := global.GVA_DB.Model(&model.ETFHistoryData{})
	if errTotal := db.Count(&total).Error; errTotal != nil {
		fmt.Println(errTotal)
	}
	response.OkWithMessage("已经开始任务，目前数据量为"+ strconv.FormatInt(total, 10)+",任务天数为 10", c)
}

func GETQuoteTodayData(c *gin.Context)  {
	currenTimeUnix:=time.Now().Unix()
	formatTime:= time.Unix(currenTimeUnix, 0).Format("20060102")
	var actress model.ETFHistoryList
	dataUrl:="https://quotedata.cnfin.com/quote/v1/kline?localDate=1663656817701&get_type=offset&prod_code=510310.SS&candle_period=1&candle_mode=1&fields=open_px%2Chigh_px%2Clow_px%2Cclose_px%2Cbusiness_amount%2Cbusiness_balance%2Cturnover_ratio%2Cpreclose_px%2Cspecial_close_px%2Csettlement%2Camount&data_count=240&date=" + formatTime + "&min_time=1501"
	res,_ :=http.Get(dataUrl)
	defer res.Body.Close()
	body,_ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body,&actress)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(dataUrl)
	/*go func() {
		var item []float64
		var info model.ETFHistoryData
		for i := 0; i <= len(actress.Data.Candle.Five10310SS)-1; i++ {
			item = actress.Data.Candle.Five10310SS[i]
			info = model.ETFHistoryData{
				MinTime:         int64(item[0]),
				OpenPx:          item[1],
				HighPx:          item[2],
				LowPx:           item[3],
				ClosePx:         item[4],
				BusinessAmount:  int(item[5]),
				BusinessBalance: int(item[6]),
				PreclosePx:      item[7],
				SpecialClosePx:  item[8],
				Settlement:      item[9],
				Amount: int(item[10]),
			}
			errInfo := global.GVA_DB.Create(&info).Error
			if errInfo != nil {
				global.GVA_LOG.Error("插入失败!", zap.Any("err", err))
			}
		}
	}()*/
	go func() {
		infoList := make([]model.ETFHistoryData,0,240)
		for i := 0; i <= len(actress.Data.Candle.Five10310SS)-1; i++ {
			item := actress.Data.Candle.Five10310SS[i]
			info := model.ETFHistoryData{
				MinTime:         int64(item[0]),
				OpenPx:          item[1],
				HighPx:          item[2],
				LowPx:           item[3],
				ClosePx:         item[4],
				BusinessAmount:  int(item[5]),
				BusinessBalance: int(item[6]),
				PreclosePx:      item[7],
				SpecialClosePx:  item[8],
				Settlement:      item[9],
				Amount: int(item[10]),
			}
			infoList = append(infoList, info)
		}
		errInfo := global.GVA_DB.Create(&infoList).Error
		if errInfo != nil {
			global.GVA_LOG.Error("插入失败!", zap.Any("err", errInfo))
		}
	}()
	response.OkWithDetailed(actress,"获取到了大批数据",c)
}

func GETQuoteDataList(c *gin.Context) {
	pageQuery := c.Query("page")
	pageSizeQuery := c.Request.URL.Query().Get("pageSize")
	var page, _ = strconv.Atoi(pageQuery)
	var pageSize, _ = strconv.Atoi(pageSizeQuery)
	var pageInfo = request.PageInfo{
		page,
		pageSize,
	}
	//_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	db := global.GVA_DB.Model(&model.ETFHistoryData{})
	var dataList []model.ETFHistoryData
	var total int64
	if errTotal := db.Count(&total).Error; errTotal != nil {
		fmt.Println(errTotal)
	}
	errList := db.Limit(pageInfo.PageSize).Offset(pageInfo.PageSize * (pageInfo.Page - 1)).Order("min_time desc").Find(&dataList).Error
	if errList != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", errList))
		response.FailWithMessage("获取失败", c)
	}else {
		response.OkWithDetailed(response.PageResult{
			List:     dataList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取到了分页数据", c)
	}
}
