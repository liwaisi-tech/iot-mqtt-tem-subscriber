package repositories

import (
	"context"

	"github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/secondary/repositories/mappers"
	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/climate_data"
	"gorm.io/gorm"
)

type ClimateDataRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ClimateDataRepository {
	return &ClimateDataRepository{
		db: db,
	}
}

func (cdr *ClimateDataRepository) CreateClimateData(ctx context.Context, entity *entities.ClimateDataEntity) (newID string, err error) {
	model, err := mappers.MapClimateDataEntityToModel(entity)
	if err != nil {
		return "", err
	}
	if err := cdr.db.Create(&model).Error; err != nil {
		return "", err
	}
	return model.ID.String(), nil
}
