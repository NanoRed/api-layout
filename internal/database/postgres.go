package database

import (
	"api-layout/pkg/logger"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/plugin/prometheus"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgres(dsn string) *Postgres {
	pg := &Postgres{}
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
		logger.Fatal("can not open database: %v", err)
	}
	pg.db = db
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("can not get database object: %v", err)
	}
	sqlDB.SetMaxOpenConns(60)
	sqlDB.SetMaxIdleConns(30)
	sqlDB.SetConnMaxIdleTime(time.Minute)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	return pg
}

func (p *Postgres) DB() *gorm.DB {
	return p.db
}

func (p *Postgres) UsePrometheus(name, addr, user, password string) (err error) {
	return p.db.Use(prometheus.New(prometheus.Config{
		DBName:          name,
		RefreshInterval: 15,
		PushAddr:        addr,
		PushUser:        user,
		PushPassword:    password,
	}))
}

func (p *Postgres) Close() (err error) {
	sqlDB, err := p.db.DB()
	if err != nil {
		return
	}
	return sqlDB.Close()
}
