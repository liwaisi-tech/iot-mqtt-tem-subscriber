package repositories

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/iot_device"
	"github.com/liwaisi-tech/iot-mqtt-tem-subscriber/mocks"
)

func TestCreateIOTDevice(t *testing.T) {
	gormDB, dbMock := mocks.NewPostgresGormDB(t)
	repository := New(gormDB)
	uuidID, err := uuid.NewV7()
	assert.Nil(t, err)
	ctx := context.TODO()
	entity := &entities.IOTDeviceEntity{
		ID: uuidID.String(),
	}
	t.Run("should fails due to error invalid uuid", func(t *testing.T) {
		entity.ID = "invalid-uuid"
		defer func() {
			entity.ID = uuidID.String()
		}()
		_, err := repository.CreateIOTDevice(ctx, entity)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "invalid UUID length")
	})

	t.Run("should fails due to error creating iot device", func(t *testing.T) {
		dbMock.ExpectBegin()
		dbMock.ExpectQuery(`INSERT INTO "iot_device" \("id","mac_address","created_at","updated_at","deleted_at"\) VALUES \(\$1,\$2,\$3,\$4,\$5\) RETURNING "id"`).
			WillReturnError(assert.AnError)
		dbMock.ExpectRollback()
		_, err := repository.CreateIOTDevice(ctx, entity)

		assert.NotNil(t, err)
	})

	t.Run("should create iot device", func(t *testing.T) {
		entity := &entities.IOTDeviceEntity{
			MACAddress: "mac-address",
		}
		dbMock.ExpectBegin()
		dbMock.ExpectQuery(`INSERT INTO "iot_device" \("id","mac_address","created_at","updated_at","deleted_at"\) VALUES \(\$1,\$2,\$3,\$4,\$5\) RETURNING "id"`).
			WillReturnRows(dbMock.NewRows([]string{"id"}).AddRow(uuidID))
		dbMock.ExpectCommit()
		newID, err := repository.CreateIOTDevice(ctx, entity)

		assert.Nil(t, err)
		assert.Equal(t, uuidID.String(), newID)
	})
}
