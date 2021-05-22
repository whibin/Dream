package services

import (
	"Dream/models"
	"fmt"
)

func SaveUser(u models.User) bool {
	c, _ := models.IsOpenIDExist(u.OpenId)
	if c > int64(0) {
		return true
	}
	err := u.Save()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func UpdateUser(u models.User) bool {
	err := u.Update()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func GetUser(openId string) (models.User, bool) {
	user, err := models.GetUser(openId)
	if err != nil {
		fmt.Println(err)
		return models.User{}, false
	}
	return user, true
}
