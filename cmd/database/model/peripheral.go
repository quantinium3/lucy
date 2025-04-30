package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Peripheral struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid" json:"primaryKey"`
	UserId      uuid.UUID `gorm:"unique;not null"`
	User        User      `gorm:"foreignKey:UserID;references:ID"`
	LeftClick   uint64    `gorm:"default:0" json:"leftClick"`
	RightClick  uint64    `gorm:"default:0" json:"rightClick"`
	MouseTravel float64   `gorm:"default:0.0" json:"mouseTravel"`
	Keypress    uint64    `gorm:"default:0" json:"keypress"`
}

func (p *Peripheral) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
