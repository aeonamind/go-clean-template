package database

import (
	"os"

	"github.com/koujiman/gobox/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	cfg := config.NewConfig()
	dsn := cfg.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		os.Exit(1)
	}

	return db
}
