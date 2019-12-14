package mysqldriver

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var (
	_database = os.Getenv("DATABASE")
	_dbHost   = os.Getenv("DB_HOST")
)

type (
	DriverWithGorm struct {
		db *gorm.DB
	}
)

func NewDriverWithGorm() (*DriverWithGorm, error) {
	driver := &DriverWithGorm{}
	if err := driver.init(); err != nil {
		return nil, err
	}
	return driver, nil
}

func (d *DriverWithGorm) init() error {
	var err error
	log.Printf("info: config -> %v, %v", _database, _dbHost)
	d.db, err = gorm.Open(_database, _dbHost)
	if err != nil {
		return err
	}
	log.Printf("info: establised db connection\n")
	return nil
}
