package database

import (
	"api-layout/pkg/logger"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rdb *redis.Client
}

func NewRedis(dsn string) *Redis {
	rds := &Redis{}
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		logger.Fatal("can not parse dsn: %s %v", dsn, err)
	}
	rdb := redis.NewClient(opt)
	rds.rdb = rdb
	return rds
}

func (r *Redis) DB() *redis.Client {
	return r.rdb
}

func (r *Redis) Close() (err error) {
	return r.rdb.Close()
}
