package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClimateDataModel struct {
	ID          *uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	DeviceID    *uuid.UUID `gorm:"type:uuid;not null"`
	Temperature float64    `gorm:"type:float;not null"`
	Humidity    float64    `gorm:"type:float;not null"`
	gorm.Model

	// Relationships
	Device *IOTDeviceModel `gorm:"foreignKey:DeviceID;references:ID"`
}

func (rm *ClimateDataModel) BeforeCreate(tx *gorm.DB) (err error) {
	newID, err := uuid.NewV7()
	if err != nil {
		return err
	}
	rm.ID = &newID
	return
}

func (rm *ClimateDataModel) TableName() string {
	return "climate_data"
}
