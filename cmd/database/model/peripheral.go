package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Peripheral struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    UserId      uuid.UUID `gorm:"unique;not null" json:"userId"`
	User        User      `gorm:"foreignKey:UserId;references:ID"`
	LeftClick   uint64    `gorm:"default:0" json:"leftClick"`
	RightClick  uint64    `gorm:"default:0" json:"rightClick"`
	MouseTravel float64   `gorm:"default:0.0" json:"mouseTravel"`
	Keypress    uint64    `gorm:"default:0" json:"keypress"`
}

func (p *Peripheral) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
