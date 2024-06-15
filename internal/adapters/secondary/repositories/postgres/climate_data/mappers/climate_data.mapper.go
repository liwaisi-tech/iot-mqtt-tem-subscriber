package mappers

import (
	"github.com/google/uuid"
	models "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/database/models/postgres"
	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/climate_data"
	"github.com/rs/zerolog/log"
)

func MapClimateDataEntityToModel(entity *entities.ClimateDataEntity) (model *models.ClimateDataModel, err error) {
	id, err := uuid.Parse(entity.ID)
	if err != nil {
		log.Error().Err(err).Msg("error parsing uuid from entity.ID")
		return nil, err
	}
	deviceID, err := uuid.Parse(entity.DeviceID)
	if err != nil {
		log.Error().Err(err).Msg("error parsing uuid from entity.DeviceID")
		return nil, err
	}
	model = &models.ClimateDataModel{
		ID:          &id,
		Temperature: entity.Temperature,
		Humidity:    entity.Humidity,
		DeviceID:    &deviceID,
	}
	return
}
