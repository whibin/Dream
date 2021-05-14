package services

import (
	"Dream/models"
	"Dream/utils"
	"fmt"
	"time"
)

func SelectOwnDream(uId int) ([]models.Dream, bool) {
	dreams, err := models.SelectOwnDream(uId)
	if err == nil {
		return dreams, true
	}
	fmt.Println(err)
	return nil, false
}

func SaveDream(dream models.Dream) bool {
	// 转成url ----------------
	dream.Draw = utils.LocalPathToUrl(dream.Draw, 1)
	dream.Sound = utils.LocalPathToUrl(dream.Sound, 2)
	// -----------------------
	err := models.Save(dream)
	if err == nil {
		return true
	}
	fmt.Println(err)
	return false
}

func CountByDreamType(uId, t string) (int64, bool) {
	count, err := models.CountByDreamType(uId, t)
	if err != nil {
		fmt.Println(err)
		return 0, false
	}
	return count, true
}

func CountByTime() ([]int64, bool) {
	var counts []int64
	for i := 0; i < 6; i++ {
		start := utils.GetFirstDateOfMonth(time.Now().AddDate(0, -i, 0)).Unix()
		end := utils.GetLastDateOfMonth(time.Now().AddDate(0, -i, 0)).Unix() - 1
		count, err := models.CountByTime(start, end)
		if err != nil {
			fmt.Println(err)
			return nil, false
		}
		counts = append(counts, count)
	}
	return counts, true
}

func DeleteDream(uid, id string) bool {
	err := models.Delete(uid, id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func UpdateDream(dream models.Dream) bool {
	err := dream.Update()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
