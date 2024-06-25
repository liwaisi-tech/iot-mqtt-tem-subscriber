package mappers

import (
	"github.com/google/uuid"
	models "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/database/models/postgres"
	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/iot_device"
)

func MapIOTDeviceEntityToModel(entity *entities.IOTDeviceEntity) (model *models.IOTDeviceModel, err error) {
	var entityID uuid.UUID
	if entity.ID != "" {
		entityID, err = uuid.Parse(entity.ID)
		if err != nil {
			return nil, err
		}
	}

	model = &models.IOTDeviceModel{
		ID:         &entityID,
		MACAddress: entity.MACAddress,
	}
	return
}

func MapIOTDeviceModelToEntity(model *models.IOTDeviceModel) (entity *entities.IOTDeviceEntity) {
	if model == nil {
		return nil
	}
	var id string
	if model.ID != nil {
		id = model.ID.String()
	}
	entity = &entities.IOTDeviceEntity{
		ID:         id,
		MACAddress: model.MACAddress,
	}
	return
}
