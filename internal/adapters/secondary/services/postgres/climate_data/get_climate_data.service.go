package services

import (
	models "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/database/models/postgres"
	"github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/secondary/mappers"
	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/climate_data"
	"gorm.io/gorm"
)

type ClimateDataService struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ClimateDataService {
	return &ClimateDataService{
		db: db,
	}
}

func (cds *ClimateDataService) GetLatestClimateData() (climateData *entities.ClimateDataEntity, err error) {
	var model *models.ClimateDataModel
	err = cds.db.
		Model(&models.ClimateDataModel{}).
		Order("created_at desc").
		Limit(1).
		Scan(&model).
		Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	climateData = mappers.MapClimateDataModelToEntity(model)
	return
}
