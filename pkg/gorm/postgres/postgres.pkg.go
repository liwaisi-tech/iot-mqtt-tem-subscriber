package postgres

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	utils "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/utils"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

const Format = "%v, %v"

func newDBConnection() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("PG_DATABASE_HOST"),
		os.Getenv("PG_DATABASE_USER"),
		os.Getenv("PG_DATABASE_PASSWORD"),
		os.Getenv("PG_DATABASE_NAME"),
		os.Getenv("PG_DATABASE_PORT"),
	)
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	maxIdleConns, _ := strconv.Atoi(os.Getenv("GORM_MAX_IDLE_CONNS"))
	sqlDB.SetMaxIdleConns(maxIdleConns)
	maxOpenConns, _ := strconv.Atoi(os.Getenv("GORM_MAX_OPEN_CONNS"))
	sqlDB.SetMaxOpenConns(maxOpenConns)
	maxLifeTime, _ := strconv.Atoi(os.Getenv("GORM_MAX_LIFE_TIME"))
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(maxLifeTime))
	if !utils.IsProduction() {
		db = db.Debug()
	}
	return
}

func NewPostgresDBConnection() *gorm.DB {
	if dbInstance == nil {
		once.Do(func() {
			var err error
			db, err := newDBConnection()
			if err != nil {
				panic(err)
			}
			dbInstance = db
		})
		return dbInstance
	}
	return dbInstance
}
