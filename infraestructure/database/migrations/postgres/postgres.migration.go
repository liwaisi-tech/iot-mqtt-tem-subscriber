package migrations

import (
	models "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/database/models/postgres"
	gormpkg "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/pkg/gorm/postgres"
)

func RunMigrations() (err error) {
	gormDB := gormpkg.NewPostgresDBConnection()

	err = gormDB.AutoMigrate(
		&models.IOTDeviceModel{},
		&models.ClimateDataModel{},
	)
	return
}
