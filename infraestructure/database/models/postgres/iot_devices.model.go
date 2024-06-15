package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IOTDeviceModel struct {
	ID         *uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	MACAddress string     `gorm:"type:varchar(17);not null;unique"`
	gorm.Model
}

func (rm *IOTDeviceModel) BeforeCreate(tx *gorm.DB) (err error) {
	newID, err := uuid.NewV7()
	if err != nil {
		return err
	}
	rm.ID = &newID
	return
}

func (rm *IOTDeviceModel) TableName() string {
	return "iot_device"
}
