package service

import (
	"api-layout/internal/model"
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	defaultExpiration time.Duration = time.Minute
)

type Example struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewExample(db *gorm.DB, redis *redis.Client) *Example {
	example := &Example{
		db:    db,
		redis: redis,
	}
	return example
}

func (e *Example) Find(ctx context.Context, id uint64) (data *model.Example, err error) {
	data = &model.Example{}
	key := strconv.FormatUint(id, 10)
	buf, err := e.redis.Get(ctx, key).Bytes()
	if err == nil {
		if err = json.Unmarshal(buf, data); err == nil {
			return
		}
		e.redis.Del(ctx, key)
	}
	result := e.db.WithContext(ctx).
		Where("id = ?", id).
		Find(data)
	if err = result.Error; err != nil {
		return
	}
	if buf, err := json.Marshal(data); err == nil {
		e.redis.SetEx(ctx, key, buf, defaultExpiration)
	}
	return
}

func (e *Example) Save(ctx context.Context, data *model.Example) (err error) {
	tx := e.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()
	err = tx.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "code"},
				{Name: "number"},
			},
			DoUpdates: clause.Assignments(map[string]any{
				"text":       data.Text,
				"created_at": time.Now(),
				"deleted_at": nil,
			}),
		}).
		Where("deleted_at IS NOT NULL").
		Create(data).Error
	return
}
