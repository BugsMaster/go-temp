package model

import "gorm.io/gorm"

type ChartInfo struct {
	ID			uint `json:"id"`
	Name  		string `json:"name"`
	ChartType  	int `json:"chartType"`
	FileName  	string `json:"fileName"`
	FileUrl 	string `json:"fileUrl"`
	IsShare 	bool `json:"isShare"`
	Desc  		string `json:"desc"`
	CoverImg	string `json:"coverImg"`
}
type SysChartInfo struct {
	gorm.Model
	Name  		string 	`json:"name"`
	ChartType  	int 	`json:"chartType"`
	FileName  	string 	`json:"fileName"`
	FileUrl 	string 	`json:"fileUrl"`
	IsShare 	bool 	`json:"isShare"`
	Desc  		string 	`json:"desc"`
	CoverImg	string `json:"coverImg" 	gorm:"default:/img/cloud_red.png;comment:用户ID"`
	AuthorId 	int  `json:"authorId" 		gorm:"default:1;comment:用户ID"`
	AuthorName 	string  `json:"authorName" 	gorm:"default:漆黑小T;comment:作者名称"`
}
type ChartInfoList struct {
	Info []SysChartInfo `json:"info"`
	Total int64  `json:"total"`
}

