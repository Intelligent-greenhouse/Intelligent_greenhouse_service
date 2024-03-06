package model

import (
	"github.com/cypunsource/cypunsource-tool/database"
	"time"
)

type Device struct {
	ID       int32  `json:"id" gorm:"colum:id; primaryKey"`
	DeviceId string `json:"deviceId" gorm:"colum:device_id"`

	Co2              float32 `json:"co2" gorm:"colum:co2"`
	LightIntensity   float32 `json:"lightIntensity" gorm:"colum:light_intensity"`
	AirTemperature   float32 `json:"airTemperature" gorm:"colum:air_temperature"`
	AirHumidity      float32 `json:"airHumidity" gorm:"colum:air_humidity"`
	SoilTemperature  float32 `json:"soilTemperature" gorm:"colum:soil_temperature"`
	SoilMoisture     float32 `json:"soilMoisture" gorm:"colum:soil_moisture"`
	SoilConductivity float32 `json:"soilConductivity" gorm:"colum:soil_conductivity"`
	SoilPH           float32 `json:"soilPH" gorm:"colum:soil_ph"`

	Led                 bool `json:"led" gorm:"colum:led"`
	Fan                 bool `json:"fan" gorm:"colum:fan"`
	Water               bool `json:"water" gorm:"colum:water"`
	ChemicalFertilizer  bool `json:"chemicalFertilizer" gorm:"colum:chemical_fertilizer"`
	IncreaseTemperature bool `json:"increaseTemperature" gorm:"colum:increase_temperature"`
	ReduceTemperature   bool `json:"reduceTemperature" gorm:"colum:reduce_temperature"`
	Buzzer              bool `json:"buzzer" gorm:"colum:buzzer"`

	IsActivation bool      `json:"isActivation" gorm:"colum:is_activation"`
	RunTime      time.Time `json:"runtime" gorm:"colum:run_time"`
	Des          string    `json:"des" gorm:"colum:des"`

	database.DeleteModel
	database.MetadataTimeModel
}

func (Device) TableName() string {
	return "device"
}
