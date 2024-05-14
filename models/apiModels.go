package models

import (
    "time"
    "github.com/jinzhu/gorm"
)

// Define the Go structs to represent the JSON data
type Zone struct {
    ID          uint   `gorm:"primary_key" json:"-"`
    ZoneName    string `gorm:"not null" json:"-"`
    ZoneArea    float64 `json:"zone_area"`
    ZoneP       float64 `json:"zone_p"`
    Fertilizer  float64 `json:"fertilizer"`
    // Geometry    Geometry `gorm:"embedded;embedded_prefix:geometry_" json:"geometry"`
    Kmean       []float64 `json:"kmean"`
}

// type Geometry struct {
//     Type        string `gorm:"column:type" json:"type"`
//     Coordinates [][][]float64 `gorm:"column:coordinates" json:"coordinates"`
// }

type Field struct {
    gorm.Model `json:"-"`
    FieldID                 uint `gorm:"primary_key" json:"field_id"`
    ZmapID                  string `gorm:"not null" json:"zmap_id"`
    TypeZmap                string `gorm:"not null" json:"type_zmap"`
    VegetationIndex         string `gorm:"not null" json:"vegetation_index"`
    Date                    time.Time `gorm:"not null" json:"date"`
    Zones                   []Zone `gorm:"foreignkey:FieldID" json:"zones"`
    TotalFertilizerConsumption float64 `gorm:"not null" json:"total_fertilizer_consumption"`
    ImageLink               string `gorm:"not null" json:"image_link"`
}


type ZoneStatus struct{
	Status string `gorm:"not null" json:"status"`
	RequestUrl string `gorm:"not null" json:"request_url"`
}