package models

import (
	. "Dream/database"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

func LikeAdd(userId, dreamId string) error {
	_, err := RedisDB.Do("SADD", dreamId, userId)
	return err
}

func LikeDelete(userId, dreamId string) error {
	_, err := RedisDB.Do("SREM", dreamId, userId)
	return err
}

func GetLikeAmount(dreamId string) (int, error) {
	return redis.Int(RedisDB.Do("SCARD", dreamId))
}

func HasLike(userId, dreamId string) (bool, error) {
	nums, err := redis.Ints(RedisDB.Do("SMEMBERS", dreamId))
	if err != nil {
		return false, err
	}
	for _, num := range nums {
		n, _ := strconv.Atoi(userId)
		if num == n {
			return true, nil
		}
	}
	return false, nil
}

func GetDreamByLike() ([]Dream, error) {
	var dreams []Dream
	err := DB.Table("dream").Order("likes desc").Find(&dreams).Error
	return dreams, err
}
