package migrations

import (
	models "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/database/models/postgres"
	"github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/utils"
	gormpkg "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/pkg/gorm/postgres"
)

func RunMigrations() (err error) {
	gormDB := gormpkg.NewPostgresDBConnection()
	if !utils.IsProduction() {
		gormDB = gormDB.Debug()
	}
	err = gormDB.AutoMigrate(
		&models.DevicesModel{},
		&models.ClimateDataModel{},
	)
	return
}
