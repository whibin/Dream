package services

import (
	. "Dream/log"
	"Dream/models"
	"Dream/utils"
)

func SelectOwnDream(uId int) ([]models.Dream, bool) {
	dreams, err := models.SelectOwnDream(uId)
	if err == nil {
		return dreams, true
	}
	Log.WithField("SelectOwnDream", uId).Error(err)
	return nil, false
}

func Save(dream models.Dream) bool {
	// 转成url ----------------
	dream.Draw = utils.LocalPathToUrl(dream.Draw, 1)
	dream.Sound = utils.LocalPathToUrl(dream.Sound, 2)
	// -----------------------
	err := models.Save(dream)
	if err == nil {
		return true
	}
	Log.WithField("Save", dream).Error(err)
	return false
}

func CountByDreamType(uId, t string) (int64, bool) {
	count, err := models.CountByDreamType(uId, t)
	if err != nil {
		Log.WithField("CountByDreamType", uId+"_"+t).Error(err)
		return 0, false
	}
	return count, true
}

func CountByTime() ([]int64, bool) {
	return nil, false
}
