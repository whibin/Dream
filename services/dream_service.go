package services

import (
	"Dream/models"
	"Dream/utils"
	"fmt"
	"strconv"
	"time"
)

type dayCount struct {
	Day   string
	Count int64
}

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

func CountByTime() ([]dayCount, bool) {
	var dayCounts []dayCount
	for i := 0; i < 6; i++ {
		start := utils.GetFirstDateOfMonth(time.Now().AddDate(0, -i, 0)).Unix()
		end := utils.GetLastDateOfMonth(time.Now().AddDate(0, -i, 0)).Unix() - 1
		count, err := models.CountByTime(start, end)
		if err != nil {
			fmt.Println(err)
			return nil, false
		}
		dayCount := dayCount{
			Day:   time.Now().AddDate(0, 0, -i).Format("01/02"),
			Count: count,
		}
		if i == 0 {
			dayCount.Day = "今日"
		}
		if i == 1 {
			dayCount.Day = "昨日"
		}
		dayCounts = append(dayCounts, dayCount)
	}
	return dayCounts, true
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

func GetDreamByTime() ([]models.Dream, bool) {
	dreams, err := models.GetDreamByTime()
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return dreams, true
}

func GetDreamByType(t string) ([]models.Dream, bool) {
	dreams, err := models.GetDreamByType(t)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return dreams, true
}

func CountDreamsByUser(uId string) (int64, bool) {
	count, err := models.CountDreamsByUser(uId)
	if err != nil {
		fmt.Println(err)
		return 0, false
	}
	return count, true
}

func GetReceivedLikes(uid string) (int, bool) {
	i, _ := strconv.Atoi(uid)
	dreams, err := models.SelectOwnDream(i)
	if err != nil {
		fmt.Println(err)
		return 0, false
	}
	counts := 0
	for _, dream := range dreams {
		counts += dream.Like
	}
	return counts, true
}

func DreamMatch(uid, id int) (models.Dream, bool) {
	var dream models.Dream
	dreams, _ := models.SelectOwnDream(uid)
	for _, dream := range dreams {
		if dream.Id == id {
			dream, _ = models.DreamMatch(dream.Dream)
			break
		}
	}
	return dream, true
}
