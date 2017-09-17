package database

import (
	"fmt"

	kitlog "github.com/go-kit/kit/log"
	"github.com/hashknife/api/config"
	"github.com/jinzhu/gorm"
	// used to access the PostgreSQL driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database
type Database struct {
	conf   *config.Config
	DB     *gorm.DB
	logger kitlog.Logger
}

// NewDatabase
func NewDatabase(c *config.Config, l kitlog.Logger) (*Database, error) {
	var d Database
	d.conf = c
	if err := d.connect(); err != nil {
		return nil, err
	}
	// setup the connection pool
	d.DB.DB().SetMaxIdleConns(10)
	d.DB.DB().SetMaxOpenConns(10)
	d.logger = l
	return &d, nil
}

// connect connects to the database
func (d *Database) connect() error {
	db, err := gorm.Open("postgres", *d.conf.DBURL+"?sslmode=disable&connect_timeout=5")
	if err != nil {
		return fmt.Errorf("Unable to connecto RDS PostgreSQL instance: %s", err)
	}
	db.LogMode(false)
	d.DB = db
	return nil
}
