package models

import (
	. "Dream/database"
)

type Dream struct {
	Id      int    `json:"id"`
	Uid     int    `gorm:"column:u_id" json:"uid"`
	Dream   string `json:"dream"`
	Privacy string `json:"privacy"`
	Time    string `json:"time"`
	Type    int    `json:"type"`
	Like    int    `json:"like"`
	Draw    string `json:"draw"`
	Sound   string `json:"sound"`
	KeyWord string `json:"keyWord" gorm:"column:key_word"`
}

func (Dream) TableName() string {
	return "dream"
}

// SelectOwnDream 查询自己的梦
func SelectOwnDream(uId int) ([]Dream, error) {
	var dreams []Dream
	err := DB.Where("u_id = ?", uId).Find(&dreams).Error
	return dreams, err
}

func Save(dream Dream) error {
	return DB.Create(&dream).Error
}

func CountByDreamType(uId, t string) (int64, error) {
	var count int64
	err := DB.Where("u_id = ? and type = ?", uId, t).Count(&count).Error
	return count, err
}

func CountByTime(start, end string) (int64, error) {
	var count int64
	err := DB.Where("time >= ? and time <= ?", start, end).Count(&count).Error
	return count, err
}
