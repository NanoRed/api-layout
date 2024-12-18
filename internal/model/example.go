package model

import (
	"api-layout/internal/model/common"
	"time"

	"gorm.io/gorm"
)

type Example struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;column:id;type:bigint;not null" json:"id"`
	Code      uint8          `gorm:"column:code;type:smallint;not null;default:0;uniqueIndex:unique_index" json:"code"`
	Number    uint32         `gorm:"column:number;type:integer;not null;default:0;uniqueIndex:unique_index" json:"number"`
	Hash      string         `gorm:"column:hash;type:char(32);not null" json:"hash"`
	Varchar   string         `gorm:"column:varchar;type:varchar(255);not null" json:"varchar"`
	Text      string         `gorm:"column:text;type:text;not null" json:"text"`
	Time      common.Time    `gorm:"column:time;type:timestamp;not null" json:"time"`
	CreatedAt time.Time      `gorm:"index" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;default:null" json:"-"` // soft delete
}

func (Example) TableName() string {
	return "example"
}
