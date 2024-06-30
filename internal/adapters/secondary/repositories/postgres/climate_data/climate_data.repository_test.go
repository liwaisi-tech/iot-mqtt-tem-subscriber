package repositories

import (
	"context"
	"testing"

	"github.com/google/uuid"
	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/climate_data"
	"github.com/liwaisi-tech/iot-mqtt-tem-subscriber/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateClimateData(t *testing.T) {
	gormDB, dbMock := mocks.NewPostgresGormDB(t)
	repository := New(gormDB)
	uuidID, err := uuid.NewV7()
	assert.Nil(t, err)
	entity := &entities.ClimateDataEntity{
		ID: uuidID.String(),
	}
	ctx := context.TODO()

	t.Run("should fails due to error invalid uuid", func(t *testing.T) {
		entity.ID = "invalid-uuid"
		defer func() {
			entity.ID = uuidID.String()
		}()
		_, err := repository.CreateClimateData(ctx, entity)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "invalid UUID length")
	})

	t.Run("should fails due to error creating climate data", func(t *testing.T) {
		dbMock.ExpectBegin()
		dbMock.ExpectQuery(`INSERT INTO "climate_data" \("id","device_id","temperature","humidity","created_at","updated_at","deleted_at"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6,\$7\) RETURNING "id"`).
			WillReturnError(assert.AnError)
		dbMock.ExpectRollback()
		_, err := repository.CreateClimateData(nil, nil)

		assert.NotNil(t, err)
	})

	t.Run("should create climate data", func(t *testing.T) {
		dbMock.ExpectBegin()
		dbMock.ExpectQuery(`INSERT INTO "climate_data" \("id","device_id","temperature","humidity","created_at","updated_at","deleted_at"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6,\$7\) RETURNING "id"`).
			WillReturnRows(dbMock.NewRows([]string{"id"}).AddRow(uuidID))
		dbMock.ExpectCommit()
		newID, err := repository.CreateClimateData(ctx, entity)

		assert.Nil(t, err)
		assert.Equal(t, uuidID.String(), newID)
	})
}
