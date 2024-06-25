package ports

import entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/climate_data"

type ClimateDataServicePorts interface {
	GetLatestClimateData() (climateData *entities.ClimateDataEntity, err error)
}
