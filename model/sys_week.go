package model

type Weekbase struct {
	Date string `json:"date"`
	Value int `json:"value"`
}
type WeeklistData struct {
	//gorm.Model
	Info map[string]int `json:"infoData"`
	Date []string `json:"date"`
	Value []int `json:"value"`
	Slim []Weekbase `json:"slimData"`
}
