package services

import (
	. "Dream/database"
	"Dream/models"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"sync"
)

func HasLike(userId, dreamId string) (hasLike bool, isError bool) {
	hasLike, err := models.HasLike(userId, dreamId)
	if err != nil {
		fmt.Println(err)
		return true, true
	}
	return hasLike, false
}

func Like(userId, dreamId string) (hasLike bool, isError bool) {
	lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	hasLike, isError = HasLike(userId, dreamId)
	if hasLike {
		return true, isError
	}
	err := models.LikeAdd(userId, dreamId)
	if err != nil {
		fmt.Println(err)
		return hasLike, true
	}
	return false, false
}

func Unlike(userId, dreamId string) (hasLike bool, isError bool) {
	lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	hasLike, isError = HasLike(userId, dreamId)
	if !hasLike {
		return false, isError
	}
	err := models.LikeDelete(userId, dreamId)
	if err != nil {
		fmt.Println(err)
		return hasLike, true
	}
	return true, false
}

func GetLikeAmount(dreamId string) (int, bool) {
	amount, err := models.GetLikeAmount(dreamId)
	if err != nil {
		fmt.Println(err)
		return -1, false
	}
	return amount, true
}

func LikeSave2MySQL() {
	keys, _ := redis.Strings(RedisDB.Do("KEYS", "*"))
	for _, dreamId := range keys {
		_, err := strconv.Atoi(dreamId)
		if err != nil {
			continue
		}
		amount, _ := models.GetLikeAmount(dreamId)
		DB.Table("dream").Where("id = ?", dreamId).Update("likes", amount)
	}
}

func GetDreamByLike() ([]models.Dream, bool) {
	dreams, err := models.GetDreamByLike()
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return dreams, true
}
