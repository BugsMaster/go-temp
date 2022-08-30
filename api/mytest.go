package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
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
func Spider(c *gin.Context) {
	myColly := colly.NewCollector()
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	myColly.AllowedDomains = []string{"hackerspaces.org", "wiki.hackerspaces.org"}
	// On every a element which has href attribute call callback
	myColly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		myColly.Visit(e.Request.AbsoluteURL(link))
		response.OkWithMessage(e.Text, c)
	})
	// Before making a request print "Visiting ..."
	myColly.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	// Start scraping on https://hackerspaces.org
	myColly.Visit("https://hackerspaces.org/")
}
