package controllers

import (
	"Dream/common"
	"Dream/models"
	"Dream/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddComment(c *gin.Context) {
	var chat models.Chat
	err := c.ShouldBindJSON(&chat)
	if err != nil {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "please input right format!",
		})
		return
	}
	ok := services.AddComment(chat)
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

func DeleteComment(c *gin.Context) {
	ok := services.DeleteComment(c.Param("id"))
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

func GetCommentsByDream(c *gin.Context) {
	dreams, ok := services.GetCommentsByDream(c.Param("id"))
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
