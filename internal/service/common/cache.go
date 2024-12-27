package common

import (
	"api-layout/pkg/logger"
	"strconv"
)

type CacheKey string

const (
	CacheKeyExample CacheKey = "example_"
)

func (k CacheKey) Suffix(v any) (key CacheKey) {
	switch v := v.(type) {
	case string:
		key = k + CacheKey(v)
	case int:
		key = k + CacheKey(strconv.Itoa(v))
	case int8:
		key = k + CacheKey(strconv.FormatInt(int64(v), 10))
	case int16:
		key = k + CacheKey(strconv.FormatInt(int64(v), 10))
	case int32:
		key = k + CacheKey(strconv.FormatInt(int64(v), 10))
	case int64:
		key = k + CacheKey(strconv.FormatInt(int64(v), 10))
	case uint:
		key = k + CacheKey(strconv.FormatUint(uint64(v), 10))
	case uint8:
		key = k + CacheKey(strconv.FormatUint(uint64(v), 10))
	case uint16:
		key = k + CacheKey(strconv.FormatUint(uint64(v), 10))
	case uint32:
		key = k + CacheKey(strconv.FormatUint(uint64(v), 10))
	case uint64:
		key = k + CacheKey(strconv.FormatUint(uint64(v), 10))
	default:
		logger.Panic("unsupported type: %T", v)
	}
	return
}

func (k CacheKey) String() string {
	if k[len(k)-1] == '_' {
		logger.Panic("%s must have a suffix", k)
	}
	return string(k)
}
