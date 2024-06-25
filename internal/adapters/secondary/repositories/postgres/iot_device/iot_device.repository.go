package repositories

import (
	"context"

	"gorm.io/gorm"

	models "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/database/models/postgres"
	mappers "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/secondary/mappers"
	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/iot_device"
)

type IOTDeviceRepository struct {
	gormDB *gorm.DB
}

func New(gormDB *gorm.DB) *IOTDeviceRepository {
	return &IOTDeviceRepository{
		gormDB: gormDB,
	}
}

func (idr *IOTDeviceRepository) CreateIOTDevice(ctx context.Context, entity *entities.IOTDeviceEntity) (newID string, err error) {
	model, err := mappers.MapIOTDeviceEntityToModel(entity)
	if err != nil {
		return "", err
	}
	response := idr.gormDB.
		Model(&models.IOTDeviceModel{}).
		Create(&model)

	if response.Error != nil {
		return "", response.Error
	}
	newID = model.ID.String()
	return
}
