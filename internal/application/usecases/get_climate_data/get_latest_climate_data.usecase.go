package usecases

import (
	"context"

	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/climate_data"
	climatedataportsrv "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/ports/services/climate"
)

type GetLatestClimateDataUseCase struct {
	climateDataSrv climatedataportsrv.ClimateDataServicePorts
}

func New(
	climateDataSrv climatedataportsrv.ClimateDataServicePorts,
) *GetLatestClimateDataUseCase {
	return &GetLatestClimateDataUseCase{
		climateDataSrv: climateDataSrv,
	}
}

func (uc *GetLatestClimateDataUseCase) Execute(ctx context.Context) (climateData *entities.ClimateDataEntity, err error) {
	return uc.climateDataSrv.GetLatestClimateData()
}
