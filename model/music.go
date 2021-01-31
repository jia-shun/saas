package model

import (
	"music-saas/global"
	"time"
)

type Music struct {
	global.MODEL
	UserID       int       `json:"userId" gorm:"not null;index:idx_uid;comment:用户id"`
	SongName     string    `json:"songName" gorm:"not null;comment:歌曲名字"`
	CustomerName string    `json:"customerName" gorm:"not null;comment:客户名字"`
	Price        float64   `json:"price" gorm:"not null;type:decimal(10,2);comment:歌曲制作费用"`
	BeganAt      time.Time `json:"beganAt" gorm:"comment:接手时间"`
	FinishedAt   time.Time `json:"finishedAt" gorm:"comment:完成时间"`
	FinishStatus bool      `json:"finishStatus" gorm:"not null;default:false;comment:是否完成"`
}

func (Music) TableName() string {
	return "music"
}