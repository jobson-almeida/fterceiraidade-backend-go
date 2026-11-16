package entity

type ChartData struct {
	Name  string `json:"name" gorm:"size:255;not null"`
	Value int    `json:"value" gorm:"default:0"`
}
