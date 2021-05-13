package controllers

import (
	"Dream/common"
	"Dream/conf"
	"Dream/models"
	"Dream/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	ok := services.Save(dream)
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
		drawPath := conf.Config.O.LocalPathPrefix + "/draw/" + draw.Filename
		c.SaveUploadedFile(draw, drawPath)
	}
}

func SoundUpload(c *gin.Context) {
	sound, _ := c.FormFile("sound")
	if sound != nil {
		drawPath := conf.Config.O.LocalPathPrefix + "/sound/" + sound.Filename
		c.SaveUploadedFile(sound, drawPath)
	}
}
