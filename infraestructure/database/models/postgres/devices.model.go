package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DevicesModel struct {
	ID         *uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	MACAddress string     `gorm:"type:varchar(17);not null;unique"`
	gorm.Model
}

func (rm *DevicesModel) BeforeCreate() {
	newID, err := uuid.NewV7()
	if err != nil {
		newID = uuid.New()
	}
	rm.ID = &newID
}
