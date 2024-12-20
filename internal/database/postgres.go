package database

import (
	"api-layout/pkg/logger"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/plugin/prometheus"
)

type Postgres struct {
	*gorm.DB
}

func NewPostgres(dsn string) (pg *Postgres, err error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.New(
			logger.Default().Logger,
			gormLogger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  gormLogger.Warn,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			}),
	})
	if err != nil {
		err = fmt.Errorf("can not open database: %v", err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		err = fmt.Errorf("can not get database object: %v", err)
		return
	}
	sqlDB.SetMaxOpenConns(60)
	sqlDB.SetMaxIdleConns(30)
	sqlDB.SetConnMaxIdleTime(time.Minute)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	pg = &Postgres{db}
	return
}

func (p *Postgres) UsePrometheus(dbname, paddr, puser, ppsw string) (err error) {
	return p.Use(prometheus.New(prometheus.Config{
		DBName:          dbname,
		RefreshInterval: 15,
		PushAddr:        paddr,
		PushUser:        puser,
		PushPassword:    ppsw,
	}))
}

func (p *Postgres) Close() (err error) {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return
	}
	return sqlDB.Close()
}
