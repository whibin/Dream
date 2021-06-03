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
	var d []models.Dream
	if err == nil {
		for _, dream := range dreams {
			amount, _ := models.GetLikeAmount(strconv.Itoa(dream.Id))
			dream.Like = amount
			d = append(d, dream)
		}
		return d, true
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

// CountDreamByType 用户梦境类型与其数量
type CountDreamByType struct {
	TypeId int
	Count  int64
}

// CountByDreamType 统计用户梦境类型数量
func CountByDreamType(uId string) ([]CountDreamByType, bool) {
	var cdbts []CountDreamByType
	for i := 0; i < 8; i++ {
		count, err := models.CountByDreamType(uId, strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return nil, false
		}
		c := CountDreamByType{
			TypeId: i,
			Count:  count,
		}
		cdbts = append(cdbts, c)
	}
	return cdbts, true
}

func CountByTime(uid string) ([]dayCount, bool) {
	var dayCounts []dayCount
	for i := 0; i < 6; i++ {
		start := utils.GetFirstDateOfMonth(time.Now().AddDate(0, -i, 0)).Unix()
		end := utils.GetLastDateOfMonth(time.Now().AddDate(0, -i, 0)).Unix() - 1
		count, err := models.CountByTime(start*1000, end*1000, uid)
		if err != nil {
			fmt.Println(err)
			return nil, false
		}
		dayCount := dayCount{
			Day:   time.Now().AddDate(0, -i, 0).Format("1月"),
			Count: count,
		}
		if i == 0 {
			dayCount.Day = "本月"
		}
		if i == 1 {
			dayCount.Day = "上月"
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
	var d2 []models.Dream
	for _, dream := range dreams {
		dream.Nickname = models.GetNickname(dream.Uid)
		amount, _ := models.GetLikeAmount(strconv.Itoa(dream.Id))
		dream.Like = amount
		d2 = append(d2, dream)
	}
	return d2, true
}

func GetDreamByType(t string) ([]models.Dream, bool) {
	dreams, err := models.GetDreamByType(t)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	var d2 []models.Dream
	for _, dream := range dreams {
		dream.Nickname = models.GetNickname(dream.Uid)
		amount, _ := models.GetLikeAmount(strconv.Itoa(dream.Id))
		dream.Like = amount
		d2 = append(d2, dream)
	}
	return d2, true
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
	for _, d := range dreams {
		if d.Id == id {
			dream, _ = models.DreamMatch(d.Dream, uid)
			break
		}
	}
	dream.Nickname = models.GetNickname(dream.Uid)
	amount, _ := models.GetLikeAmount(strconv.Itoa(dream.Id))
	dream.Like = amount
	return dream, true
}
