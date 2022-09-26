package model

import (
	"gorm.io/gorm"
)

type Weekbase struct {
	Date  string `json:"date"`
	Value int    `json:"value"`
}
type WeeklistData struct {
	//gorm.Model
	Info  map[string]int `json:"infoData"`
	Date  []string       `json:"date"`
	Value []int          `json:"value"`
	Slim  []Weekbase     `json:"slimData"`
}
type LinksInfo struct {
	gorm.Model
	Name  string `json:"name"`
	Value string `json:"value"`
}
type ETFHistoryList struct {
	Data struct {
		Candle struct {
			Fields      []string    `json:"fields"`
			Five10310SS [][]float64 `json:"510310.SS"`
		} `json:"candle"`
	} `json:"data"`
}
type ETFHistoryData struct {
	MinTime         int64       `json:"min_time" gorm:"min_time;primarykey"`
	OpenPx          float64   `json:"open_px" gorm:"open_px"`
	HighPx          float64   `json:"high_px" gorm:"high_px"`
	LowPx           float64   `json:"low_px" gorm:"low_px"`
	ClosePx         float64   `json:"close_px" gorm:"close_px"`
	BusinessAmount  int       `json:"business_amount" gorm:"business_amount"`
	BusinessBalance int       `json:"business_balance" gorm:"business_balance"`
	PreclosePx      float64   `json:"preclose_px" gorm:"preclose_px"`
	SpecialClosePx  float64   `json:"special_close_px" gorm:"special_close_px"`
	Settlement      float64   `json:"settlement" gorm:"settlement"`
	Amount           int       `json:"amount" gorm:"amount"`
}
//type ETFRealTimeDataInfo struct {
//	hq_type_code,
//	trade_status,
//	prod_name,
//	prod_code,
//	data_timestamp,
//	last_px,
//	px_change,
//	preclose_px,
//	open_px,
//	high_px,
//	low_px,
//	amplitude,
//	up_px,
//	down_px
//}
type ETFRealTimeData struct {
	Data struct {
		Snapshot struct {
			Fields     []string      `json:"fields"`
			Five10310SS []interface{} `json:"510310.SS"`
		} `json:"snapshot"`
	} `json:"data"`
}
