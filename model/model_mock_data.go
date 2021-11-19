package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ChartColorList struct {
	ColorList  []string `json:"colorList"`
}
type SysChartcolor struct {
	gorm.Model
	UUID   uuid.UUID    `json:"uuid" gorm:"comment:颜色UUID"`
	ColorList  string `json:"colorList" gorm:"comment:颜色列表"`
	Author	string  `json:"author" gorm:"default:漆黑小T"`
}