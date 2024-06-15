package services

import (
	"context"

	models "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/database/models/postgres"
	"github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/secondary/repositories/mappers"
	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/iot_device"
	"gorm.io/gorm"
)

type IOTDeviceService struct {
	gormDB *gorm.DB
}

func New(gormDB *gorm.DB) *IOTDeviceService {
	return &IOTDeviceService{
		gormDB: gormDB,
	}
}

func (idr *IOTDeviceService) GetIOTDeviceByMACAddress(ctx context.Context, macAddress string) (iotDevice *entities.IOTDeviceEntity, err error) {
	var model *models.IOTDeviceModel
	err = idr.gormDB.
		Model(&models.IOTDeviceModel{}).
		Where("mac_address = ?", macAddress).
		Scan(&model).
		Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	iotDevice = mappers.MapIOTDeviceModelToEntity(model)

	return
}
