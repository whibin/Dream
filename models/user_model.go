package models

import . "Dream/database"

type User struct {
	Id       int    `json:"id"`
	OpenId   string `gorm:"column:open_id" json:"open_id"`
	Nickname string `json:"nickname"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) Save() error {
	return DB.Create(u).Error
}

func (u *User) Update() error {
	return DB.Model(&User{}).Where("id = ?", u.Id).Update("nickname", u.Nickname).Error
}

func GetUser(openId string) (User, error) {
	var user User
	err := DB.Where("open_id = ?", openId).Find(&user).Error
	return user, err
}
