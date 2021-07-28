package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB_USER     = "root"
	DB_PASSWORD = "admin"
	DB_NAME     = "classicmodels"
	DB_DRIVER   = "mysql"
)

var sqlDB *sql.DB

var gormDB *gorm.DB

func SQLDB() *sql.DB {
	return sqlDB
}

func GormDB() *gorm.DB {
	return gormDB
}

func InitDBConnection() {
	connStr := fmt.Sprintf("%s:%s@/%s", DB_USER, DB_PASSWORD, DB_NAME)
	sqldb, err := sql.Open(DB_DRIVER, connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database, exiting: %v", err)
	}
	sqlDB = sqldb
	//Configure Database connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	gormDatabase, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to Gorm, exiting: %v", err)
	}
	gormDB = gormDatabase
}
