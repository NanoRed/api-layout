package model

import (
	"time"

	"gorm.io/gorm"
)

type Example struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;column:id;type:bigint;not null"`
	Code      uint8          `gorm:"column:code;type:smallint;not null;default:0;uniqueIndex:unique_index"`
	Number    uint32         `gorm:"column:number;type:integer;not null;default:0;uniqueIndex:unique_index"`
	Hash      string         `gorm:"column:hash;type:char(32);not null"`
	Varchar   string         `gorm:"column:varchar;type:varchar(255);not null"`
	Text      string         `gorm:"column:text;type:text;not null"`
	Time      Time           `gorm:"column:time;type:timestamp;not null"`
	CreatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index;default:null"` // soft delete
}

func (Example) TableName() string {
	return "example"
}
