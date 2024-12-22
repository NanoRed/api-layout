package model

import (
	"api-layout/internal/model/common"
	"time"

	"gorm.io/gorm"
)

type Example struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;column:id;type:bigint;not null" json:"id" example:"10001"`
	Code      uint8          `gorm:"column:code;type:smallint;not null;default:0;uniqueIndex:unique_index" json:"code" example:"1"`
	Number    uint32         `gorm:"column:number;type:integer;not null;default:0;uniqueIndex:unique_index" json:"number" example:"1"`
	Hash      string         `gorm:"column:hash;type:char(32);not null" json:"hash" example:"a8306c3367a271d245ba59fe4a8f9eaf"`
	Varchar   string         `gorm:"column:varchar;type:varchar(255);not null" json:"varchar" example:"varchar"`
	Text      string         `gorm:"column:text;type:text;not null" json:"text" example:"text"`
	Time      common.Time    `gorm:"column:time;type:timestamp;not null" json:"time" example:"2021-01-01T00:00:00Z"`
	CreatedAt time.Time      `gorm:"index" json:"created_at" example:"2021-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `gorm:"index;default:null" json:"-"` // soft delete
}

func (Example) TableName() string {
	return "example"
}
