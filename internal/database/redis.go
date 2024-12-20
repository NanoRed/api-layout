package database

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	*redis.Client
}

func NewRedis(dsn string) (r *Redis, err error) {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		err = fmt.Errorf("can not parse dsn: %s %v", dsn, err)
		return
	}
	r = &Redis{redis.NewClient(opt)}
	return
}

func (r *Redis) Close() (err error) {
	return r.Client.Close()
}
