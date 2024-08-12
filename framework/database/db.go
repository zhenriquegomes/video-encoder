package database

import (
	"log"

	"github.com/zhenriquegomes/video-encoder/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB            *gorm.DB
	DSN           string
	Debug         bool
	AutoMigrateDB bool
	Env           string
}

func NewDB() *Database {
	return &Database{}
}

func NewDBTest() *gorm.DB {
	DBInstance := NewDB()
	DBInstance.Env = "test"
	DBInstance.DSN = ":memory:"
	DBInstance.AutoMigrateDB = true
	DBInstance.Debug = true

	conn, err := DBInstance.Connect()

	if err != nil {
		log.Fatalf("test db error: %v", err)
	}
	return conn
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error
	d.DB, err = gorm.Open(sqlite.Open(d.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.DB.Debug()
	}

	if d.AutoMigrateDB {
		err := d.DB.AutoMigrate(&domain.Video{}, &domain.Job{})
		if err != nil {
			return nil, err
		}
	}

	return d.DB, nil
}
