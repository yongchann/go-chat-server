package mysql

import (
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb(dsn string) error {
	if db != nil {
		return errors.New("db already initialized")
	}

	gormDb, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	db = gormDb

	return nil
}
