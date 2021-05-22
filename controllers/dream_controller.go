package controllers

import (
	"Dream/common"
	"Dream/conf"
	"Dream/models"
	"Dream/reptile"
	"Dream/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SelectOwnDream(c *gin.Context) {
	uId, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "not a number",
		})
		return
	}
	dreams, ok := services.SelectOwnDream(uId)
	if !ok {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
		Data:    dreams,
	})
}

func Save(c *gin.Context) {
	var dream models.Dream
	err := c.ShouldBindJSON(&dream)
	if err != nil {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "please input right format!",
		})
		return
	}
	ok := services.SaveDream(dream)
	if !ok {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
	})
}

func DrawUpload(c *gin.Context) {
	draw, _ := c.FormFile("draw")
	if draw != nil {
		drawPath := conf.Config.Other.LocalPathPrefix + "/draw/" + draw.Filename
		c.SaveUploadedFile(draw, drawPath)
	}
}

func SoundUpload(c *gin.Context) {
	sound, _ := c.FormFile("sound")
	if sound != nil {
		drawPath := conf.Config.Other.LocalPathPrefix + "/sound/" + sound.Filename
		c.SaveUploadedFile(sound, drawPath)
	}
}

func CountByDreamType(c *gin.Context) {
	uid := c.Param("uid")
	tid := c.Param("type")
	count, ok := services.CountByDreamType(uid, tid)
	if !ok {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
		Data:    count,
	})
}

func CountByTime(c *gin.Context) {
	counts, ok := services.CountByTime()
	if !ok {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
		Data:    counts,
	})
}

func Delete(c *gin.Context) {
	ok := services.DeleteDream(c.Param("uid"), c.Param("id"))
	if !ok {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
	})
}

func Update(c *gin.Context) {
	var dream models.Dream
	err := c.ShouldBindJSON(&dream)
	if err != nil {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "please input right format!",
		})
		return
	}
	ok := services.UpdateDream(dream)
	if !ok {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
	})
}

func GetDreamByTime(c *gin.Context) {
	dreams, ok := services.GetDreamByTime()
	if !ok {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
		Data:    dreams,
	})
}

func GetDreamByType(c *gin.Context) {
	t := c.Param("type")
	dreams, ok := services.GetDreamByType(t)
	if !ok {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
		Data:    dreams,
	})
}

func ExplainDream(c *gin.Context) {
	dreamExplains, err := reptile.ExplainDream(c.Query("keyword"))
	if err != nil {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "explain fail",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
		Data:    dreamExplains,
	})
}

func CountDreamsByUser(c *gin.Context) {
	uid := c.Param("uid")
	count, ok := services.CountDreamsByUser(uid)
	if !ok {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
		Data:    count,
	})
}

func GetReceivedLikes(c *gin.Context) {
	uid := c.Param("uid")
	count, ok := services.GetReceivedLikes(uid)
	if !ok {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
		Data:    count,
	})
}

func DreamMatch(c *gin.Context) {
	uId, _ := strconv.Atoi(c.Param("uid"))
	id, _ := strconv.Atoi(c.Param("id"))
	dream, _ := services.DreamMatch(uId, id)
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
		Data:    dream,
	})
}
