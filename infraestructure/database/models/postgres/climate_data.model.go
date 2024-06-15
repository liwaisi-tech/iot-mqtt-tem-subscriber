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
	Device *DevicesModel `gorm:"foreignKey:DeviceID;references:ID"`
}

func (rm *ClimateDataModel) BeforeCreate() {
	newID, err := uuid.NewV7()
	if err != nil {
		newID = uuid.New()
	}
	rm.ID = &newID
}
