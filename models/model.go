package models

import (
	. "Dream/database"
)

type Dream struct {
	Id      int    `json:"id"`
	Uid     string `gorm:"column:u_id" json:"uid"`
	Dream   string `json:"dream"`
	Privacy string `json:"privacy"`
	Time    string `json:"time"`
	Type    string `json:"type"`
	Like    string `json:"like"`
	Draw    string `json:"draw"`
	Sound   string `json:"sound"`
	KeyWord string `json:"title" gorm:"column:key_word"`
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
