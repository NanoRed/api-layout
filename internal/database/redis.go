package database

import (
	"api-layout/pkg/logger"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Key string

func (k Key) Suffix(v any) Key {
	return Key(fmt.Sprintf("%s%v", k, v))
}

func (k Key) String() string {
	if k[len(k)-1] == '_' {
		logger.Panic("cache key must have a suffix: %s", k)
	}
	return string(k)
}

type Redis struct {
	*redis.Client
}

func NewRedis(dsn string) (rds *Redis, err error) {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		err = fmt.Errorf("can not parse dsn: %s %v", dsn, err)
		return
	}
	rds = &Redis{redis.NewClient(opt)}
	return
}

func (r *Redis) Get(ctx context.Context, key Key) *redis.StringCmd {
	return r.Client.Get(ctx, key.String())
}

func (r *Redis) SetEx(ctx context.Context, key Key, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.Client.SetEx(ctx, key.String(), value, expiration)
}

func (r *Redis) Del(ctx context.Context, keys ...Key) *redis.IntCmd {
	ks := make([]string, len(keys))
	for i, key := range keys {
		ks[i] = key.String()
	}
	return r.Client.Del(ctx, ks...)
}

func (r *Redis) Close() (err error) {
	return r.Client.Close()
}
