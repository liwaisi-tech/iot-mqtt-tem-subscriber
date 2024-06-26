package mappers

import (
	"github.com/google/uuid"
	models "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/database/models/postgres"
	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/climate_data"
	"github.com/rs/zerolog/log"
)

func MapClimateDataEntityToModel(entity *entities.ClimateDataEntity) (model *models.ClimateDataModel, err error) {
	if entity == nil {
		return nil, nil
	}
	var id, deviceID uuid.UUID
	if entity.ID != "" {
		id, err = uuid.Parse(entity.ID)
		if err != nil {
			log.Error().Err(err).Msg("error parsing uuid from entity.ID")
			return nil, err
		}
	}
	if entity.DeviceID != "" {
		deviceID, err = uuid.Parse(entity.DeviceID)
		if err != nil {
			log.Error().Err(err).Msg("error parsing uuid from entity.DeviceID")
			return nil, err
		}
	}
	model = &models.ClimateDataModel{
		ID:          &id,
		Temperature: entity.Temperature,
		Humidity:    entity.Humidity,
		DeviceID:    &deviceID,
	}
	return
}

func MapClimateDataModelToEntity(model *models.ClimateDataModel) (entity *entities.ClimateDataEntity) {
	if model == nil {
		return nil
	}
	entity = &entities.ClimateDataEntity{
		ID:          model.ID.String(),
		Temperature: model.Temperature,
		Humidity:    model.Humidity,
		DeviceID:    model.DeviceID.String(),
		CreatedAt:   model.CreatedAt,
	}
	return
}
