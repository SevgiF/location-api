package mysql

import (
	"fmt"
	config "github.com/SevgiF/location-api/internal/config/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type GormDBManager struct {
	DB *gorm.DB
}

// GormDBManager, starts the MySQL connection from config variables and returns a GormDBManager.
func NewGormDBManager() *GormDBManager {
	db := initGormDB()
	return &GormDBManager{DB: db}
}

func initGormDB() *gorm.DB {
	user := config.MYSQL_USER
	pass := config.MYSQL_PASS
	addr := config.MYSQL_ADDR
	dbName := config.MYSQL_DB

	// dns for gorm Open
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, addr, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	// Connection pool settings (accessing via sql.DB)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB object from GORM: %v", err)
	}

	sqlDB.SetMaxOpenConns(config.MYSQL_MAX_OPEN_CONNS)
	sqlDB.SetMaxIdleConns(config.MYSQL_MAX_IDLE_CONNS)
	sqlDB.SetConnMaxLifetime(time.Minute * config.MYSQL_CONN_MAX_LIFETIME)
	sqlDB.SetConnMaxIdleTime(time.Minute * config.MYSQL_CONN_MAX_IDLE_TIME)

	//Check if the connection is alive
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Cannot access MySQL database: %v", err)
	}

	return db
}
