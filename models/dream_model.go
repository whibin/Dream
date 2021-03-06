package models

import (
	. "Dream/database"
)

type Dream struct {
	Id       int    `json:"id"`
	Uid      int    `gorm:"column:u_id" json:"uid"`
	Dream    string `json:"dream"`
	Privacy  string `json:"privacy"`
	Time     int64  `json:"time"`
	Type     int    `json:"type"`
	Like     int    `gorm:"column:likes" json:"like"`
	Draw     string `json:"draw"`
	Sound    string `json:"sound"`
	KeyWord  string `json:"keyWord" gorm:"column:key_word"`
	Nickname string `gorm:"-" json:"nickname"`
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
	err := DB.Table("dream").Where("u_id = ? and type = ?", uId, t).Count(&count).Error
	return count, err
}

func CountByTime(start, end int64, uid string) (int64, error) {
	var count int64
	err := DB.Table("dream").Where("time between ? and ? and u_id = ?", start, end, uid).Count(&count).Error
	return count, err
}

func Delete(uid, id string) error {
	return DB.Where("id = ? and u_id = ?", id, uid).Delete(&Dream{}).Error
}

func (d *Dream) Update() error {
	return DB.Model(&Dream{Id: d.Id, Uid: d.Uid}).Updates(d).Error
}

func GetDreamByTime() ([]Dream, error) {
	var dreams []Dream
	err := DB.Table("dream").Where("privacy = 'n'").Order("time desc").Find(&dreams).Error
	return dreams, err
}

func GetDreamByType(t string) ([]Dream, error) {
	var dreams []Dream
	err := DB.Table("dream").Where("type = ? and privacy = 'n'", t).Order("time desc").Find(&dreams).Error
	return dreams, err
}

func CountDreamsByUser(uId string) (int64, error) {
	var count int64
	err := DB.Table("dream").Where("u_id = ?", uId).Count(&count).Error
	return count, err
}

func DreamMatch(d string, uid int) (Dream, error) {
	var dream Dream
	sql := "select * from dream where match(dream) against(?) and u_id != ?"
	err := DB.Table("dream").Raw(sql, d, uid).First(&dream).Error
	return dream, err
}
