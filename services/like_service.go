package services

import (
	"Dream/models"
	"fmt"
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
