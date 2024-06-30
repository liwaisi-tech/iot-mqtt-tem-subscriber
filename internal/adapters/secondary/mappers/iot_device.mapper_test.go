package mappers

import (
	"testing"

	"github.com/google/uuid"
	models "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/database/models/postgres"
	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/iot_device"
	"github.com/stretchr/testify/assert"
)

const (
	macTest = "00:11:22:33:44:55"
)

func TestMapIOTDeviceEntityToModel(t *testing.T) {
	t.Run("should return nil when entity is nil", func(t *testing.T) {
		result, err := MapIOTDeviceEntityToModel(nil)

		assert.Nil(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when uuid id is invalid", func(t *testing.T) {
		entity := &entities.IOTDeviceEntity{
			ID: "invalid-uuid",
		}

		result, err := MapIOTDeviceEntityToModel(entity)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return model with valid entity", func(t *testing.T) {
		uuidv7, err := uuid.NewV7()
		assert.Nil(t, err)
		entity := &entities.IOTDeviceEntity{
			ID:         uuidv7.String(),
			MACAddress: macTest,
		}

		result, err := MapIOTDeviceEntityToModel(entity)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, entity.ID, result.ID.String())
		assert.Equal(t, entity.MACAddress, result.MACAddress)
	})
}

func TestMapIOTDeviceModelToEntity(t *testing.T) {
	t.Run("should return nil when model is nil", func(t *testing.T) {
		result := MapIOTDeviceModelToEntity(nil)

		assert.Nil(t, result)
	})

	t.Run("should return entity with valid model and id empty", func(t *testing.T) {
		model := &models.IOTDeviceModel{
			MACAddress: macTest,
		}

		result := MapIOTDeviceModelToEntity(model)

		assert.NotNil(t, result)
		assert.Equal(t, "", result.ID)
		assert.Equal(t, model.MACAddress, result.MACAddress)
	})

	t.Run("should return entity with valid model and id not empty", func(t *testing.T) {
		uuidv7, err := uuid.NewV7()
		assert.Nil(t, err)
		model := &models.IOTDeviceModel{
			ID:         &uuidv7,
			MACAddress: macTest,
		}

		result := MapIOTDeviceModelToEntity(model)

		assert.NotNil(t, result)
		assert.Equal(t, uuidv7.String(), result.ID)
		assert.Equal(t, model.MACAddress, result.MACAddress)
	})

}
