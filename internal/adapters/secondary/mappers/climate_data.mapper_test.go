package mappers

import (
	"testing"

	"github.com/google/uuid"
	models "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/database/models/postgres"
	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/climate_data"
	"github.com/stretchr/testify/assert"
)

func TestMapClimateDataEntityToModel(t *testing.T) {
	t.Run("should return nil due to mapping nil entity", func(t *testing.T) {
		result, err := MapClimateDataEntityToModel(nil)

		assert.Nil(t, err)
		assert.Nil(t, result)

	})

	t.Run("should return error due to invalid id uuid", func(t *testing.T) {
		entity := &entities.ClimateDataEntity{
			ID: "invalid-uuid",
		}

		result, err := MapClimateDataEntityToModel(entity)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error due to invalid deviceID uuid", func(t *testing.T) {
		uuidv7, err := uuid.NewV7()
		assert.NoError(t, err)
		entity := &entities.ClimateDataEntity{
			ID:       uuidv7.String(),
			DeviceID: "invalid-uuid",
		}

		result, err := MapClimateDataEntityToModel(entity)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return model with valid entity", func(t *testing.T) {
		uuidv7, err := uuid.NewV7()
		assert.NoError(t, err)
		entity := &entities.ClimateDataEntity{
			ID:          uuidv7.String(),
			Temperature: 25.5,
			Humidity:    50.5,
			DeviceID:    uuidv7.String(),
		}

		result, err := MapClimateDataEntityToModel(entity)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, entity.ID, result.ID.String())
		assert.Equal(t, entity.Temperature, result.Temperature)
		assert.Equal(t, entity.Humidity, result.Humidity)
		assert.Equal(t, entity.DeviceID, result.DeviceID.String())
	})
}

func TestMapClimateDataModelToEntity(t *testing.T) {
	t.Run("should return nil due to mapping nil model", func(t *testing.T) {
		result := MapClimateDataModelToEntity(nil)

		assert.Nil(t, result)
	})

	t.Run("should return entity with valid model", func(t *testing.T) {
		uuidv7, err := uuid.NewV7()
		assert.NoError(t, err)
		model := &models.ClimateDataModel{
			ID:          &uuidv7,
			Temperature: 25.5,
			Humidity:    50.5,
			DeviceID:    &uuidv7,
		}

		result := MapClimateDataModelToEntity(model)

		assert.NotNil(t, result)
		assert.Equal(t, model.ID.String(), result.ID)
		assert.Equal(t, model.Temperature, result.Temperature)
		assert.Equal(t, model.Humidity, result.Humidity)
		assert.Equal(t, model.DeviceID.String(), result.DeviceID)
	})
}
