package models

import . "Dream/database"

type Chat struct {
	Id           int    `json:"id"`
	MainId       int    `json:"main_id"`
	SendId       int    `json:"send_id"`
	Content      string `json:"content"`
	DreamId      int    `gorm:"column:d_id" json:"dream_id"`
	MainNickname string `gorm:"-" json:"main_nickname"`
	SendNickname string `gorm:"-" json:"send_nickname"`
}

func (Chat) TableName() string {
	return "chat"
}

func (c *Chat) AddComment() error {
	return DB.Create(c).Error
}

func DeleteComment(id string) error {
	return DB.Table("chat").Where("id = ?", id).Delete(&Chat{}).Error
}

func GetCommentsByDream(id string) ([]Chat, error) {
	var chats []Chat
	err := DB.Table("chat").Where("d_id = ?", id).Find(&chats).Error
	return chats, err
}
