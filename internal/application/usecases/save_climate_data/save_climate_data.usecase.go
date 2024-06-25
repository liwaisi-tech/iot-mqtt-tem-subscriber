package usecases

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"

	eventmsgdtos "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/dtos/events"
	climatedataentities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/climate_data"
	iotdeviceentities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/iot_device"
	climatedataerr "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/errors/climate_data"
	climatedataportrepo "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/ports/repositories/climate_data"
	iotdeviceportrepo "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/ports/repositories/iot_device"
	iotdeviceportsrv "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/ports/services/iot_device"
)

type SaveClimateDataUseCase struct {
	climateDataRepo  climatedataportrepo.ClimateDataRepositoryPort
	iotDeviceRepo    iotdeviceportrepo.IOTDeviceRepositoryPort
	iotDeviceService iotdeviceportsrv.IOTDeviceServicePort
}

func New(
	climateDataRepo climatedataportrepo.ClimateDataRepositoryPort,
	iotDeviceRepo iotdeviceportrepo.IOTDeviceRepositoryPort,
	iotDeviceService iotdeviceportsrv.IOTDeviceServicePort,
) *SaveClimateDataUseCase {
	return &SaveClimateDataUseCase{
		climateDataRepo:  climateDataRepo,
		iotDeviceRepo:    iotDeviceRepo,
		iotDeviceService: iotDeviceService,
	}
}

func (usecase *SaveClimateDataUseCase) Execute(ctx context.Context, eventMessage []byte) (err error) {
	dto, err := usecase.getDTOFromEventMessage(eventMessage)
	if err != nil {
		return
	}

	climateDataEntity, err := usecase.getEntityFromDTO(dto)
	if err != nil {
		return
	}

	deviceEntity, err := usecase.getDeviceEntity(ctx, dto.MACAddress)
	if err != nil {
		return
	}

	climateDataEntity.DeviceID = deviceEntity.ID

	id, err := usecase.climateDataRepo.CreateClimateData(ctx, climateDataEntity)
	if err != nil {
		log.Error().Err(err).Msg("failed to create climate data")
		return
	}
	log.Info().Str("id", id).Msg("climate data created")
	return
}

func (usecase *SaveClimateDataUseCase) getDTOFromEventMessage(eventMessage []byte) (
	dto *eventmsgdtos.IOTMessageEventDTO,
	err error,
) {
	err = json.Unmarshal(eventMessage, &dto)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal event message")
		return nil, err
	}
	return
}

func (usecase *SaveClimateDataUseCase) getEntityFromDTO(dto *eventmsgdtos.IOTMessageEventDTO) (
	entity *climatedataentities.ClimateDataEntity,
	err error,
) {
	if dto == nil {
		return nil, climatedataerr.ErrIOTDeviceInformationNotFound
	}
	entity = &climatedataentities.ClimateDataEntity{
		Temperature: dto.Temperature,
		Humidity:    dto.Humidity,
	}
	return
}

func (usecase *SaveClimateDataUseCase) getDeviceEntity(
	ctx context.Context,
	macAddress string,
) (
	entity *iotdeviceentities.IOTDeviceEntity,
	err error,
) {
	entity, err = usecase.iotDeviceService.GetIOTDeviceByMACAddress(ctx, macAddress)
	if err != nil {
		log.Error().Err(err).Msg("failed to get iot device by mac address")
		return nil, err
	}
	if entity == nil {
		entity = &iotdeviceentities.IOTDeviceEntity{
			MACAddress: macAddress,
		}
		id, err := usecase.iotDeviceRepo.CreateIOTDevice(ctx, entity)
		if err != nil {
			log.Error().Err(err).Msg("failed to create iot device")
			return nil, err
		}
		entity.ID = id
	}
	return
}
