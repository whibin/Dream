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

func IsOpenIDExist(openID string) (int64, error) {
	var c int64
	err := DB.Table("user").Where("open_id = ?", openID).Count(&c).Error
	return c, err
}

// GetNickname 根据id获取用户昵称
func GetNickname(uid int) string {
	var s string
	DB.Table("user").Where("id = ?", uid).Select("nickname").Find(&s)
	return s
}
