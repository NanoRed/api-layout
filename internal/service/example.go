package service

import (
	"api-layout/internal/database"
	"context"
)

type Example struct {
	pg  *database.Postgres
	rds *database.Redis
}

func NewExample(pg *database.Postgres, rds *database.Redis) *Example {
	example := &Example{
		pg:  pg,
		rds: rds,
	}
	return example
}

func (e *Example) Find(ctx context.Context, id uint64) {

}
